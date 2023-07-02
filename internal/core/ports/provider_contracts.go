package ports

type QueueProviderContract interface {
	SendMessage(sender, msg string, receivers interface{}) (isSuccessful bool)
}

type MessageProvider interface {
	Publisher(sender, msg string, receivers interface{})
}
