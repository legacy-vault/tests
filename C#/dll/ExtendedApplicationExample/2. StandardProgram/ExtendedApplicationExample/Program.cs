using System;
using System.Reflection;

using StandardStorageInterfaceLibrary;
using StandardStorageLibrary;

namespace ExtendedApplicationExample
{
	class Program
	{
		static void Main(string[] args)
		{
			ulong[] ids = { 1, 2, 3, 4 };

			try
			{
				UseStandardStorage(ids);

				string pathToPlugin = "../../../plugin/storage.dll";
				object pluggedObject = EnablePlugin(pathToPlugin);
				IStandardStorage extendedStorage = pluggedObject as IStandardStorage;
				UseExtendedStorage(extendedStorage, ids);
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}
		}

		static void UseStandardStorage(ulong[] ids)
		{
			IStandardStorage storage = new StandardStorage();
			Console.WriteLine("Using the built-in Storage...");
			UseStorage(storage, ids);
		}

		static void UseStorage(IStandardStorage storage, ulong[] ids)
		{
			if (storage == null)
			{
				throw new Exception("Storage is not set");
			}
			storage.Connect();
			foreach (var id in ids)
			{
				GetPrintRecord(storage, id);
			}
			storage.Disconnect();
			Console.WriteLine();
		}

		static void GetPrintRecord(IStandardStorage storage, ulong id)
		{
			IStandardStorageRecord record = storage.GetRecordById(id);
			if (record != null)
			{
				Console.WriteLine($"Obtained a Record (ID={id}): [{record.Data}].");
			}
			else
			{
				Console.WriteLine($"Record with ID={id} does not exist.");
			}
		}

		// Loads a compiled external Library (DLL), creates an Object Instance of the loaded Class and returns it.
		static object EnablePlugin(string pathToPlugin)
		{
			Assembly pluginAssembly = Assembly.LoadFrom(pathToPlugin);
			Type extendedStorageType = pluginAssembly.GetType("StandardStorageLibraryExtension.StandardStorageExtended", true, true);
			object extendedStorage = Activator.CreateInstance(extendedStorageType);
			return extendedStorage;
		}

		static void UseExtendedStorage(IStandardStorage storage, ulong[] ids)
		{
			Console.WriteLine("Using an extended Storage...");
			UseStorage(storage, ids);
		}
	}
}
