using System;
using System.Text;
using System.Net.Sockets;
using System.Threading.Tasks;
using System.Threading;

// This is a Demonstration of a very simple Console Chat Server.
// As opposed to the Server, the Client has been written just to test the Server fast.

namespace ConsoleChatClient
{
	class Program
	{
		static void Main(string[] args)
		{
			string host;
			switch (args.Length)
			{
				case 2:
					host = args[0];
					ushort port = ushort.Parse(args[1]);
					Client(host, port);
					return;

				case 1:
					host = args[0];
					Client(host);
					return;

				default:
					Console.WriteLine("Parameters are not set!\r\n\r\n" +
					"Usage:\r\n" +
					"  <app> host port\r\n" +
					"  <app> host\r\n\r\n" +
					"Example:\r\n" +
					"  ConsoleChatClient.exe localhost 3000");
					return;
			}
		}

		public const string NoticeLoggingOut = "You are now being logged out.";
		public const string SystemMessagePing = "/ping";
		public const string SystemMessagePong = "/pong";
		public const int KeepAlivePingIntervalMs = 60_000 - 5_000;

		static void Client(string host = "127.0.0.1", ushort port = 3000)
		{
			TcpClient client = new TcpClient();
			client.Client.ExclusiveAddressUse = true;
			client.Client.ReceiveTimeout = 60 * 1000; // 1 Minute.
			client.Client.SendTimeout = 60 * 1000; // 1 Minute.
			try
			{
				bool success;
				LogoutConfirmation logoutConfirmation = new LogoutConfirmation { HasBeenReceived = false };

				// Connect the Client to the Server.
				client.Connect(host, port);
				NetworkStream stream = client.GetStream();
				Console.WriteLine($"Connected to {host}:{port}.");

				// Log in.
				success = ClientLogin(stream);
				if (!success)
				{
					Console.WriteLine("Failed to log in.");
					return;
				}

				// Send KeepAlivePing Messages in a Parallel Task.
				SendKeepAlivePingsAsync(stream);

				// Receive Messages in a Parallel Task.
				ReceiveMessagesAsync(stream, logoutConfirmation);

				string unicodeMessage;
				while (true)
				{
					// Get a Message from User.
					unicodeMessage = Console.ReadLine();
					while (unicodeMessage.Length == 0)
					{
						unicodeMessage = Console.ReadLine();
					}
					Console.WriteLine();

					// Send a Message.
					success = SendUnicodeMessage(stream, unicodeMessage);
					if (!success)
					{
						Console.WriteLine("[ERROR]");
						continue;
					}
					if (MessageIsLogout(unicodeMessage)) break;
				}

				// Await for Logout Confirmation.
				while (!logoutConfirmation.HasBeenReceived)
				{
					Thread.Sleep(1000);
				}

				// Disconnect.
				Console.WriteLine("Shutdown.");
				stream.Close();
				client.Close();
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}
		}

		static bool ClientLogin(NetworkStream stream)
		{
			string intro;
			string userName = "";
			string expectedReply;
			string expectedReply2;
			string response;
			bool success;
			bool loggedIn = false;

			try
			{
				// Introduction.
				intro = ReceiveUnicodeMessage(stream);
				if (intro.Length == 0) { return false; }
				Console.WriteLine(intro);

				while (!loggedIn)
				{
					// Get a Message from User.
					userName = Console.ReadLine();
					while (userName.Length == 0)
					{
						userName = Console.ReadLine();
					}
					Console.WriteLine();

					// Send a Message.
					success = SendUnicodeMessage(stream, userName);
					if (!success) Console.WriteLine("[ERROR]");
					expectedReply = $"{userName}, welcome to the chat!\r\nType '/help' for help or type a message for everyone.";

					// Await the Confirmation.
					response = ReceiveUnicodeMessage(stream);
					if (response.Length == 0) { return false; }
					Console.WriteLine(response);
					if (response == expectedReply)
					{
						loggedIn = true;
					}
				}

				// Await the second Confirmation.
				expectedReply2 = $"User '{userName}' has logged in.";
				response = ReceiveUnicodeMessage(stream);
				if (response.Length == 0) { return false; }
				Console.WriteLine(response);
				if (response != expectedReply2) return false;

			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
				return false;
			}
			return true;
		}

		static async Task ReceiveMessagesAsync(NetworkStream stream, LogoutConfirmation logoutConfirmation)
		{
			await Task.Run(() => ReceiveMessages(stream, logoutConfirmation));
		}

		static void ReceiveMessages(NetworkStream stream, LogoutConfirmation logoutConfirmation)
		{
			string unicodeMessage;
			while (true)
			{
				unicodeMessage = ReceiveUnicodeMessage(stream);
				if (unicodeMessage.Length == 0) { return; }
				switch (unicodeMessage)
				{
					case SystemMessagePong:
						continue;

					case NoticeLoggingOut:
						logoutConfirmation.HasBeenReceived = true;
						return;
				}
				Console.WriteLine($"{unicodeMessage}");
			}
		}

		static async Task SendKeepAlivePingsAsync(NetworkStream stream)
		{
			await Task.Run(() => SendKeepAlivePings(stream));
		}

		static void SendKeepAlivePings(NetworkStream stream)
		{
			int failuresCount = 0;
			bool ok;
			while (failuresCount < 3)
			{
				ok = SendUnicodeMessage(stream, SystemMessagePing);
				if (!ok)
				{
					Console.WriteLine("SendKeepAlivePings: Failure.");
					failuresCount++;
				}
				Thread.Sleep(KeepAlivePingIntervalMs);
			}
		}

		static bool SendUnicodeMessage(NetworkStream stream, string message)
		{
			byte[] sendBuffer = Encoding.Unicode.GetBytes(message);
			try
			{
				stream.Write(sendBuffer, 0, sendBuffer.Length);
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
				return false;
			}
			return true;
		}

		static string ReceiveUnicodeMessage(NetworkStream stream)
		{
			StringBuilder receivedString = new StringBuilder();
			int receivedDataSize;
			byte[] receiveBuffer = new byte[4096];

			try
			{
				do
				{
					receivedDataSize = stream.Read(receiveBuffer, 0, receiveBuffer.Length);
					receivedString.Append(Encoding.Unicode.GetString(receiveBuffer, 0, receivedDataSize));
				}
				while (stream.DataAvailable);
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}

			return receivedString.ToString();
		}

		static bool MessageIsLogout(string message)
		{
			switch (message)
			{
				case "/exit":
				case "/quit":
				case "/logout":
					return true;

				default: return false;
			}
		}
	}

	class LogoutConfirmation
	{
		public bool HasBeenReceived { get; set; }
	}
}
