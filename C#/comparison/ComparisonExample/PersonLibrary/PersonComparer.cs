using System.Collections.Generic;

namespace PersonLibrary
{
	// A Person Comparer.
	public class PersonComparer : IComparer<Person>
	{
		public int Compare(Person p1, Person p2)
		{
			if (p1 != null)
			{
				return p1.CompareTo(p2);
			}
			else
			{
				if (p2 == null)
				{
					// Both Objects are null!
					return 0;
				}
				else
				{
					return -1;
				}
			}
		}
	}
}
