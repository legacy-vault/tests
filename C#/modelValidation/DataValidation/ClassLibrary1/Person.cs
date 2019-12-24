using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;

namespace PersonWithValidation
{
	public class Person
	{
		[Required]
		[StringLength(255, MinimumLength = 1)]
		public string Name { get; set; }

		[Required]
		[Range(18, 255)]
		public int Age { get; set; }

		[Required]
		[Phone]
		public string PhoneNumber { get; set; }

		[Required]
		[EmailAddress]
		public string EmailAddress { get; set; }

		[Required]
		[CreditCard]
		public string CreditCardNumber { get; set; }

		[Required]
		[Url]
		public string Website { get; set; }

		[Required]
		[UserName]
		public string UserName { get; set; }

		[Required]
		[Password]
		public string Password { get; set; }

		// Validation without verbose Results.
		public bool IsValid()
		{
			var context = new ValidationContext(this);
			return Validator.TryValidateObject(this, context, new List<ValidationResult>(), true);
		}

		// Validation with verbose Results.
		public bool Validate(out List<ValidationResult> errors)
		{
			errors = new List<ValidationResult>();
			var context = new ValidationContext(this);
			return Validator.TryValidateObject(this, context, errors, true);
		}
	}
}
