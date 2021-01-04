package v1

import (
	"ginblog/model"
	"ginblog/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == msg.SUCCESS {
		model.CreateUser(&data)
	}

	if code == msg.ErrUsernameUsed {
		code = msg.ErrUsernameUsed
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"msg": msg.GetMsg(code),
	})
}

// 查询单个用户


// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize,_ := strconv.Atoi(c.Query("pageSize"))
	pageNum,_ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetUsers(pageSize, pageNum)
	code = msg.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"data":data,
		"msg":msg.GetMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {

}

// 删除用户
func DeleteUser(c *gin.Context) {

}