using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace ClassLibrary
{
	[Table("Product")]
	public class Product
	{
		[Required]
		[DatabaseGenerated(DatabaseGeneratedOption.Identity)]
		public int Id { get; set; }

		[Required]
		[MaxLength(45)]
		public string Name { get; set; }

		public List<Sale> Sales { get; set; }
	}
}
