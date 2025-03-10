package chat

import (
	"BriefChat/global"
	"BriefChat/middleware"
	"BriefChat/model/entity"
	"BriefChat/model/request"
	"BriefChat/utils/status"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
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
func (c *ChatService) Login(r request.Login) (string, *entity.Info, error) {
	var chat = entity.Chat{}
	err := global.Global_Db.Where("account = ?", r.Account).First(&chat).Error
	//println(chat.Account)
	if err != nil {
		return "", nil, status.NoAccountExist
	}
	if err3 := bcrypt.CompareHashAndPassword([]byte(chat.Password), []byte(r.Password)); err3 != nil {
		err3 = status.WrongPassword
		return "", nil, err3
	}
	token, err2 := middleware.GenerateToken(r.Account)
	if err != nil {
		return "", nil, err2
	}
	return token, &entity.Info{
		Name:   chat.SelfInfo.Name,
		Avatar: chat.SelfInfo.Avatar,
	}, nil
}

//func (c *ChatService) Find(account string) (string, error) {
//	return c.Login(request.Login{
//		account,
//		"",
//	})
//}

func (c *ChatService) History(send, receive string) ([]entity.Message, error) {
	var records []entity.Message
	if err := global.Global_Db.Where("sender = ? AND receiver = ? OR sender = ? AND receiver = ?", send, receive, receive, send).Order("date ASC").Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (c *ChatService) AddFriend(send, receive string) error {
	var records []entity.Friend
	global.Global_Db.Where("account1 = ? AND account2 = ? OR account1 = ? AND account2 = ?", send, receive, receive, send).Find(&records)
	if records != nil {
		return errors.New("friend exist")
	}
	if err := global.Global_Db.Create(&entity.Friend{
		Account1: send,
		Account2: receive,
	}); err != nil {
		return err.Error
	}
	return nil
}

func (c *ChatService) AllFriend(user string) ([]entity.FriendInfo, error) {
	var friendAccounts []string
	// 从 Friend 表中分别获取 account1 和 account2
	err := global.Global_Db.Model(&entity.Friend{}).
		Where("account1 = ? OR account2 = ?", user, user).
		Pluck("account1", &friendAccounts).
		Error
	if err != nil {
		return nil, err
	}

	// 获取 account2 的值，并追加到 friendAccounts 中
	var friendAccounts2 []string
	err = global.Global_Db.Model(&entity.Friend{}).
		Where("account1 = ? OR account2 = ?", user, user).
		Pluck("account2", &friendAccounts2).
		Error
	if err != nil {
		return nil, err
	}

	// 将 account1 和 account2 的数据合并，避免重复账号
	friendAccounts = append(friendAccounts, friendAccounts2...)
	var filteredAccounts []string
	for _, account := range friendAccounts {
		if account != user {
			filteredAccounts = append(filteredAccounts, account)
		}
	}
	//println(filteredAccounts)
	var users []entity.FriendInfo
	if err := global.Global_Db.Model(&entity.Chat{}).Select("name,avatar,account").Where("account IN ?", filteredAccounts).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (c *ChatService) GetOneUser(account string) (*entity.Chat, error) {
	var user *entity.Chat
	if err := global.Global_Db.Where("account = ?", account).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (c *ChatService) SendMessage(sender, receiver, txt, jpg string) error {
	if err := global.Global_Db.Create(&entity.Message{
		Sender:   sender,
		Receiver: receiver,
		Txt:      txt,
		Jpg:      jpg,
		Date:     time.Now().Format("2006-01-02 15:04:05"),
	}).Error; err != nil {
		return err
	}
	return nil
}

func (c *ChatService) Upload(account string, info entity.Info) error {
	updateFields := make(map[string]interface{})
	if info.Name != "" {
		updateFields["name"] = info.Name // 如果有 name 字段，才会更新
	}
	if info.Avatar != "" {
		updateFields["avatar"] = "http://" + global.Global_APP_SETTING.Domain + "/" + *global.Global_FILE + "/" + info.Avatar // 如果有 avatar 字段，才会更新
	}
	fmt.Println(global.Global_FILE)
	if err := global.Global_Db.Model(&entity.Chat{}).Where("account = ?", account).Updates(updateFields).Error; err != nil {
		return err
	}
	return nil
}
