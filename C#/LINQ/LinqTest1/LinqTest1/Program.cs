using System;
using System.Collections.Generic;
using System.Linq;

namespace LinqTest1
{
	class Program
	{
		static void Main(string[] args)
		{
			List<Pet> pets = new List<Pet>
			{
				new Pet{ Age=1,Name="Tobi",Size=10},
				new Pet{ Age=2,Name="Gavka",Size=12},
				new Pet{ Age=1,Name="Alice",Size=11},
				new Pet{ Age=3,Name="Yoyo",Size=12},
				new Pet{ Age=1,Name="Dodo",Size=10},
				new Pet{ Age=2,Name="Zozo",Size=13},
				new Pet{ Age=2,Name="Bob",Size=12},
				new Pet{ Age=3,Name="Xaxa",Size=10},
			};

			// Single Grouping.
			var petsQuery = from pet in pets
							orderby pet.Age descending
							group pet by pet.Age;
			foreach (var petsByAge in petsQuery)
			{
				Console.Write($"Age {petsByAge.Key}: ");
				foreach (var pet in petsByAge)
				{
					Console.Write($"{pet.Name} (S-{pet.Size}) ");
				}
				Console.WriteLine();
			}

			// Double Grouping.
			Console.WriteLine();
			var petsQuery2 = from pet in pets
							 orderby pet.Age descending
							 group pet by pet.Age into petsByAge
							 from petsBySize in
							 (
								 from pets in petsByAge
								 orderby pets.Size ascending
								 group pets by pets.Size
							 )
							 group petsBySize by petsByAge.Key;
			foreach (var petsByAgeSize in petsQuery2)
			{
				Console.Write($"Age {petsByAgeSize.Key}: ");
				foreach (var petsByAge in petsByAgeSize)
				{
					Console.Write($"\r\n\tSize {petsByAge.Key}: ");
					foreach (var pet in petsByAge)
					{
						Console.Write($"{pet.Name} ");
					}
				}
				Console.WriteLine();
			}

			// Double Grouping (alternative).
			Console.WriteLine();
			var petsQuery3 = from pet in pets
							 group pet by new { pet.Age, pet.Size } into tmp
							 orderby tmp.Key.Age descending, tmp.Key.Size ascending
							 select tmp;
			foreach (var x in petsQuery3)
			{
				Console.Write($"{x.Key}");
				foreach (var y in x)
				{
					Console.Write($"\t[{y.Age} {y.Size} {y.Name}]");
				}
				Console.WriteLine();
			}
		}
	}

	class Pet
	{
		public int Age;
		public string Name;
		public int Size;
	}
}
