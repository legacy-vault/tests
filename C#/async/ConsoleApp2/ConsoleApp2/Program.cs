using System;
using System.Collections.Generic;
using System.Diagnostics;
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
				TaskChainExample();
				SubTaskExample();
				WorkSimple();
				WorkWithResults();
				WorkWithResultsAlternative();
			}
			catch (Exception e)
			{
				Console.WriteLine(e.ToString());
			}
		}

		static void TaskChainExample()
		{
			Stopwatch timer = Stopwatch.StartNew();
			timer.Start();
			string id = "xyz-id";
			Task<int> t = Task.Factory.StartNew(() => AsyncDemo.GetSomeDataById(id)).
				ContinueWith(previousTask => AsyncDemo.ProcessSomeData(previousTask.Result));
			t.Wait();
			timer.Stop();
			Console.WriteLine($"Result: {t.Result}.");
			Console.WriteLine($"Chain of Tasks has been completed in {timer.Elapsed}.");
		}

		static void SubTaskExample()
		{
			Stopwatch timer = Stopwatch.StartNew();
			timer.Start();
			string id = "xyz-id";
			Task<int> outerTask = Task.Factory.StartNew(() => AsyncDemo.OuterJob(id));
			outerTask.Wait();
			timer.Stop();
			Console.WriteLine($"Result: {outerTask.Result}.");
			Console.WriteLine($"Chain of Tasks has been completed in {timer.Elapsed}.");
		}

		static void WorkSimple()
		{
			int tasksCount = Environment.ProcessorCount;
			Console.WriteLine($"Running {tasksCount} Tasks...");
			List<Task> tasks = new List<Task>();
			for (int i = 1; i <= tasksCount; i++)
			{
				int n = 5_000 * i;
				Task t = AsyncDemo.CalculateSomeDataAsync(n);
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
				Task<double> t = AsyncDemo.CalculateSomeDataWithResultAsync(n);
				tasks.Add(t);
			}
			Task workTask = Task.WhenAll(tasks);
			workTask.Wait();
			foreach (Task<double> task in tasks)
			{
				Console.WriteLine($"{task.Result.ToString()}");
			}
		}

		// An alternative Way to wait for all Tasks to finish.
		static void WorkWithResultsAlternative()
		{
			int tasksCount = Environment.ProcessorCount;
			Console.WriteLine($"Running {tasksCount} Tasks...");
			Task<double>[] tasks = new Task<double>[tasksCount];
			for (int i = 0; i < tasksCount; i++)
			{
				int n = 5_000 * (i + 1);
				tasks[i] = AsyncDemo.CalculateSomeDataWithResultAsync(n);
			}
			Task.WaitAll(tasks);
			foreach (Task<double> task in tasks)
			{
				Console.WriteLine($"{task.Result.ToString()}");
			}
		}
	}
}
