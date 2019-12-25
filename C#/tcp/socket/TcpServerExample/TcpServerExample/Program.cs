using System;
using System.Text;
using System.Net;
using System.Net.Sockets;

namespace TcpServerExample
{
	class Program
	{
		static void Main(string[] args)
		{
			Listen();
		}

		static void Listen()
		{
			string host = "0.0.0.0";
			ushort port = 3000;
			IPEndPoint address = new IPEndPoint(IPAddress.Parse(host), port);
			Socket socket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
			try
			{
				// Take the Socket.
				socket.ExclusiveAddressUse = true;
				socket.ReceiveTimeout = 60 * 1000; // 1 Minute.
				socket.SendTimeout = 60 * 1000; // 1 Minute.
				socket.Bind(address);
				socket.Listen(8);
				Console.WriteLine($"Listening at {host}:{port}...");

				string unicodeMessage;
				string currentTime;
				while (true)
				{
					Socket handler = socket.Accept();

					// Receive a Message.
					unicodeMessage = ReceiveUnicodeMessage(handler);
					currentTime = DateTime.Now.ToShortTimeString();
					Console.WriteLine($"[{currentTime}] [ IN] {unicodeMessage}");

					// Send a Message.
					string returnedMessage = "Got it!";
					bool success = SendUnicodeMessage(handler, returnedMessage);
					if (success)
					{
						Console.WriteLine($"[{currentTime}] [OUT] {returnedMessage}");
					}
					else
					{
						Console.WriteLine($"[{currentTime}] Failed to send a Message!");
					}

					// Finalization.
					handler.Shutdown(SocketShutdown.Both);
					handler.Close();
				}
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}
		}

		static string ReceiveUnicodeMessage(Socket handler)
		{
			StringBuilder receivedString = new StringBuilder();
			int receivedDataSize = 0;
			byte[] receiveBuffer = new byte[4096];

			try
			{
				do
				{
					receivedDataSize = handler.Receive(receiveBuffer);
					receivedString.Append(Encoding.Unicode.GetString(receiveBuffer, 0, receivedDataSize));
				}
				while (handler.Available > 0);
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}

			return receivedString.ToString();
		}

		static bool SendUnicodeMessage(Socket handler, string message)
		{
			byte[] sendBuffer = Encoding.Unicode.GetBytes(message);
			int sentBytesCount = handler.Send(sendBuffer);
			if (sendBuffer.Length == sentBytesCount) return true;
			return false;
		}
	}
}
