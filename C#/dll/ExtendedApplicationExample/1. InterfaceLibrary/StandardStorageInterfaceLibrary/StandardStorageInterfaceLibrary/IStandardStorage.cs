namespace StandardStorageInterfaceLibrary
{
	public interface IStandardStorage
	{
		// Connects to the Storage.
		public void Connect();

		// Disconnects from the Storage.
		public void Disconnect();

		// Gets a Record from the Storage. 
		// Returns null if a Record is not found by its ID.
		public IStandardStorageRecord GetRecordById(ulong id);
	}
}
