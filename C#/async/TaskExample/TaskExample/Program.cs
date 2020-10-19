using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace TaskExample
{
	class Program
	{
		// Parameter Range for the benchmarked Methods.
		const double iFirst = 1;
		const double iLast = 1_000_000_000;

		static void Main()
		{
			double durationSeconds;

			RunBenchmark(MethodA, out durationSeconds);
			Console.WriteLine(MakeBenchmarkReport("MethodA", durationSeconds));

			RunBenchmark(MethodB, out durationSeconds);
			Console.WriteLine(MakeBenchmarkReport("MethodB", durationSeconds));
		}

		// Composes the Benchmark Report Text.
		static string MakeBenchmarkReport(string methodName, double durationSeconds)
		{
			return $"{methodName} has been performed in {durationSeconds} seconds.";
		}

		// Runs the specified Method and calculates the Time used for its Execution.
		// Returns the spent Time as Seconds.
		static void RunBenchmark(Action action, out double durationSeconds)
		{
			DateTime timeOfStart = DateTime.Now;
			action.Invoke();
			DateTime timeOfStop = DateTime.Now;
			durationSeconds = timeOfStop.Subtract(timeOfStart).TotalSeconds;
			return;
		}

		// A sample Method using the multi-threaded Function.
		static void MethodA()
		{
			ShortFunctionRangeCollectorMT(iFirst, iLast);
		}

		// A sample Method using the single-threaded Function.
		static void MethodB()
		{
			ShortFunctionRangeCollector(iFirst, iLast);
		}

		// Makes a Series of some Calculations for the specified Parameter Range.
		// Returns the Sum of the Calculations.
		// Uses multiple Threads (Tasks).
		static double ShortFunctionRangeCollectorMT(double iFirst, double iLast)
		{
			// Create the Tasks. Separate the input Parameters into different Tasks (CPU Threads).
			int tasksCount = Environment.ProcessorCount;
			List<Task<double>> tasks = new List<Task<double>>();
			Task<double> t;
			double iterationsPerTask = Math.Ceiling((iLast - iFirst + 1) / (double)(tasksCount));
			double i2i1delta = iterationsPerTask - 1; // Cached Difference between "Edges".
			for (int i = 0; i < tasksCount - 1; i++)
			{
				double i1 = iFirst + (i * iterationsPerTask);
				t = Task.Run(() => ShortFunctionRangeCollector(i1, i1 + i2i1delta));
				tasks.Add(t);
			}
			double i1LastTask = iFirst + (tasksCount - 1) * iterationsPerTask; // Last Task may be shorter than others.
			t = Task.Run(() => ShortFunctionRangeCollector(i1LastTask, iLast));
			tasks.Add(t);

			// Wait for all the Tasks to complete.
			Task tAll = Task.WhenAll(tasks);
			tAll.Wait();
			var sum = 0.0;
			foreach (Task<double> task in tasks)
			{
				sum += task.Result;
			}
			return sum;
		}

		// Makes a Series of some Calculations for the specified Parameter Range.
		// Returns the Sum of the Calculations.
		static double ShortFunctionRangeCollector(double i1, double i2)
		{
			//Console.WriteLine($"ShortFunctionRangeCollector({i1},{i2});"); // Debug.
			var sum = 0.0;
			for (double i = i1; i <= i2; i++)
			{
				sum += ShortFunction(i);
			}
			return sum;
		}

		// Makes some Calculations for the specified Parameter.
		static double ShortFunction(double i)
		{
			return i / Math.PI;
		}
	}
}
