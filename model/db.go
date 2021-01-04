package model

import (
	"fmt"
	"ginblog/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDb()  {
	db,err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,utils.DbPass,utils.DbHost,utils.DbPort,utils.DbName))

	if err != nil {
		fmt.Println("连接数据库失败：",err)
	}

	db.SingularTable(true)
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(10*time.Second)
	db.DB().SetMaxOpenConns(100)

	//db.Close()
	//defer db.Close()
}
