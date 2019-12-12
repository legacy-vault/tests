using System;
using System.IO;
using System.IO.Compression;
using System.Text;
using System.Xml;


namespace WorkingWithFileSystems
{
	class Program
	{
		static void Main(string[] args)
		{
			OutputSystemParameters();
			Console.WriteLine();

			GetSystemDrivesInfo();
			Console.WriteLine();

			//Test1();
			Console.WriteLine();

			//Test2();
			Console.WriteLine();

			//TestXmlFile();
			Console.WriteLine();

			TestEncoding();
			Console.WriteLine();
		}

		static void OutputSystemParameters()
		{
			Console.WriteLine($"Path Separator: {Path.PathSeparator}");
			Console.WriteLine($"Directory Separator: {Path.DirectorySeparatorChar}");
			Console.WriteLine($"Current Directory: {Directory.GetCurrentDirectory()}");
			Console.WriteLine($"Current Directory: {Environment.CurrentDirectory}");
			Console.WriteLine($"System Directory: {Environment.SystemDirectory}");
			Console.WriteLine($"Temp Path: {Path.GetTempPath()}");
			Console.WriteLine($"{Environment.GetFolderPath(Environment.SpecialFolder.System)}");
			Console.WriteLine($"{Environment.GetFolderPath(Environment.SpecialFolder.Cookies)}");
			Console.WriteLine($"64-bit OS: {Environment.Is64BitOperatingSystem}");
			Console.WriteLine($"Processor Count: {Environment.ProcessorCount}");
		}

		static void GetSystemDrivesInfo()
		{
			foreach (DriveInfo drive in DriveInfo.GetDrives())
			{
				switch (drive.IsReady)
				{
					case true:
						Console.WriteLine($"{drive.Name} {drive.DriveType} {drive.DriveFormat} TotalSize:{drive.TotalSize} FreeSize:{drive.AvailableFreeSpace}");
						break;

					default:
						Console.WriteLine($"{drive.Name}");
						break;
				}
			}
		}

		static void Test1()
		{
			string userFolder = Environment.GetFolderPath(Environment.SpecialFolder.Personal);
			string dirPath = Path.Combine(new string[] { userFolder, "Code", "Chapter09", "NewFolder" });
			Console.WriteLine($"Working Directory: {dirPath}");
			Console.WriteLine($"Does it exist? {Directory.Exists(dirPath)}");
			Console.WriteLine("Creating it...");
			Directory.CreateDirectory(dirPath);
			Console.WriteLine("Deleting it...");
			Directory.Delete(path: Path.Combine(new string[] { userFolder, "Code" }), recursive: true);
		}

		static void Test2()
		{
			string userFolderPath = Environment.GetFolderPath(Environment.SpecialFolder.Personal);
			string filePath = Path.Combine(new string[] { userFolderPath, "File.txt" });

			// Create a File.
			if (!File.Exists(filePath))
			{
				Console.WriteLine($"Creating a File '{filePath}'...");
				using (StreamWriter file = File.CreateText(filePath))
				{
					file.WriteLine("C# is cool!");
				}
			}

			// Read a File.
			string fileName = Path.GetFileName(filePath);
			string fileNameExt = Path.GetExtension(fileName);
			Console.WriteLine($"Reading a File '{fileName}'... [Ext={fileNameExt}]");
			using (StreamReader file = File.OpenText(filePath))
			{
				string fileContents = file.ReadToEnd();
				Console.WriteLine(fileContents);
			}

			// Pause.
			Console.WriteLine("Press any key to continue...");
			Console.ReadKey();

			// Delete a File.
			Console.WriteLine($"Deleting a File '{filePath}'...");
			File.Delete(filePath);
		}

		static void TestXmlFile()
		{
			Console.WriteLine("Enter the File's Name (without an Extension):");
			string fileName = Console.ReadLine();
			fileName += ".xml.gzip";
			string folderPath = Environment.GetFolderPath(Environment.SpecialFolder.Personal);
			string filePath = Path.Combine(new string[] { folderPath, fileName });
			Console.WriteLine($"Using the File: {filePath}...");
			using (FileStream file = File.OpenWrite(filePath))
			{
				using (GZipStream gzipFile = new GZipStream(file, CompressionLevel.Optimal))
				{
					XmlWriterSettings xmlSettings = new XmlWriterSettings
					{
						Indent = true,
						ConformanceLevel = ConformanceLevel.Document
					};
					using (XmlWriter xmlFile = XmlWriter.Create(gzipFile, xmlSettings))
					{
						try
						{
							xmlFile.WriteStartDocument();
							xmlFile.WriteStartElement("names");
							var names = new string[] { "John", "Alice", "Douglas" };
							foreach (var name in names)
							{
								xmlFile.WriteElementString("name", name);
							}
							xmlFile.WriteEndElement(); // names.
						}
						catch (Exception ex)
						{
							Console.WriteLine($"{ex.GetType()} Exception: {ex.ToString()}");
						}
					}
				}
			}
		}

		static void TestEncoding()
		{
			string message = "AZaz09АЯая";
			Encoding encoding = Encoding.UTF8;

			// Encode.
			byte[] messageEncoded = encoding.GetBytes(message);
			Console.WriteLine($"Message: {message}");
			Console.WriteLine($"Message encoded into UTF-8 ({messageEncoded.Length} Bytes):");
			foreach (byte b in messageEncoded)
			{
				Console.Write($"[{b}]");
			}
			Console.WriteLine();

			// Decode.
			string messageDecoded = encoding.GetString(messageEncoded);
			if (message != messageDecoded)
			{
				throw new Exception("Codec Error");
			}
		}
	}
}
