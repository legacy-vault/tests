using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using PersonLibrary;

namespace ComparisonExample
{
	class Program
	{
		static void Main(string[] args)
		{
			// Simple Sort Test.
			Console.WriteLine("Simple Sort Test");
			List<int?> l = new List<int?> { 10, 1, null, 5, 3 };
			l.Sort();
			foreach (var item in l) Console.Write($"[{item}] ");
			Console.WriteLine();
			Console.WriteLine();

			// Create an original Copy of a List.
			Person personJohn10 = new Person("John", 10);
			Person personJohn5 = new Person("John", 5);
			Person personJack20 = new Person("Jack", 20);
			Person personJuliette3 = new Person("Juliette", 3);
			List<Person> people = new List<Person>();
			people.Add(personJuliette3);
			people.Add(null);
			people.Add(personJohn10);
			people.Add(personJack20);
			people.Add(personJohn5);
			Console.WriteLine("Original List");
			PrintPeopleCollection(people);

			// Make a List Copy, and try a simple Sorting.
			// Seems that the default Sorter skips null Items.
			var peopleForTest = people.ToList();
			peopleForTest.Sort();
			Console.WriteLine("List sorted with the default Sorter");
			PrintPeopleCollection(peopleForTest);

			// Make a List Copy, and try an advanced Sorting.
			var peopleForTest2 = people.ToList();
			peopleForTest2.Sort(new PersonComparer());
			Console.WriteLine("List sorted with the advanced Sorter");
			PrintPeopleCollection(peopleForTest2);

			// Make another List Copy into an Array, and try an advanced Sorting.
			var peopleForTest3 = people.ToArray();
			Array.Sort(peopleForTest3, new PersonComparer());
			Console.WriteLine("Sorted Array");
			PrintPeopleCollection(peopleForTest3);

			// Equality Test.
			Person personNone = null;
			Person personNone2 = null;
			Console.WriteLine($"{personJack20 == personNone} {personNone == personNone2}");
		}

		static void PrintPeopleCollection(IEnumerable people)
		{
			Console.Write("People: ");
			foreach (Person p in people)
			{
				if (p != null)
				{
					Console.Write($"[{p.Name}:{p.Age}] ");
				}
				else
				{
					Console.Write("[null] ");
				}
			}
			Console.WriteLine();
			Console.WriteLine();
		}
	}
}
