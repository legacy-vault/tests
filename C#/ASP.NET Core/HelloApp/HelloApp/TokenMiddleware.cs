using Microsoft.AspNetCore.Http;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace HelloApp
{
	public class TokenMiddleware
	{
		private readonly RequestDelegate nextMiddleware;
		private readonly string correctToken;
		public const string TokenQueryParameter = "token";

		public TokenMiddleware(RequestDelegate nextMiddleware, string correctToken)
		{
			this.nextMiddleware = nextMiddleware;
			this.correctToken = correctToken;
		}

		public async Task InvokeAsync(HttpContext context)
		{
			// Get the Token.
			if (!context.Request.Query.ContainsKey(TokenQueryParameter))
			{
				context.Response.StatusCode = 400;
				return;
			}
			var token = context.Request.Query[TokenQueryParameter];

			// Check the Token.
			if (token != correctToken)
			{
				context.Response.StatusCode = 403;
				return;
			}

			// Success.
			await nextMiddleware.Invoke(context);
		}
	}
}
