using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using HelloMvcApp.Models;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;

namespace HelloMvcApp
{
	public class Program
	{
		public static void Main(string[] args)
		{
			IHost host = CreateHostBuilder(args).Build();
			if (!InitializeDatabaseWithSampleData(host)) return;
			host.Run();
		}

		public static IHostBuilder CreateHostBuilder(string[] args)
		{
			return Host.CreateDefaultBuilder(args)
				.ConfigureWebHostDefaults(webBuilder =>
				{
					webBuilder.UseStartup<Startup>();
				});
		}

		static private bool InitializeDatabaseWithSampleData(IHost host)
		{
			using (var serviceScope = host.Services.CreateScope())
			{
				IServiceProvider serviceProvider = serviceScope.ServiceProvider;
				ILogger logger;

				// Get the Logger Service.
				try
				{
					logger = serviceProvider.GetRequiredService<ILogger<Program>>();
				}
				catch (Exception ex)
				{
					Console.WriteLine("Can not get the Logger.");
					Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}.");
					return false;
				}

				// Get the Database Context, initialize the Database with sample Data.
				try
				{
					StoreDbContext dbContext = serviceProvider.GetRequiredService<StoreDbContext>();
					SampleData.Initialize(dbContext);
					logger.LogInformation("Database has been initialized with sample data.");
					return true;
				}
				catch (Exception ex)
				{
					logger.LogError(ex, "Database Initialization Error");
					return false;
				}
			}
		}
	}
}
