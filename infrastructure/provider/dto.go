package provider

type Message struct {
	Sender    string   `json:"sender"`
	Receivers []string `json:"receivers"`
	Content   string   `json:"content"`
}
