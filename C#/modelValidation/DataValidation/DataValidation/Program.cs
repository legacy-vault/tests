using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using PersonWithValidation;

namespace DataValidation
{
	class Program
	{
		static void Main(string[] args)
		{
			Person p1 = new Person
			{
				Name = "",
				Age = 1,
				PhoneNumber = "what?",
				EmailAddress = "junky",
				CreditCardNumber = "empty",
				Website = "Website is not here",
				UserName = "",
				Password = ""
			};
			CheckPrintPerson(p1);

			p1.Name = "John";
			p1.Age = 123;
			p1.PhoneNumber = "9999999999999999999999";
			p1.EmailAddress = "xxx@yyy@zzz";
			p1.CreditCardNumber = "123";
			p1.Website = "http://localhost";
			p1.UserName = "a";
			p1.Password = "a";
			CheckPrintPerson(p1);

			p1.PhoneNumber = "+0 (000) 111-22-33";
			p1.EmailAddress = "john@example.org";
			p1.CreditCardNumber = "1111-2222-3333-4444";
			p1.UserName = "****";
			p1.Password = "..................";
			CheckPrintPerson(p1);

			p1.UserName = "ZZZZ";
			p1.Password = "abcdefgHIJKLMNOP1234567";
			CheckPrintPerson(p1);

			p1.UserName = "john";
			CheckPrintPerson(p1);
		}

		static void CheckPrintPerson(Person p)
		{
			if (!(p.IsValid()))
			{
				// Short Check.
				Console.WriteLine("Person is not valid.");
				List<ValidationResult> errors;

				// Long Check.
				p.Validate(out errors);
				foreach (var error in errors)
				{
					Console.WriteLine($"  {error.ErrorMessage}");
				}

				Console.WriteLine();
			}
			else
			{
				Console.WriteLine("Person is valid.");
			}
		}
	}
}
