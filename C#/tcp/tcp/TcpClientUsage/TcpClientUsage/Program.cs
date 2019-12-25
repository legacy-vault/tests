using System;
using System.Text;
using System.Net;
using System.Net.Sockets;

namespace TcpClientUsage
{
	class Program
	{
		static void Main()
		{
			Send();
		}

		static void Send()
		{
			string host = "127.0.0.1";
			ushort port = 3000;

			try
			{
				string unicodeMessage;
				string currentTime;
				bool success;
				while (true)
				{
					TcpClient client = new TcpClient();

					// Get a Message from User.
					Console.WriteLine("Enter a Message to be sent. An empty Message stops the Program.");
					unicodeMessage = Console.ReadLine();
					if (unicodeMessage.Length == 0) break;

					// Configure the Socket.
					client.Client.ExclusiveAddressUse = true;
					client.Client.ReceiveTimeout = 60 * 1000; // 1 Minute.
					client.Client.SendTimeout = 60 * 1000; // 1 Minute.

					// Connect the Client to the Server.
					client.Connect(host, port);
					NetworkStream stream = client.GetStream();
					Console.WriteLine($"Connected to {host}:{port}.");

					// Send a Message.
					success = SendUnicodeMessage(stream, unicodeMessage);
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
					unicodeMessage = ReceiveUnicodeMessage(stream);
					currentTime = DateTime.Now.ToShortTimeString();
					Console.WriteLine($"[{currentTime}] [REPLY] {unicodeMessage}");

					// Disconnect.
					stream.Close();
					client.Close();
				}
				Console.WriteLine("Shutdown.");
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
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
	}
}
