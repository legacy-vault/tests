using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Library;

namespace ConsoleApp2
{
	class Program
	{
		static void Main(string[] args)
		{
			try
			{
				WorkSimple();
				WorkWithResults();
			}
			catch (Exception e)
			{
				Console.WriteLine(e.ToString());
			}
		}

		static void WorkSimple()
		{
			int tasksCount = Environment.ProcessorCount;
			Console.WriteLine($"Running {tasksCount} Tasks...");
			List<Task> tasks = new List<Task>();
			for (int i = 1; i <= tasksCount; i++)
			{
				int n = 5_000 * i;
				Task t = Task.Run(() => AsyncDemo.CalculateSomeDataAsync(n));
				tasks.Add(t);
			}
			Task workTask = Task.WhenAll(tasks);
			workTask.Wait();
		}

		static void WorkWithResults()
		{
			int tasksCount = Environment.ProcessorCount;
			Console.WriteLine($"Running {tasksCount} Tasks...");
			List<Task<double>> tasks = new List<Task<double>>();
			for (int i = 1; i <= tasksCount; i++)
			{
				int n = 5_000 * i;
				Task<double> t = Task<double>.Run(() => AsyncDemo.CalculateSomeDataWithResultAsync(n));
				tasks.Add(t);
			}
			Task workTask = Task.WhenAll(tasks);
			workTask.Wait();
			foreach (Task<double> task in tasks)
			{
				Console.WriteLine($"{task.Result.ToString()}");
			}
		}
	}
}
