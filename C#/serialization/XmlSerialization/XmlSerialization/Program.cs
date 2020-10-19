using System;
using System.Collections.Generic;
using System.IO;
using System.Xml.Serialization;

namespace XmlSerialization
{
	[XmlType("person")]
	public class Person
	{
		[XmlAttribute("name")]
		public string Name { get; set; }
		[XmlAttribute("age")]
		public int Age { get; set; }

		[XmlArray("books")]
		public Book[] Books { get; set; }
	}

	[XmlType("book")]
	public class Book
	{
		[XmlAttribute("title")]
		public string Title { get; set; }
	}

	class Program
	{
		static void Main(string[] args)
		{
			try
			{
				SerializeData();
				Console.WriteLine();

				DeserializeData();
				Console.WriteLine();
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.ToString()}.");
			}
		}

		static readonly string CollectionNameLowCase = "people";

		static void SerializeData()
		{
			var people = new List<Person>();
			var booksCollectionA = new Book[] { new Book { Title = "Red Book" }, new Book { Title = "Green Book" } };
			var booksCollectionB = new Book[] { new Book { Title = "Yellow Book" }, new Book { Title = "Blue Book" } };
			people.Add(new Person() { Name = "John", Age = 123, Books = booksCollectionA });
			people.Add(new Person() { Name = "Jack", Age = 456, Books = booksCollectionB });
			string filePath = Path.Combine(
				Environment.GetFolderPath(Environment.SpecialFolder.Personal),
				CollectionNameLowCase + ".xml"
			);
			Console.WriteLine($"Wrting Data to File: {filePath}");
			var xs = new XmlSerializer(typeof(List<Person>), new XmlRootAttribute(CollectionNameLowCase));
			using (var file = File.Create(filePath))
			{
				xs.Serialize(file, people);
			}
		}

		static void DeserializeData()
		{
			var people = new List<Person>();
			var xs = new XmlSerializer(typeof(List<Person>), new XmlRootAttribute(CollectionNameLowCase));
			string filePath = Path.Combine(
				Environment.GetFolderPath(Environment.SpecialFolder.Personal),
				CollectionNameLowCase + ".xml"
			);
			Console.WriteLine($"Reading Data from File: {filePath}");
			using (var file = File.OpenRead(filePath))
			{
				people = (List<Person>)xs.Deserialize(file);
			}
			foreach (var person in people)
			{
				Console.Write($"Name: {person.Name}, Age: {person.Age}, Books ({person.Books.Length}): ");
				foreach (var book in person.Books)
				{
					Console.Write($"\"{book.Title}\" ");
				}
				Console.WriteLine();
			}
		}
	}
}
