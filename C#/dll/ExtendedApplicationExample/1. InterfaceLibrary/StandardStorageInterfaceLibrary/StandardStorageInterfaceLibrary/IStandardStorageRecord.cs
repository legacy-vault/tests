namespace StandardStorageInterfaceLibrary
{
	public interface IStandardStorageRecord
	{
		// Record's Identifier in the Storage.
		public ulong Id { get; set; }

		// Record's useful Data.
		public string Data { get; set; }
	}
}
