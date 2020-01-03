using HttpServerLibrary;
using System.Threading;

namespace HttpListenerUsage
{
	class Program
	{
		static void Main()
		{
			Server server = new Server("http", "127.0.0.1", "3000");
			bool ok = server.Start();
			if (!ok) return;
			Thread.Sleep(15_000);
			server.Stop();
		}
	}
}
