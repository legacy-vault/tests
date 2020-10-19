using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleChatServerLibrary
{
	public class ChatServer
	{
		// A Link to the Parent (Chat) Level.
		// It is used for Synchronization of the active Users List.
		private Chat Chat;
		public string Host { get; }
		public ushort Port { get; }
		public int NetReadTimeoutMs { get; }
		public int NetWriteTimeoutMs { get; }
		private ConcurrentDictionary<string, User> ActiveUsers;
		private TcpListener tcpListener;
		public Logger Logger { get; }

		// Text Messages.
		public const string QuestionEnterYourName = "Enter your name.";
		public const string QuestionEnterAnotherName = "This name is already used, enter another name.";
		public const string NoticeLoggingOut = "You are now being logged out.";
		public const string SystemMessageUnknownAction = "Unknown action.";

		// System Messages.
		public const char SystemPrefix = '/';
		public const string SystemMessageHelp = "/help";
		public const string SystemMessageMembersList = "/members";
		public const string SystemMessageExit = "/exit";
		public const string SystemMessageQuit = "/quit";
		public const string SystemMessageLogout = "/logout";
		public const string SystemMessagePing = "/ping";
		public const string SystemMessagePong = "/pong";


		public ChatServer
		(
			Chat chat,
			string host,
			ushort port,
			ConcurrentDictionary<string, User> activeUsers,
			Logger logger
		)
		{
			Chat = chat;

			// Get the IP Address or resolve the Host Name into an IP Address.
			IPAddress hostIPAddress;
			bool ok = IPAddress.TryParse(host, out hostIPAddress);
			if (ok)
			{
				Host = host;
			}
			else
			{
				IPHostEntry hostEntry = Dns.GetHostEntry(host);
				if (hostEntry.AddressList.Length == 0)
				{
					throw new Exception($"Host Name '{host}' can not be resolved into an I.P. Address");
				}
				Host = hostEntry.AddressList[0].ToString();
			}

			Port = port;
			NetReadTimeoutMs = 60_000;
			NetWriteTimeoutMs = 60_000;
			ActiveUsers = activeUsers;
			Logger = logger;
		}

		public void Run()
		{
			Start();
			ServeClients();
		}

		// Starts the Server.
		private void Start()
		{
			tcpListener = new TcpListener(IPAddress.Parse(Host), Port);
			tcpListener.Server.ExclusiveAddressUse = true;
			tcpListener.Server.ReceiveTimeout = NetReadTimeoutMs;
			tcpListener.Server.SendTimeout = NetWriteTimeoutMs;
			tcpListener.Start();
			Console.WriteLine($"Server has been started at {Host}:{Port}.");
		}

		// Accepts Clients Connections and serves them in separate Tasks.
		private void ServeClients()
		{
			try
			{
				while (true)
				{
					TcpClient client = tcpListener.AcceptTcpClient();
					Task t = ServeClientAsync(client);
				}
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}
			finally
			{
				Stop();
			}
		}

		// Processes the Client's Connection in a separate Task.
		private async Task ServeClientAsync(TcpClient client)
		{
			await Task.Run(() => ServeClient(client));
		}

		// Processes the Client's Connection.
		private void ServeClient(TcpClient client)
		{
			Message message;

			// Try to log the User into the Chat.
			User user = new User(this, client);
			if (!user.IsActive)
			{
				_ = user.Disconnect();
				return;
			}

			// Greet the User.
			bool ok = user.Greet();
			if (!ok)
			{
				user.LogOutWithBroadcast();
				return;
			}

			// Broadcast a Message about User's successfull Log-In.
			BroadcastMessageUserLogin(user);

			// Receive Messages from the User. Messages, which require logging-out, break the Loop.
			bool logoutIsRequired;
			while (true)
			{
				try
				{
					// Get the Message.
					message = user.ReceiveMessage();
					if (message == null)
					{
						Logger.LogErrorFailure("ServeClient", "ReceiveUnicodeMessage");
						break;
					}

					// Process the Message.
					logoutIsRequired = ProcessMessage(message);
					if (logoutIsRequired) break;
				}
				catch (Exception ex)
				{
					Logger.LogException("ServeClient", ex);
					break;
				}
			}

			// Log the User out.
			user.LogOutWithBroadcast();
		}

		// Processes the Message.
		// Returns 'true' when the Log Out is required.
		private bool ProcessMessage(Message message)
		{
			if (message.IsSystem())
			{
				return ProcessSystemMessage(message);
			}
			else
			{
				return ProcessTextMessage(message);
			}
		}

		// Processes the System Message.
		// Returns 'true' when the Log Out is required.
		private bool ProcessSystemMessage(Message message)
		{
			if (message.SenderUser == null) { return true; }

			switch (message.Text)
			{
				case SystemMessageHelp:
					return !(message.SenderUser.ShowHelp());

				case SystemMessageMembersList:
					return !(message.SenderUser.ShowMembersList());

				case SystemMessageExit:
				case SystemMessageQuit:
				case SystemMessageLogout:
					return true;

				case SystemMessagePing:
					return !(message.SenderUser.Pong());

				default:
					return !(message.SenderUser.ShowWarningUnknownAction());
			}
		}

		// Processes the Text Message.
		// Returns 'true' when the Log Out is required.
		private bool ProcessTextMessage(Message message)
		{
			BroadcastMessage(message);
			return false;
		}

		// Checks whether the specified User Name is free (not used in the active Users List).
		public bool UserNameIsFree(string userName)
		{
			return !this.ActiveUsers.ContainsKey(userName);
		}

		// Creates a Message Text with the List of active Users.
		public string MakeMessageTextMembersList()
		{
			StringBuilder messsage = new StringBuilder();
			messsage.Append("List of active chat members:\r\n");
			messsage.Append("----------------------------\r\n");
			foreach (var member in this.ActiveUsers)
			{
				messsage.Append(member.Value.Name + "\r\n");
			}
			messsage.Append("----------------------------");
			return messsage.ToString();
		}

		// Broadcasts a Message to all the active Users.
		private void BroadcastMessage(Message message, bool logIsRequired = false)
		{
			Task taskMessage = BroadcastMessageAsync(message);
			if (logIsRequired)
			{
				Logger.LogMessage(message.Text);
			}
		}

		// Asynchronously broadcasts a Message to all the the active Users.
		private async Task BroadcastMessageAsync(Message message)
		{
			await Task.Run(() => BroadcastMessageInternal(message));
		}

		// Broadcasts a Message to all the the active Users.
		private void BroadcastMessageInternal(Message message)
		{
			List<Task> tasks = new List<Task>();
			Task t;
			foreach (var item in ActiveUsers)
			{
				User user = item.Value;
				t = Task.Run(() => user.SendMessage(message));
				tasks.Add(t);
			}
			Task.WhenAll(tasks).Wait();
		}

		// Broadcasts a Message about the User's successfull Log-In to all the active Users.
		private void BroadcastMessageUserLogin(User user)
		{
			Message message = Message.UserLogIn(user);
			BroadcastMessage(message, true);
		}

		// Broadcasts a Message about the User's normal Log-Out to all the active Users.
		public void BroadcastMessageUserLogoutNormal(User user)
		{
			Message message = Message.UserLogOutNormal(user);
			BroadcastMessage(message, true);
		}

		// Broadcasts a Message about the User's problematic Log-Out to all the active Users.
		public void BroadcastMessageUserLogoutWithProblems(User user)
		{
			Message message = Message.UserLogOutWithProblem(user);
			BroadcastMessage(message, true);
		}

		// Tries to add a User to a List of active Users.
		// Returns 'false' and sets the 'IsActive' Field of the User to 'false' on Failure.
		public bool MakeUserActive(User user)
		{
			user.IsActive = true;
			bool ok = ActiveUsers.TryAdd(user.Name, user);
			if (!ok)
			{
				user.IsActive = false;
				return false;
			}

			// Success.
			return true;
		}

		// Tries to delete a User from a List of active Users.
		// Returns 'false' on Failure.
		public bool MakeUserInactive(User user)
		{
			User userInList = new User();
			bool ok = ActiveUsers.TryRemove(user.Name, out userInList);
			try
			{
				if (!ok)
				{
					Exception x = new Exception($"User '{user.Name}' is going to be logged out, but it is not active.");
					throw x;
				}
				if ((user.Name != userInList.Name) ||
					(user.TcpClient != userInList.TcpClient) ||
					(user.NetStream != userInList.NetStream))
				{
					Exception x = new Exception($"User '{user.Name}' has damaged Data in the List.");
					throw x;
				}
			}
			catch (Exception ex)
			{
				Logger.LogException("MakeUserInactive", ex);
				return false;
			}

			// Success.
			user.IsActive = false;
			return true;
		}

		// Stops the Server.
		private void Stop()
		{
			if (tcpListener != null)
			{
				tcpListener.Stop();
			}
			Console.WriteLine($"Server at {Host}:{Port} has been stopped.");
		}
	}
}
