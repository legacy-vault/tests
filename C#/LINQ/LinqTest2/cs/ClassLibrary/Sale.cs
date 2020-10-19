using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace ClassLibrary
{
	[Table("Sale")]
	public class Sale
	{
		[Required]
		[DatabaseGenerated(DatabaseGeneratedOption.Identity)]
		public int Id { get; set; }

		[Required]
		public int ProductId { get; set; }

		[Required]
		public int ProductQuantity { get; set; }

		[Required]
		[DatabaseGenerated(DatabaseGeneratedOption.Computed)]
		public DateTime Time { get; set; }
	}
}
