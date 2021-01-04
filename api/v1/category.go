package v1

import (
	"ginblog/model"
	"ginblog/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == msg.SUCCESS {
		model.CreateCategory(&data)
	}

	if code == msg.ErrCategoryUsed {
		code = msg.ErrCategoryUsed
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"msg": msg.GetMsg(code),
	})
}

// 查询单个用户


// 查询用户列表
func GetCategories(c *gin.Context) {
	pageSize,_ := strconv.Atoi(c.Query("pageSize"))
	pageNum,_ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetCategories(pageSize, pageNum)
	code = msg.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"data":data,
		"msg":msg.GetMsg(code),
	})
}

// 编辑用户
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _:= strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == msg.SUCCESS {
		model.EditCategory(id,&data)
	}
	if code == msg.ErrCategoryUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"msg":msg.GetMsg(code),
	})
}

// 删除用户
func DeleteCategory(c *gin.Context) {
	id, _:= strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg": msg.GetMsg(code),
	})
}