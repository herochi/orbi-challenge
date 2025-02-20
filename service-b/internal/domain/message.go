package domain

type Message struct {
	To      string `json:"userId"`
	Message string `json:"message"`
}
