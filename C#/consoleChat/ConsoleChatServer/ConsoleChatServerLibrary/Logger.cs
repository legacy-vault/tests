using System;

namespace ConsoleChatServerLibrary
{
	public class Logger
	{
		public void LogMessage(string message)
		{
			Console.WriteLine(message);
		}
		public void LogException(string method, Exception ex)
		{
			if (ex != null)
			{
				Console.WriteLine($"{method} has encountered a {ex.GetType()} Exception: {ex.Message}");
			}
		}
		public void LogError(string method, string error)
		{
			Console.WriteLine($"{method} Method Error: {error}");
		}
		public void LogErrorFailure(string hostMethod, string calledMethod)
		{
			string error = $"{calledMethod} has failed.";
			LogError(hostMethod, error);
		}
	}
}
