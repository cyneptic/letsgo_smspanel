package ports

type QueueProviderContract interface {
	SendMessage(sender, msg string, receiver []string) (isSuccessful bool)
}

type MessageProvider interface {
	Publisher(sender, msg string, receivers []string)
}
