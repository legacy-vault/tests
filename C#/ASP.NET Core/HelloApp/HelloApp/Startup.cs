using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.FileProviders;
using Microsoft.Extensions.Hosting;

namespace HelloApp
{
	public class Startup
	{
		IWebHostEnvironment environment;
		public Startup(IWebHostEnvironment env)
		{
			environment = env;
		}

		// This method gets called by the runtime. Use this method to add services to the container.
		// For more information on how to configure your application, visit https://go.microsoft.com/fwlink/?LinkID=398940
		public void ConfigureServices(IServiceCollection services)
		{
		}

		// This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
		public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
		{
			// Environment Check.
			if (env.IsProduction())
			{
				//app.UseHsts();
				app.UseExceptionHandler("/error");
			}
			else
			{
				if (env.IsDevelopment())
				{
					app.UseDeveloperExceptionPage();
				}
				else if (env.IsEnvironment("Test"))
				{
					app.Run(async (context) =>
					{
						await context.Response.WriteAsync("Test Environment is not yet supported.");
					});
				}
				else
				{
					app.Run(async (context) =>
					{
						await context.Response.WriteAsync("Unknown Environment.");
					});
				}
			}
			//app.UseHttpsRedirection();

			//app.Map("/system", SystemHandler);
			app.Map("/error", ap => ap.Run(async context =>
			{
				await context.Response.WriteAsync("Error");
			}));

			app.Use(async (context, next) =>
			{
				// Do Something...
				await next.Invoke();
			});

			// Status Codes (Errors).
			//app.UseMiddleware<ErrorHandlingMiddleware>();
			app.UseStatusCodePages();
			//app.UseStatusCodePages("text/plain", "Error. Status code : {0}");
			//app.UseStatusCodePagesWithRedirects("/error?code={0}");
			//app.UseStatusCodePagesWithReExecute("/error", "?code={0}");

			//app.UseMiddleware<TokenMiddleware>();
			app.UseToken("CS");

			app.UseRouting();

			// Default Files.
			DefaultFilesOptions defaultFilesOptions = new DefaultFilesOptions();
			defaultFilesOptions.DefaultFileNames.Add("default.xhtml");
			defaultFilesOptions.DefaultFileNames.Add("index.xhtml");
			app.UseDefaultFiles();

			// Directory Browser.
			string pathToPublicFolder = Path.Combine(Directory.GetCurrentDirectory(), "Content", "public");
			string publicFolderAlias = "/files";
			DirectoryBrowserOptions directoryBrowserOptions = new DirectoryBrowserOptions()
			{
				FileProvider = new PhysicalFileProvider(pathToPublicFolder),
				RequestPath = new PathString(publicFolderAlias)
			};
			app.UseDirectoryBrowser(directoryBrowserOptions);

			// Static Files.
			app.UseStaticFiles(); // [/] -> [/www].
			StaticFileOptions staticFileOptions = new StaticFileOptions()
			{
				FileProvider = new PhysicalFileProvider(pathToPublicFolder),
				RequestPath = new PathString(publicFolderAlias)
			};
			app.UseStaticFiles(staticFileOptions); // [/files] -> [/public].

			//app.UseFileServer(enableDirectoryBrowsing: true);

			app.UseEndpoints(endpoints =>
			{
				endpoints.MapGet("/", async context =>
				{
					int x = 0;
					x = x / x; // => Exception.
					await context.Response.WriteAsync($"Hello World! Running at {environment.ContentRootPath}.");
				});
			});

			/*
			app.Run(async (context) =>
			{
				await context.Response.WriteAsync("End of the Line!");
			});
			*/
		}
	}
}
