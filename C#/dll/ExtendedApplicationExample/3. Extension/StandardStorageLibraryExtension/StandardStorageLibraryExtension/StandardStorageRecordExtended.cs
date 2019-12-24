using StandardStorageInterfaceLibrary;

namespace StandardStorageLibraryExtension
{
	public class StandardStorageRecordExtended : IStandardStorageRecord
	{
		public ulong Id { get; set; }
		public string Data { get; set; }
	}
}
