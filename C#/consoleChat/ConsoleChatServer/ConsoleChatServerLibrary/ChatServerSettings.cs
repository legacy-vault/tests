using System;

namespace ConsoleChatServerLibrary
{
	public class ChatServerSettings
	{
		public const string DefaultHost = "0.0.0.0";
		public const ushort DefaultPort = 3000;
		public string Host { get; private set; }
		public ushort Port { get; private set; }

		public ChatServerSettings(string host = DefaultHost, ushort port = DefaultPort)
		{
			Init(host, port);
		}
		public ChatServerSettings(string[] programArguments)
		{
			if (programArguments == null)
			{
				throw new Exception("Arguments are not set");
			}
			switch (programArguments.Length)
			{
				case 0:
					Init();
					return;

				case 1:
					Init(programArguments[0]);
					return;

				case 2:
					ushort port = ushort.Parse(programArguments[1]);
					Init(programArguments[0], port);
					return;

				default:
					throw new Exception("Unsupported Arguments Format");
			}
		}

		private void Init(string host = DefaultHost, ushort port = DefaultPort)
		{
			this.Host = host;
			this.Port = port;
		}
	}
}
