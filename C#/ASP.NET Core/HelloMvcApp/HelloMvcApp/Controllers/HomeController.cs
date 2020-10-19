using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using HelloMvcApp.Models;

namespace HelloMvcApp.Controllers
{
	public class HomeController : Controller
	{
		private readonly ILogger<HomeController> logger;
		private StoreDbContext db;

		public HomeController(
			ILogger<HomeController> logger,
			StoreDbContext dbContext)
		{
			this.logger = logger;
			this.db = dbContext;
		}

		[HttpGet]
		// GET /Home/Buy/3.
		public IActionResult Buy(int? id)
		{
			if (id == null) return RedirectToAction("Index");
			ViewBag.PhoneId = id;
			
			return View();
		}

		[HttpPost]
		// POST /Home/Buy.
		public string Buy(Order order)
		{
			// Add an Order to DB.
			db.Orders.Add(order);
			db.SaveChanges();

			return "Dear, " + order.User + "! We thank you for your purchase.";
		}

		[ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
		public IActionResult Error()
		{
			return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
		}

		public IActionResult Index()
		{
			return View(
				db.Phones.ToList()
			);
		}

		public IActionResult Privacy()
		{
			return View();
		}
	}
}
