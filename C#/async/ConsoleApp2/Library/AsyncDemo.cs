using System;
using System.Threading;
using System.Threading.Tasks;

namespace Library
{
	public class AsyncDemo
	{
		public static async Task CalculateSomeDataAsync(int n)
		{
			double result = await Task.Run(() => CalculateSomeData(n));
			Console.WriteLine($"{result}");
		}

		public static async Task<double> CalculateSomeDataWithResultAsync(int n)
		{
			return await Task.Run(() => CalculateSomeData(n));
		}

		private static double CalculateSomeData(int n)
		{
			double result = 0;
			for (int i = 1; i <= n; i++)
			{
				for (int j = 1; j <= n; j++)
				{
					try
					{
						checked
						{
							result = Convert.ToDouble(i + j) * Convert.ToDouble(i + j);
						}
					}
					catch (OverflowException e)
					{
						Console.WriteLine($"{e.ToString()}");
						return 0;
					}
				}
			}
			return result;
		}

		public static int GetSomeDataById(string id)
		{
			Console.WriteLine($"GetSomeDataById('{id}')");
			Thread.Sleep(2000);
			return 123;
		}

		public static int ProcessSomeData(int data)
		{
			Console.WriteLine($"ProcessSomeData({data})");
			Thread.Sleep(3000);
			return -data;
		}

		public static int OuterJob(string id)
		{
			Console.WriteLine($"OuterJob('{id}')");
			Thread.Sleep(2000);
			var subTask = Task.Factory.StartNew(() => InnerJob(123), TaskCreationOptions.AttachedToParent);
			Thread.Sleep(1000); // Parallel with a sub-Task.
			return subTask.Result;
		}

		public static int InnerJob(int data)
		{
			Console.WriteLine($"InnerJob({data})");
			Thread.Sleep(5000);
			return data * (-2);
		}
	}
}
