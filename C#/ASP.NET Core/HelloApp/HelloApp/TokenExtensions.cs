using HelloApp;
using Microsoft.AspNetCore.Builder;

public static class TokenExtensions
{
	public static IApplicationBuilder UseToken(this IApplicationBuilder builder, string correctToken)
	{
		return builder.UseMiddleware<TokenMiddleware>(correctToken);
	}
}
