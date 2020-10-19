using HelloMvcApp.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace HelloMvcApp
{
	public class SampleData
	{
		public static void Initialize(StoreDbContext databaseContext)
		{
			if (!databaseContext.Phones.Any())
			{
				databaseContext.Phones.AddRange(
					new Phone
					{
						Name = "iPhone X",
						Company = "Apple",
						Price = 600
					},
					new Phone
					{
						Name = "Samsung Galaxy Edge",
						Company = "Samsung",
						Price = 550
					},
					new Phone
					{
						Name = "Pixel 3",
						Company = "Google",
						Price = 500
					}
				);
				databaseContext.SaveChanges();
			}
		}
	}
}
