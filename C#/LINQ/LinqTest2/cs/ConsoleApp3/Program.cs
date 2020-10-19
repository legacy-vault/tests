using ClassLibrary;
using Microsoft.EntityFrameworkCore;
using System;
using System.Linq;
using System.Threading;

namespace ConsoleApp3
{
	class Program
	{
		static void Main(string[] args)
		{
			try
			{
				// Get Data from the Database.
				GetPrintSomeData();

				// Add a Product and its Sale to the Database.
				Console.Write("To proceed with Addition of a new Product, press the 'Y' Key...");
				ConsoleKeyInfo key = Console.ReadKey();
				if (key.Key != ConsoleKey.Y) return;
				Console.WriteLine();
				AddSomeData();
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}, {ex.InnerException?.Message}");
			}
		}

		static void GetPrintSomeData()
		{
			using (ProductSalesContext db = new ProductSalesContext())
			{
				GetPrintProducts(db);
				GetPrintSales(db);
				GetPrintProductWithSales(db);
			}
		}

		static void GetPrintProducts(ProductSalesContext db)
		{
			// Get Products.
			Console.WriteLine("Products");
			var products = db.Products
				.Where(p => p.Name.Length > 0)
				.OrderBy(p => p.Name)
				.ToList();
			foreach (var product in products)
			{
				Console.WriteLine($"Product: Id={product.Id} Name='{product.Name}'.");
			}
			Console.WriteLine();
		}

		static void GetPrintSales(ProductSalesContext db)
		{
			// Get Sales.
			Console.WriteLine("Sales");
			var sales = db.Sales
				.Where(s => s.ProductQuantity > 0)
				.OrderBy(s => s.Time)
				.ToList();
			foreach (var sale in sales)
			{
				Console.WriteLine($"Sale: Id={sale.Id} ProductId={sale.ProductId} ProductQuantity={sale.ProductQuantity} Time={sale.Time}.");
			}
			Console.WriteLine();
		}

		static void GetPrintProductWithSales(ProductSalesContext db)
		{
			// Get Products with their Sales.
			Console.WriteLine("Product with Sales Statistics");
			var productsWithSales = db.Products
				.Include(p => p.Sales) // ORM Behaviour: Make additional Queries for related Sales.
				.Where(p => p.Name.Length > 0)
				.OrderBy(p => p.Name)
				.ToList();
			foreach (var product in productsWithSales)
			{
				Console.WriteLine($"Product: Id={product.Id} Name='{product.Name}' Sales:{product.Sales.Count}.");
			}
			Console.WriteLine();
		}

		static void AddSomeData()
		{
			using (var db = new ProductSalesContext())
			{
				// Add a Product.
				AddProduct(db);

				// Show the Results.
				GetPrintProductWithSales(db);
				GetPrintSales(db);
			}
		}

		static void AddProduct(ProductSalesContext db)
		{
			// Add a Product.
			Console.Write("Adding a Product...\r\nEnter the Product's Name: ");
			string newProductName = Console.ReadLine();
			Console.Write("Enter the Count of sold Products: ");
			string newProductSoldCountStr = Console.ReadLine();
			int newProductSoldCount = Convert.ToInt32(newProductSoldCountStr, Thread.CurrentThread.CurrentCulture);
			var newProduct = new Product { Name = newProductName };
			db.Products.Add(newProduct);
			db.SaveChanges();

			// Add a Sale.
			// Notes:
			// The ID of the inserted Product is updated in the C# Object 
			// (synchronized with the Database) automagically due to the LINQ ORM Behaviour!
			var newSale = new Sale { ProductId = newProduct.Id, ProductQuantity = newProductSoldCount };
			db.Sales.Add(newSale);
			db.SaveChanges();
		}
	}
}
