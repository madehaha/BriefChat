package entity

type Message struct {
	Sender   string `json:"sender" binding:"required"`
	Receiver string `json:"receiver" binding:"required"`
	Txt      string `json:"txt"`
	Jpg      string `json:"jpg"`
	Date     string `json:"date"`
}
