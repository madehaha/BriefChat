package request

type Register struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type Login struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type MessageReq struct {
	Receiver string `json:"receiver" gorm:"column:receiver;not null" binding:"required"`
	Txt      string `json:"txt" gorm:"column:txt"`
	Jpg      string `json:"jpg" gorm:"column:jpg"`
}

type Info struct {
	Name string `json:"name" binding:"required,max=10"`
}
