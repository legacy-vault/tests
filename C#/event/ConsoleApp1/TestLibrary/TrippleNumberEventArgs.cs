using System;
using System.Collections.Generic;
using System.Text;

namespace TestLibrary
{
	public class TrippleNumberEventArgs : EventArgs
	{
		public int NumberFound { get; }
		public TrippleNumberEventArgs(int numberFound)
		{
			this.NumberFound = numberFound;
		}
	}
}
