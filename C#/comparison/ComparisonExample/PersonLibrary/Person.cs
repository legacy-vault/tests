using System;

namespace PersonLibrary
{
	// A Person.
	public class Person : ICloneable, IComparable<Person>, IEquatable<Person>
	{
		public string Name { get; set; }
		public int Age { get; set; }

		public Person(string name, int age)
		{
			Name = name;
			Age = age;
		}
		public Person(string name) : this(name, 0) { }
		public Person(int age) : this(String.Empty, age) { }

		// ICloneable Methods.
		public object Clone()
		{
			return new Person(this.Name, this.Age);
		}

		// IComparable<Person> Methods.
		public int CompareTo(Person that)
		{
			// By definition, any object compares greater than (or follows) null, and two null references compare equal to each other.
			// https://docs.microsoft.com/en-us/dotnet/api/system.icomparable.compareto?redirectedfrom=MSDN&view=netframework-4.8#remarks 
			if (that == null) return 1;

			// Compare by Name.
			int comparisonByName = string.Compare(this.Name, that.Name, System.StringComparison.CurrentCultureIgnoreCase);
			if (comparisonByName != 0) return comparisonByName;

			// Names are equal, compare by Age.
			int comparisonByAge = this.Age.CompareTo(that.Age);
			if (comparisonByAge != 0) return comparisonByAge;

			// Total Equality.
			return 0;
		}

		// IEquatable<Person> Methods.
		public bool Equals(Person that)
		{
			return (this.CompareTo(that) == 0);
		}
	}
}
