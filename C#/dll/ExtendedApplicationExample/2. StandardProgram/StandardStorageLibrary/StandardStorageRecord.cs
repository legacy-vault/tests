using StandardStorageInterfaceLibrary;

namespace StandardStorageLibrary
{
	public class StandardStorageRecord : IStandardStorageRecord
	{
		public ulong Id { get; set; }
		public string Data { get; set; }
	}
}
