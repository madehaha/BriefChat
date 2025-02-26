package request

type Register struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type Login struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
