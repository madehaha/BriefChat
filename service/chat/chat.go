package chat

import (
	"BriefChat/global"
	"BriefChat/middleware"
	"BriefChat/model/entity"
	"BriefChat/model/request"
	"BriefChat/utils/status"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ChatService struct {
}

func (c *ChatService) Register(r request.Register) error {
	if !errors.Is(global.Global_Db.Where("account = ?", r.Account).First(&entity.Chat{}).Error,
		gorm.ErrRecordNotFound) {
		return status.SameAccountExists
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = global.Global_Db.Create(&entity.Chat{
		Account:  r.Account,
		Password: string(encryptedPassword),
	}).Error
	return err
}
func (c *ChatService) Login(r request.Login) (string, error) {
	var chat = entity.Chat{}
	println(r.Account)
	err := global.Global_Db.Where("account = ?", r.Account).First(&chat).Error
	println(chat.Account)
	if err != nil {
		return "", status.NoAccountExist
	}
	if err3 := bcrypt.CompareHashAndPassword([]byte(chat.Password), []byte(r.Password)); err3 != nil {
		err3 = status.WrongPassword
		return "", err3
	}
	token, err2 := middleware.GenerateToken(r.Account)
	if err != nil {
		return "", err2
	}
	return token, nil
}

func (c *ChatService) Find(account string) (string, error) {
	return c.Login(request.Login{
		account,
		"",
	})
}

func (c *ChatService) History(send, receive string) ([]entity.Message, error) {
	var records []entity.Message
	if err := global.Global_Db.Where("sender = ? AND receiver = ? OR sender = ? AND receiver = ?", send, receive, receive, send).Order("date DESC").Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
