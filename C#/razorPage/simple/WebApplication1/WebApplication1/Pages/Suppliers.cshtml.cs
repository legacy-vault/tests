using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using NorthwindEntitiesLib;

namespace WebApplication1.Pages
{
	public class SuppliersModel : PageModel
	{
		public IEnumerable<string> Suppliers { get; set; }

		[BindProperty]
		public Supplier Supplier { get; set; }

		public void OnGet()
		{
			ViewData["Title"] = "Northwind Web Site - Suppliers";
			Suppliers = new[] { "Alpha Co", "Beta Limited", "Gamma Corp" };
		}

		public IActionResult OnPost()
		{
			ViewData["Title"] = "Northwind Web Site - Supplier Addition";
			Suppliers = new[] { "Alpha Co", "Beta Limited", "Gamma Corp" };
			if (ModelState.IsValid)
			{
				ViewData["Hint"] = "You have tried to add a new supplier: " +
					Supplier.CompanyName;
			}
			return Page();
		}
	}
}
