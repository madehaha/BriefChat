package entity

type Chat struct {
	Id       int64  `gorm:"column:id; not null; primaryKey; autoIncrement"`
	Account  string `gorm:"column:account;not null; unique"`
	Password string `gorm:"column:password;not null"`
	SelfInfo Info   `gorm:"embedded"`
}

type Info struct {
	Name   string `json:"name" gorm:"column:name;" example:"MOMO"`
	Avatar string `json:"avatar" gorm:"column:avatar;" example:"example.com/1919810.jpg"`
}

type FriendInfo struct {
	Name    string `json:"name" gorm:"column:name;" example:"MOMO"`
	Avatar  string `json:"avatar" gorm:"column:avatar;" example:"example.com/1919810.jpg"`
	Account string `json:"account" gorm:"column:account;not null; unique"`
}
