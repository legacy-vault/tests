using System;

namespace TestLibrary
{
	public class Counter
	{
		private int firstNumber;
		private int lastNumber;
		private int step;
		public Counter(int firstNumber, int lastNumber, int step)
		{
			this.firstNumber = firstNumber;
			this.lastNumber = lastNumber;
			this.step = step;
		}

		public void Count()
		{
			for (int i = firstNumber; i <= lastNumber; i += step)
			{
				if (i % 3 == 0)
				{
					TrippleNumber?.Invoke(this, new TrippleNumberEventArgs(i));
				}
			}
			Console.WriteLine("Counter has finished counting.");
		}

		public event EventHandler<TrippleNumberEventArgs> TrippleNumber;
	}
}
