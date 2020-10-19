using System;

namespace ConsoleChatServerLibrary
{
	public enum MessageType
	{
		None,
		System,
		Text
	}

	public enum MessageSenderType
	{
		None,
		System,
		User
	}

	public class Message
	{
		public static readonly string TimeFormat = "yyyy-MM-dd HH:mm:ss";
		public MessageType Type { get; }
		public MessageSenderType SenderType { get; }
		public User SenderUser { get; }
		public string Text { get; }
		public DateTime Time { get; }

		public Message(
			string text,
			MessageSenderType senderType = MessageSenderType.User,
			User senderUser = null,
			MessageType type = MessageType.Text
		)
		{
			Type = type;
			SenderType = senderType;
			SenderUser = senderUser;
			Text = text;
			Time = DateTime.Now;
		}

		// Checks whether the Message is a System Message.
		public bool IsSystem()
		{
			return (Type == MessageType.System);
		}

		// Checks whether the Message is a Text Message.
		public bool IsText()
		{
			return (Type == MessageType.Text);
		}

		// Checks whether the Message is from the System.
		public bool IsFromSystem()
		{
			return (SenderType == MessageSenderType.System);
		}

		// Checks whether the Message is from a User.
		public bool IsFromUser()
		{
			return (SenderType == MessageSenderType.User);
		}

		// Creates a Text Message sent from the System.
		public static Message TextMessageFromSystem(string text)
		{
			return new Message(text, MessageSenderType.System);
		}

		// Creates a Text Message sent from a User.
		public static Message TextMessageFromUser(User user, string text)
		{
			if (user == null) { throw new Exception("Text Message from an unspecified User."); }
			return new Message(text, MessageSenderType.User, user);
		}

		// Creates a System Message sent from a User.
		public static Message SystemMessageFromUser(User user, string text)
		{
			if (user == null) { throw new Exception("System Message from an unspecified User."); }
			return new Message(text, MessageSenderType.User, user, MessageType.System);
		}

		// Creates a User Greeting Message.
		public static Message UserGreeting(User user)
		{
			string text = $"{user.Name}, welcome to the chat!\r\n" +
				"Type '/help' for help or type a message for everyone.";
			return TextMessageFromSystem(text);
		}

		// Creates a Message about User's normal Log out.
		public static Message UserLogOutNormal(User user)
		{
			string text = $"User '{user.Name}' has logged out.";
			return TextMessageFromSystem(text);
		}

		// Creates a Message about User's problematic Log out.
		public static Message UserLogOutWithProblem(User user)
		{
			string text = $"User '{user.Name}' has been logged out.";
			return TextMessageFromSystem(text);
		}

		// Creates a Message about User's Log in.
		public static Message UserLogIn(User user)
		{
			string text = $"User '{user.Name}' has logged in.";
			return TextMessageFromSystem(text);
		}

		// Creates a Text with some Help for the User.
		public static string TextHelp()
		{
			return "Here is the list of available commands:\r\n" +
				$"  {ChatServer.SystemMessageHelp} - Show this help message;\r\n" +
				$"  {ChatServer.SystemMessageMembersList} - Show the list of chat members;\r\n" +
				$"  {ChatServer.SystemMessageExit}, {ChatServer.SystemMessageQuit} or {ChatServer.SystemMessageLogout} - Log out.";
		}
	}
}
