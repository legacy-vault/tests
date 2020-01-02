using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace HttpClientExample
{
	class Program
	{
		static void Main()
		{
			try
			{
				HttpClientUsage();
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}
		}

		// HttpClient is intended to be instantiated once per application, rather than per-use.
		// https://docs.microsoft.com/en/dotnet/api/system.net.http.httpclient
		static readonly HttpClient httpClient = new HttpClient();

		static void HttpClientUsage()
		{
			InitHttpClient();

			string[] pageUrls = new string[] {
				"https://www.yandex.ru",
				"https://www.mail.ru",
				"https://www.bing.com",
				"https://microsoft.com"
			};
			List<Task<int>> tasks = new List<Task<int>>();
			foreach (string pageUrl in pageUrls)
			{
				Task<int> t = Task.Run(() => GetPageSizeAsync(pageUrl));
				tasks.Add(t);
			}
			Task tasksController = Task.WhenAll(tasks);
			tasksController.Wait();
			for (int i = 0; i < pageUrls.Length; i++)
			{
				Console.WriteLine($"{pageUrls[i]}, Size={tasks[i].Result}");
			}

			httpClient.Dispose();
		}

		static void InitHttpClient()
		{
			httpClient.Timeout = new TimeSpan(0, 0, 60); // 1 Minute.
			httpClient.DefaultRequestHeaders.Add("X-CustomHeader", "CustomHeaderValue");
		}

		static async Task<int> GetPageSizeAsync(string pageAddress)
		{
			try
			{
				Uri requestUri = new Uri(pageAddress);
				using HttpResponseMessage response = await httpClient.GetAsync(requestUri);
				response.EnsureSuccessStatusCode();
				string responseBody = await response.Content.ReadAsStringAsync();
				return responseBody.Length;
			}
			catch (HttpRequestException e)
			{
				Console.WriteLine(e.Message);
				return -1;
			}
		}
	}
}
