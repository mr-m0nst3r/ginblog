package model

import (
	"ginblog/utils/msg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CheckCategory(name string) (code int)  {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)

	if category.ID > 0 {
		return msg.ErrCategoryUsed
	}

	return msg.SUCCESS
}

func CreateCategory(data *Category) int {
	//data.Password = EncryptPw(data.Password)
	err = db.Create(&data).Error
	if err != nil {
		return msg.ERROR
	}

	return msg.SUCCESS
}

func GetCategories(pageSize int, pageNum int) []Category {
	var categories []Category
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&categories).Error

	if err != nil && err!=gorm.ErrRecordNotFound{
		return nil
	}
	return categories
}

func DeleteCategory(id int) int {
	err = db.Where("ID = ?", id).Delete(&Category{}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func EditCategory(id int, data *Category) int {
	var category Category
	maps := make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Debug().Model(&category).Where("id = ?",id).Update(maps).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func GetCategoryArticles()  {

}