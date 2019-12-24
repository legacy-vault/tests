using System;
using System.Reflection;

namespace TypeReflection
{
	class Program
	{
		static void Main(string[] args)
		{
			Person p = new Person { Name = "John", Age = 123 };
			Type personType = typeof(Person);
			personType = p.GetType();

			MemberFilter searchFilter = (MemberInfo memberInfo, object filterCriteria) => true; // Show Everything.
			MemberInfo[] members = personType.FindMembers(
				MemberTypes.All,
				BindingFlags.Instance | BindingFlags.Public | BindingFlags.NonPublic,
				searchFilter,
				"");

			Console.WriteLine("{0} Members are found.", members.Length);
			FieldInfo fieldInfo;
			PropertyInfo propertyInfo;
			string publicityType;
			foreach (var member in members)
			{
				switch (member.MemberType)
				{
					case MemberTypes.Field:
						fieldInfo = member as FieldInfo;
						if (fieldInfo.IsPublic) { publicityType = "public"; } else { publicityType = "hidden"; }
						Console.WriteLine($"  {member.Name,32} ({publicityType} {member.MemberType}, {fieldInfo.FieldType.FullName})");
						break;

					case MemberTypes.Property:
						propertyInfo = member as PropertyInfo;
						Console.WriteLine($"  {member.Name,32} ({member.MemberType}, {propertyInfo.PropertyType.FullName})");
						break;

					default:
						Console.WriteLine($"  {member.Name,32} ({member.MemberType})");
						break;
				}
			}
		}
	}

	class Person
	{
		public string Name { get; set; }
		public int Age { get; set; }
		public int X;
		public int Y;
		private string Secret;
	}
}
