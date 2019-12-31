using ConsoleChatServerLibrary;
using System;

// This is a Demonstration of a very simple Console Chat Server.

namespace ConsoleChatServer
{
	class Program
	{
		static void Main(string[] args)
		{
			try
			{
				ChatServerSettings settings = new ChatServerSettings(args);
				Chat chat = new Chat(settings.Host, settings.Port);
				chat.Run();
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}
		}
	}
}
