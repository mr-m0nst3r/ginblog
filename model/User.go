package model

import (
	"encoding/base64"
	"ginblog/utils/msg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string  `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role int `gorm:"type:int" json:"role"`
}

func CheckUser(name string) (code int)  {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)

	if users.ID > 0 {
		return msg.ErrUsernameUsed
	}

	return msg.SUCCESS
}

func CreateUser(data *User) int {
	data.Password = EncryptPw(data.Password)
	err = db.Create(&data).Error
	if err != nil {
		return msg.ERROR
	}

	return msg.SUCCESS
}

func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users).Error

	if err != nil && err!=gorm.ErrRecordNotFound{
		return nil
	}
	return users
}

func EncryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12,32,4,6,66,22,222,11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384,8,1,KeyLen)
	if err != nil {
		log.Fatal(err)
	}

	FinalPw := base64.StdEncoding.EncodeToString(HashPw)
	return FinalPw
}