using Microsoft.AspNetCore.Http;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace HelloApp
{
	public class ErrorHandlingMiddleware
	{
		private RequestDelegate nextMiddleware;
		public ErrorHandlingMiddleware(RequestDelegate nextMiddleware)
		{
			this.nextMiddleware = nextMiddleware;
		}
		public async Task InvokeAsync(HttpContext context)
		{
			await nextMiddleware.Invoke(context);

			if (context.Response.StatusCode == 403)
			{
				await context.Response.WriteAsync("403: Access Denied");
			}
			else if (context.Response.StatusCode == 404)
			{
				await context.Response.WriteAsync("404: Not Found");
			}
			else if (context.Response.StatusCode == 400)
			{
				await context.Response.WriteAsync("400: Bad Request");
			}
			else if (context.Response.StatusCode != 200)
			{
				await context.Response.WriteAsync($"Error {context.Response.StatusCode}");
			}
		}
	}
}
