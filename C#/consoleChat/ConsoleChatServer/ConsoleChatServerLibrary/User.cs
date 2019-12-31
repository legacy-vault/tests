using System;
using System.Net.Sockets;
using System.Text;

namespace ConsoleChatServerLibrary
{
	public class User
	{
		public ChatServer Server { get; set; }
		public string Name { get; set; }
		public TcpClient TcpClient { get; set; }
		public NetworkStream NetStream { get; set; }
		public bool IsActive { get; set; }

		public User() { }

		// The Connstructor tries to log the User into the Server.
		// Sets the 'IsActive' Field of the returned User to 'false' on Failure.
		public User(ChatServer server, TcpClient client)
		{
			// Fool Check.
			if ((server == null) ||
				(client == null) ||
				(server.Logger == null))
			{
				return;
			}

			try
			{
				this.IsActive = false;
				this.Server = server;
				this.Name = String.Empty;
				this.TcpClient = client;
				this.NetStream = client.GetStream();

				bool ok = this.LogIn();
				if (!ok)
				{
					Exception x = new Exception($"Failed to log the User '{this.Name}' into the Server.");
					throw x;
				}
			}
			catch (Exception ex)
			{
				this.Server.Logger.LogException("User", ex);
			}
		}

		// Tries to log the User into the Server.
		// Returns 'false' on Failure.
		private bool LogIn()
		{
			bool ok = this.GetFreeName();
			if (!ok)
			{
				this.IsActive = false;
				return false;
			}
			return this.Server.MakeUserActive(this);
		}

		// Gets a free Name for the User.
		// returns 'false' on Failure.
		private bool GetFreeName()
		{
			bool userNameIsTaken;
			string userName = this.GetNameFirstTime();
			if (userName.Length == 0) { return false; }
			userNameIsTaken = !this.Server.UserNameIsFree(userName);
			while (userNameIsTaken)
			{
				userName = this.GetAnotherName();
				if (userName.Length == 0) { return false; }
				userNameIsTaken = !this.Server.UserNameIsFree(userName);
			}

			// A free Name is found.
			this.Name = userName;
			return true;
		}

		// Tries to get User's Name for the first Time.
		private string GetNameFirstTime()
		{
			return this.GetNameNthTime(1);
		}

		// Tries to get User's Name again.
		private string GetAnotherName()
		{
			return this.GetNameNthTime(2);
		}

		// Tries to get User's Name for the N-th Time.
		// Returns an empty String on Failure.
		private string GetNameNthTime(int n)
		{
			if (n < 1) { return String.Empty; }

			// Ask for a Name.
			bool ok;
			Message message;
			if (n == 1)
			{
				message = Message.TextMessageFromSystem(ChatServer.QuestionEnterYourName);
				ok = this.SendMessage(message);
			}
			else // n > 1.
			{
				message = Message.TextMessageFromSystem(ChatServer.QuestionEnterAnotherName);
				ok = this.SendMessage(message);
			}
			if (!ok) return String.Empty;

			// Wait for a Reply which is not a System Message (e.g., not a Ping Message).
			message = this.ReceiveMessage();
			if (message == null) { return String.Empty; }
			while (message.IsSystem())
			{
				message = this.ReceiveMessage();
				if (message == null) { return String.Empty; }
			}
			return message.Text;
		}

		// Receives a Message from a User.
		// Returns null on Error.
		public Message ReceiveMessage()
		{
			string messageText = this.ReceiveUnicodeMessage();

			// Learn the Message Type.
			if ((messageText == null) || (messageText.Length < 1)) { return null; }
			if (messageText[0] == ChatServer.SystemPrefix)
			{
				// Message Type is System.
				return Message.SystemMessageFromUser(this, messageText);
			}
			else
			{
				// Message Type is Text.
				return Message.TextMessageFromUser(this, messageText);
			}
		}

		// Receives a Unicode String from the User.
		// Returns an empty String on Failure.
		private string ReceiveUnicodeMessage()
		{
			StringBuilder receivedString = new StringBuilder();
			int receivedDataSize;
			byte[] receiveBuffer = new byte[4096];

			try
			{
				do
				{
					receivedDataSize = this.NetStream.Read(receiveBuffer, 0, receiveBuffer.Length);
					receivedString.Append(Encoding.Unicode.GetString(receiveBuffer, 0, receivedDataSize));
				}
				while (this.NetStream.DataAvailable);
			}
			catch (Exception ex)
			{
				this.Server.Logger.LogException("ReceiveUnicodeMessage", ex);
				return String.Empty;
			}

			return receivedString.ToString();
		}

		// Sends a Message to the User.
		// Returns 'false' on Error.
		public bool SendMessage(Message message)
		{
			if (message == null)
			{
				throw new Exception("Can not send a null Message!");
			}
			switch (message.Type)
			{
				case MessageType.System:
					return this.SendSystemMessage(message);

				case MessageType.Text:
					return this.SendTextMessage(message);

				default:
					return false;
			}
		}

		// Sends a System Message to the User.
		// Returns 'false' on Error.
		private bool SendSystemMessage(Message message)
		{
			return this.SendUnicodeMessage(message.Text);
		}

		// Sends a Text Message to the User.
		// Returns 'false' on Error.
		private bool SendTextMessage(Message message)
		{
			switch (message.SenderType)
			{
				case MessageSenderType.System:
					return this.SendUnicodeMessage(message.Text);

				case MessageSenderType.User:
					return this.SendTextMessageFromUser(message);

				default:
					return false;
			}
		}

		// Sends a Text Message from a User to the User.
		// Returns 'false' on Error.
		private bool SendTextMessageFromUser(Message message)
		{
			string rawText = $"[{message.Time.ToString(Message.TimeFormat)}] Message from '{message.SenderUser.Name}'.\r\n" +
				message.Text;
			return this.SendUnicodeMessage(rawText);
		}

		// Sends a Unicode String to the User.
		// Returns 'false' on Failure.
		private bool SendUnicodeMessage(string message)
		{
			byte[] sendBuffer = Encoding.Unicode.GetBytes(message);
			try
			{
				this.NetStream.Write(sendBuffer, 0, sendBuffer.Length);
			}
			catch (Exception ex)
			{
				this.Server.Logger.LogException("SendUnicodeMessage", ex);
				return false;
			}

			return true;
		}

		// Disconnects a non-logged-in User from the Server.
		// Returns 'false' on Failure.
		public bool Disconnect()
		{
			try
			{
				if (this.IsActive)
				{
					Exception x = new Exception($"User '{this.Name}' is active and can not be disconnected.");
					throw x;
				}
				this.NetStream.Close();
				this.TcpClient.Close();
			}
			catch (Exception ex)
			{
				this.Server.Logger.LogException("Disconnect", ex);
				return false;
			}

			return true;
		}

		// Sends a Greeting Message to the User.
		public bool Greet()
		{
			Message message = Message.UserGreeting(this);
			return this.SendMessage(message);
		}

		// Logs an active User out of the Server.
		// Returns 'false' on any Failure.
		private bool LogOut()
		{
			bool success = true;
			Message message = Message.TextMessageFromSystem(ChatServer.NoticeLoggingOut);
			bool ok = this.SendMessage(message);
			if (!ok)
			{
				this.Server.Logger.LogErrorFailure("LogOut", "SendUnicodeMessage");
				success = false;
			}
			ok = this.Server.MakeUserInactive(this);
			if (!ok)
			{
				this.Server.Logger.LogErrorFailure("LogOut", "MakeUserInactive");
				success = false;
			}
			ok = this.Disconnect();
			if (!ok)
			{
				this.Server.Logger.LogErrorFailure("LogOut", "Disconnect");
				success = false;
			}
			return success;
		}

		// Logs an active User out of the Server and broadcasts a Message about the Event to all the active Users.
		public void LogOutWithBroadcast()
		{
			bool ok = this.LogOut();
			if (ok)
			{
				this.Server.BroadcastMessageUserLogoutNormal(this);
			}
			else
			{
				this.Server.BroadcastMessageUserLogoutWithProblems(this);
			}
		}

		// Sends the Help Information to the User.
		// Returns 'false' on Failure.
		public bool ShowHelp()
		{
			Message message = Message.TextMessageFromSystem(Message.TextHelp());
			return this.SendMessage(message);
		}

		// Sends the Member List Information to the User.
		// Returns 'false' on Failure.
		public bool ShowMembersList()
		{
			string text = this.Server.MakeMessageTextMembersList();
			Message message = Message.TextMessageFromSystem(text);
			return this.SendMessage(message);
		}

		// Sends the 'PONG' Message to the User.
		// Returns 'false' on Failure.
		public bool Pong()
		{
			Message message = Message.TextMessageFromSystem(ChatServer.SystemMessagePong);
			return this.SendMessage(message);
		}

		// Sends the Message about unknown Action to the User.
		// Returns 'false' on Failure.
		public bool ShowWarningUnknownAction()
		{
			Message message = Message.TextMessageFromSystem(ChatServer.SystemMessageUnknownAction);
			return this.SendMessage(message);
		}
	}
}
