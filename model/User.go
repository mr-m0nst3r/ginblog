package model

import (
	"ginblog/utils/msg"
	"github.com/jinzhu/gorm"
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