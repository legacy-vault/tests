using System;

using StandardStorageInterfaceLibrary;

namespace StandardStorageLibrary
{
	public class StandardStorage : IStandardStorage
	{
		public static readonly string ExceptionMessageConnected = "Connected already";
		public static readonly string ExceptionMessageNotConnected = "Not yet connected";

		private bool isConnected;

		public void Connect()
		{
			lock (this)
			{

				if (isConnected) throw new Exception(ExceptionMessageConnected);

				// Connect...
				//...

				isConnected = true;
			}
		}

		public void Disconnect()
		{
			lock (this)
			{

				if (!isConnected) throw new Exception(ExceptionMessageNotConnected);

				// Disconnect...
				//...

				isConnected = false;
			}
		}

		public IStandardStorageRecord GetRecordById(ulong id)
		{
			if (!isConnected) throw new Exception(ExceptionMessageNotConnected);

			// Get a Record.
			// Emulate the Search in a Storage.
			switch (id)
			{
				case 1: return new StandardStorageRecord { Id = 1, Data = "One" };

				case 2: return new StandardStorageRecord { Id = 2, Data = "Two" };

				default: return null;
			}
		}
	}
}
