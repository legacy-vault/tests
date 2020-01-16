using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace HelloMvcApp.Models
{
	public class StoreDbContext : DbContext
	{
		public DbSet<Phone> Phones { get; set; }
		public DbSet<Order> Orders { get; set; }

		public StoreDbContext(DbContextOptions<StoreDbContext> options)
			: base(options)
		{
			Database.EnsureCreated();
		}
	}
}
