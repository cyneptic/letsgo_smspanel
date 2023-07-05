package ports

type QueueContract interface {
	Publisher(sender string, message string, recevier []string)
}
