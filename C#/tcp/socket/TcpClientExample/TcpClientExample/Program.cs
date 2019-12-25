using System;
using System.Text;
using System.Net;
using System.Net.Sockets;

namespace TcpClientExample
{
	class Program
	{
		static void Main(string[] args)
		{
			Send();
		}

		static void Send()
		{
			string host = "127.0.0.1";
			ushort port = 3000;
			IPEndPoint address = new IPEndPoint(IPAddress.Parse(host), port);

			try
			{
				string unicodeMessage;
				string currentTime;
				bool success;
				while (true)
				{
					// Get a Message from User.
					Console.WriteLine("Enter a Message to be sent. An empty Message stops the Program.");
					unicodeMessage = Console.ReadLine();
					if (unicodeMessage.Length == 0) break;

					// Configure the Socket.
					Socket socket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
					socket.ExclusiveAddressUse = true;
					socket.ReceiveTimeout = 60 * 1000; // 1 Minute.
					socket.SendTimeout = 60 * 1000; // 1 Minute.

					// Connect to the Socket.
					socket.Connect(address);
					Console.WriteLine($"Connected to {host}:{port}.");

					// Send a Message.
					success = SendUnicodeMessage(socket, unicodeMessage);
					currentTime = DateTime.Now.ToShortTimeString();
					if (success)
					{
						Console.WriteLine($"[{currentTime}] Message has been sent.");
					}
					else
					{
						Console.WriteLine($"[{currentTime}] Failed to send a Message!");
						continue;
					}

					// Receive a Reply.
					unicodeMessage = ReceiveUnicodeMessage(socket);
					currentTime = DateTime.Now.ToShortTimeString();
					Console.WriteLine($"[{currentTime}] [REPLY] {unicodeMessage}");

					// Disconnect.
					socket.Shutdown(SocketShutdown.Both);
					socket.Close();
				}
				Console.WriteLine("Shutdown.");
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}
		}

		static bool SendUnicodeMessage(Socket socket, string message)
		{
			byte[] sendBuffer = Encoding.Unicode.GetBytes(message);
			int sentBytesCount = socket.Send(sendBuffer);
			if (sendBuffer.Length == sentBytesCount) return true;
			return false;
		}

		static string ReceiveUnicodeMessage(Socket socket)
		{
			StringBuilder receivedString = new StringBuilder();
			int receivedDataSize = 0;
			byte[] receiveBuffer = new byte[4096];

			try
			{
				do
				{
					receivedDataSize = socket.Receive(receiveBuffer);
					receivedString.Append(Encoding.Unicode.GetString(receiveBuffer, 0, receivedDataSize));
				}
				while (socket.Available > 0);
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}

			return receivedString.ToString();
		}
	}
}
