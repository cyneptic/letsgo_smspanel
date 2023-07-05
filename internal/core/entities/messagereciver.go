package entities

type MessageReciver struct {
	UserID   string `json:"user_id"`
	Content  string `json:"content"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
}
