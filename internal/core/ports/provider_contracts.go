package ports

type QueueContract interface {
	Publisher(sender, message, recevier string)
}
