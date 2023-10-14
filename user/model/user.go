package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User /**
type User struct {
	gorm.Model
	Username       string `gorm:"type:varchar(20);not null;unique;comment:id" json:"username"`
	PasswordDigset string
}

const (
	PassWordCost = 12
)

// 加密密码
func (user *User) SetPassword(password string) error {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigset = string(fromPassword)
	return nil
}

// 检验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigset), []byte(password))
	return err == nil
}

func (User) TableName() string {
	return "user"
}

func (user *User) GetUserModelByUsername(username string) (tx *gorm.DB) {
	return Db.Where("username = ?", username)
}
