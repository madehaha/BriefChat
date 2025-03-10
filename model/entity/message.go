package entity

type Message struct {
	Sender   string `json:"sender" gorm:"column:sender;not null" binding:"required"`
	Receiver string `json:"receiver" gorm:"column:receiver;not null" binding:"required"`
	Txt      string `json:"txt" gorm:"column:txt"`
	Jpg      string `json:"jpg" gorm:"column:jpg"`
	Date     string `json:"date" gorm:"column:date"`
}

type Friend struct {
	Account1 string `json:"account1" gorm:"column:account1"`
	Account2 string `json:"account2" gorm:"column:account2"`
}
