using System;
using TestLibrary;

namespace ConsoleApp1
{
	class Program
	{
		static void Main(string[] args)
		{
			Counter c = new Counter(1, 10, 1);
			c.Count();

			c.TrippleNumber += OnTrippleNumber;
			c.Count();

			c = new Counter(1, 10, 2);
			c.TrippleNumber += OnTrippleNumber;
			c.Count();
		}

		static void OnTrippleNumber(object sender, TrippleNumberEventArgs eventArgs)
		{
			int nf = eventArgs.NumberFound;
			Console.WriteLine($"The Counter has found a Number: {nf}");
		}
	}
}
