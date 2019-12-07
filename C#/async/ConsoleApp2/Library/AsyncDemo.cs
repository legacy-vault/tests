using System;
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
	}
}
