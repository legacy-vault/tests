using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Infrastructure;
//using MySql.Data.EntityFrameworkCore;
using Pomelo.EntityFrameworkCore.MySql;
using System;

namespace ClassLibrary
{
	public class ProductSalesContext : DbContext
	{
		public ProductSalesContext() : base() { }
		public ProductSalesContext(DbContextOptions<ProductSalesContext> options) : base(options) { }


		public DbSet<Product> Products { get; set; }
		public DbSet<Sale> Sales { get; set; }

		protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
		{
			string dsn = "server=127.0.0.1;uid=test;pwd=test;database=test";
			Action<MySqlDbContextOptionsBuilder> mySqlOptionsAction = o => o.EnableRetryOnFailure().CommandTimeout(60);
			optionsBuilder.UseMySql(dsn, mySqlOptionsAction);
		}
	}
}
