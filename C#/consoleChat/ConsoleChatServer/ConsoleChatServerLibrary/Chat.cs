using System.Collections.Concurrent;

namespace ConsoleChatServerLibrary
{
	public class Chat
	{
		// List of active Users.
		// It is placed at this (Chat) Level instead of the Server Level while there 
		// may be several Servers or Protocols for a single Chat Object and all the 
		// Servers or Protocols must be synchronized. Each Server must have a Link 
		// to the Parent (Chat) Level.
		private ConcurrentDictionary<string, User> ActiveUsers;

		// TCP Server.
		private ChatServer Server;

		// Logger.
		private Logger Logger;

		public Chat(string host, ushort port)
		{
			ActiveUsers = new ConcurrentDictionary<string, User>();
			Logger = new Logger();
			Server = new ChatServer(this, host, port, ActiveUsers, Logger);
		}

		public void Run()
		{
			Server.Run();
		}
	}
}
