using System;
using System.IO;
using System.Net;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace HttpServerLibrary
{
	public class Server
	{
		public string Protocol { get; }
		public string Host { get; }
		public string Port { get; }

		private string[] uriPrefixes;
		public string[] GetUriPrefixes()
		{
			return uriPrefixes;
		}

		// Settings.
		private bool startupLogIsOn = false;
		private bool requestLoggingIsOn = false;
		private bool highLoadEmulationIsOn = true;

		// Listener & Work Status.
		private bool isStarted = false;
		private bool newRequestsAreAccepted = false;
		private int activeRequestsCount = -1;
		private static HttpListener listener;
		private Task listenerTask;

		// Shutdown.
		private bool shutdownIsNeeded = false;
		private Task shutdownTask;

		public Server(
			string protocol,
			string host,
			string port,
			bool startupLogIsOn = false,
			bool requestLoggingIsOn = false,
			bool highLoadEmulationIsOn = false)
		{
			Protocol = protocol;
			Host = host;
			Port = port;
			uriPrefixes = new string[]
			{
				string.Format(
					Thread.CurrentThread.CurrentCulture,
					"{0}://{1}:{2}{3}",
					Protocol, Host, Port, "/")
			};
			this.startupLogIsOn = startupLogIsOn;
			this.requestLoggingIsOn = requestLoggingIsOn;
			this.highLoadEmulationIsOn = highLoadEmulationIsOn;
		}

		// Starts the Server.
		// Returns 'false' on Failure.
		public bool Start()
		{
			try
			{
				if (isStarted)
				{
					throw new Exception("Server is already started.");
				}

				listener = StartHttpListener(uriPrefixes);
				isStarted = true;
				activeRequestsCount = 0;
				newRequestsAreAccepted = true;
				listenerTask = Listen();
				return true;
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
				return false;
			}
		}

		// Stops the Server synchronously.
		public void Stop()
		{
			if (!isStarted)
			{
				throw new Exception("Server is already stopped.");
			}
			if (shutdownIsNeeded)
			{
				throw new Exception("Shutdown has already been requested.");
			}

			shutdownTask = Task.Run(() => Shutdown());
			listenerTask.Wait();
		}

		// Writes the Message to the Console.
		private void LogMessage(string message)
		{
			Console.WriteLine(message);
		}

		// Starts the HttpListener.
		// Returns null on Failure.
		private HttpListener StartHttpListener(string[] uriPrefixes)
		{
			// Check.
			if (!HttpListener.IsSupported)
			{
				throw new Exception("The Operating System does not support the HttpListener Class.");
			}
			if ((uriPrefixes == null) || (uriPrefixes.Length == 0))
			{
				throw new Exception("URI Prefixes are not set.");
			}

			// Listener Creation.
			HttpListener listener = new HttpListener();
			foreach (string uriPrefix in uriPrefixes)
			{
				listener.Prefixes.Add(uriPrefix);
				if (startupLogIsOn)
				{
					LogMessage($"Listening at {uriPrefix}...");
				}
			}
			listener.Start();
			return listener;
		}

		// Listens to the incoming Connections and processes them.
		private async Task Listen()
		{
			try
			{
				while (newRequestsAreAccepted)
				{
					await AcceptProcessRequest();
				}
				WaitForShutdown();
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}
		}

		// Accepts an incoming Connection and processes it asynchronously.
		private async Task AcceptProcessRequest()
		{
			try
			{
				// Notes:
				// Listener is stopped by the 'Shutdown' Method while 'Listen' Method is 
				// blocked and can not exit its own Loop waiting for the next Connection.
				if (listener == null) return;
				HttpListenerContext context = await listener.GetContextAsync();
				if (newRequestsAreAccepted)
				{
					Task.Run(() => ProcessRequestWithCounter(context));
				}
				else
				{
					if (requestLoggingIsOn)
					{
						LogMessage("A new Request received in the Shutdown Stage has been rejected.");
					}
				}
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}
		}

		// Waits for the Shutdown Task to complete.
		private void WaitForShutdown()
		{
			if (!shutdownIsNeeded)
			{
				throw new Exception("Shutdown has not been requested, but the main Work Loop has stopped.");
			}
			if (shutdownTask == null)
			{
				throw new Exception("Shutdown Task does not exist.");
			}

			shutdownTask.Wait();
		}

		// Processes the Request, modifies the active Requests Counter.
		private void ProcessRequestWithCounter(HttpListenerContext context)
		{
			try
			{
				Interlocked.Increment(ref activeRequestsCount);
				ProcessRequest(context);
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}
			finally
			{
				Interlocked.Decrement(ref activeRequestsCount);
			}
		}

		// Processes the Request.
		private void ProcessRequest(HttpListenerContext context)
		{
			HttpListenerRequest request = context.Request;
			HttpListenerResponse response = context.Response;

			// Read the Request.
			string requestBody;
			string requestedResource = $"{request.HttpMethod} {request.Url}";
			if (requestLoggingIsOn)
			{
				LogMessage(requestedResource);
			}
			Stream inputStream = request.InputStream;
			Encoding encoding = request.ContentEncoding;
			using StreamReader reader = new StreamReader(inputStream, encoding);
			requestBody = reader.ReadToEnd();

			// Simulation of some busy and laggy Infrastructure.
			if (highLoadEmulationIsOn)
			{
				Thread.Sleep(10_000);
			}

			// Write the Response.
			response.StatusCode = (int)HttpStatusCode.OK;
			using Stream outputStream = response.OutputStream;
			StringBuilder outputText = new StringBuilder();
			outputText.AppendLine("<!DOCTYPE html>");
			outputText.AppendLine("<html>");
			outputText.AppendLine("<head>");
			outputText.AppendLine($"\t<title>HTTP Server Page Example</title>");
			outputText.AppendLine($"\t<meta charset=\"utf-8\">");
			outputText.AppendLine("</head>");
			outputText.AppendLine("<body>");

			// Document.
			outputText.AppendLine(
				$"You ({request.RemoteEndPoint}) have requested the following Document: '{requestedResource}'.<br />" +
				$"Raw URL is '{request.RawUrl}'.<br />" +
				$"Scheme is '{request.Url.Scheme}', Port is '{request.Url.Port}', Query is '{request.Url.Query}'.<br />");
			outputText.AppendLine("<br />");

			// URL Path Segments.
			outputText.AppendLine($"URL Path Segments:<br />");
			foreach (var urlSegment in request.Url.Segments)
			{
				outputText.AppendLine($"{urlSegment}<br />");
			}
			outputText.AppendLine("<br />");

			// Headers List.
			outputText.AppendLine($"Request Headers:<br />");
			foreach (var header in request.Headers.AllKeys)
			{
				string[] headerValues = request.Headers.GetValues(header);
				foreach (var headerValue in headerValues)
				{
					outputText.AppendLine($"[{header}] = \"{headerValue}\"<br />");
				}
			}
			outputText.AppendLine("</body>");
			outputText.AppendLine("</html>");
			byte[] outputBuffer = Encoding.UTF8.GetBytes(outputText.ToString());
			outputStream.Write(outputBuffer, 0, outputBuffer.Length);
		}

		// Stops accepting incoming Requests, waits for all active Requests to be completed, stops the Listener.
		private void Shutdown()
		{
			try
			{
				newRequestsAreAccepted = false;
				shutdownIsNeeded = true;
				LogMessage("Shutdown has been initiated.");
				GracefulShutdown();
			}
			catch (Exception ex)
			{
				Console.WriteLine($"{ex.GetType()} Exception: {ex.Message}");
			}
		}

		// Waits for all active Requests to be completed, stops the Listener.
		private void GracefulShutdown()
		{
			int activeRequestsCountPrevious = -1;
			if (requestLoggingIsOn)
			{
				LogMessage($"Waiting for active Requests to be processed...");
			}
			while (activeRequestsCount > 0)
			{
				if (activeRequestsCount != activeRequestsCountPrevious)
				{
					if (requestLoggingIsOn)
					{
						Console.Write($"{activeRequestsCount}.");
					}
				}
				else
				{
					if (requestLoggingIsOn)
					{
						Console.Write(".");
					}
				}
				activeRequestsCountPrevious = activeRequestsCount;
				Thread.Sleep(1000);
			}
			if (requestLoggingIsOn)
			{
				LogMessage("");
			}
			if (activeRequestsCount == 0)
			{
				if (requestLoggingIsOn)
				{
					LogMessage($"All active Requests have been processed.");
				}
			}

			listener.Stop();
			isStarted = false;
		}
	}
}
