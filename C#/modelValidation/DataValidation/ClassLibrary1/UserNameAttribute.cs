using System.ComponentModel.DataAnnotations;

namespace PersonWithValidation
{
	// UserName Field Validation Attribute.
	// String must contain only Letters of low Case.
	public class UserNameAttribute : ValidationAttribute
	{
		public static readonly int MinLength = 4;
		public override bool IsValid(object value)
		{
			if (value == null) return false;
			string userName = value.ToString();

			if (userName.Length < MinLength)
			{
				this.ErrorMessage = "The UserName field is too short.";
				return false;
			}
			for (int i = 0; i < userName.Length; i++)
			{
				char c = userName[i];
				if (!char.IsLetter(c))
				{
					this.ErrorMessage = $"The UserName field contains a '{c}' symbol which is not a letter.";
					return false;
				}
				if (char.IsUpper(c))
				{
					this.ErrorMessage = $"The UserName field contains a '{c}' symbol which is not a low case letter.";
					return false;
				}
			}
			return true;
		}
	}
}
