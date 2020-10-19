using System;
using System.Net;
using System.Net.Sockets;
using System.Text;

namespace TcpServerUsage
{
	class Program
	{
		static void Main()
		{
			Listen();
		}

		static void Listen()
		{
			string host = "0.0.0.0";
			ushort port = 3000;
			TcpListener server = new TcpListener(IPAddress.Parse(host), port);

			try
			{
				// Configure the Socket.
				server.Server.ExclusiveAddressUse = true;
				server.Server.ReceiveTimeout = 60 * 1000; // 1 Minute.
				server.Server.SendTimeout = 60 * 1000; // 1 Minute.
				server.Start();
				Console.WriteLine($"Listening at {host}:{port}...");

				string unicodeMessage;
				string currentTime;
				while (true)
				{
					TcpClient client = server.AcceptTcpClient();
					NetworkStream stream = client.GetStream();

					// Receive a Message.
					unicodeMessage = ReceiveUnicodeMessage(stream);
					currentTime = DateTime.Now.ToShortTimeString();
					Console.WriteLine($"[{currentTime}] [ IN] {unicodeMessage}");

					// Send a Message.
					string returnedMessage = "Got it!";
					bool success = SendUnicodeMessage(stream, returnedMessage);
					if (success)
					{
						Console.WriteLine($"[{currentTime}] [OUT] {returnedMessage}");
					}
					else
					{
						Console.WriteLine($"[{currentTime}] Failed to send a Message!");
					}

					// Finalization.
					stream.Close();
					client.Close();
				}
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}
			finally
			{
				if (server != null)
				{
					server.Stop();
				}
			}
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
	}
}
