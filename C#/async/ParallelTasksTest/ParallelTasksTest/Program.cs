using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;

// Synchronization between parallel Tasks of same Hierarchy Level via an external Object.

namespace ParallelTasksTest
{
	class Program
	{
		static void Main(string[] args)
		{
			User user = new User { Name = "John" };
			UserLogin(user);
			List<Task> tasks = new List<Task>();
			Task tA = Task.Run(() => TaskA(user));
			tasks.Add(tA);
			Task tB = Task.Run(() => TaskB(user));
			tasks.Add(tB);
			Task.WhenAll(tasks).Wait();
			Console.WriteLine("Main: All Tasks have finished. Shutdown...");
		}

		static void UserLogin(User user)
		{
			user.SessionId = "ABC-1234567";
			user.IsLoggedIn = true;
		}

		static void UserLogout(User user)
		{
			user.IsLoggedIn = false;
			user.SessionId = String.Empty;
		}

		static void TaskA(User user)
		{
			Console.WriteLine("TaskA: Working...");
			Thread.Sleep(3000);
			UserLogout(user);
			Console.WriteLine("TaskA: Seems that the User has logged out. Aborting...");
		}

		static void TaskB(User user)
		{
			while (user.IsLoggedIn)
			{
				Console.WriteLine($"TaskB: User is logged in. SID='{user.SessionId}'. Working...");
				Thread.Sleep(1000);
			}
			Console.WriteLine($"TaskB: User has logged out. SID='{user.SessionId}'. Aborting...");
		}
	}

	public class User
	{
		public string Name { get; set; }
		public string SessionId { get; set; }
		public bool IsLoggedIn { get; set; }
	}
}
