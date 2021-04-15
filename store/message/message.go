package message

type Message struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	Message  string `json:"message"`
}
