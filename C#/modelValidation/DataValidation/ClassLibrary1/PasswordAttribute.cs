using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;

namespace PersonWithValidation
{
	// Password Field Validation Attribute.
	// String must conform with Security Rules.
	class PasswordAttribute : ValidationAttribute
	{
		public static readonly int MinLength = 8;
		public override bool IsValid(object value)
		{
			if (value == null) return false;
			string password = value.ToString();

			if (password.Length < MinLength)
			{
				this.ErrorMessage = "The Password field is too short";
				return false;
			}
			bool hasNumber = false;
			bool hasLowCaseLetter = false;
			bool hasHighCaseLetter = false;
			for (int i = 0; i < password.Length; i++)
			{
				char c = password[i];
				if (char.IsNumber(c)) { hasNumber = true; continue; }
				if (char.IsLetter(c))
				{
					if (char.IsLower(c)) { hasLowCaseLetter = true; continue; }
					if (char.IsUpper(c)) { hasHighCaseLetter = true; continue; }
				}
			}
			if ((hasNumber) && (hasLowCaseLetter) && (hasHighCaseLetter))
			{
				return true;
			}
			else
			{
				string resume = "The Password field violates the security rules: ";
				List<string> violations = new List<string>();
				if (!hasNumber) { violations.Add("no number is present"); }
				if (!hasLowCaseLetter) { violations.Add("no lower case letter is present"); }
				if (!hasHighCaseLetter) { violations.Add("no upper case letter is present"); }
				this.ErrorMessage = resume + string.Join(", ", violations) + ".";
				return false;
			}
		}
	}
}
