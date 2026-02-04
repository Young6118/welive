package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
}

func GetUser(c *gin.Context) {
	// TODO: 从数据库获取用户信息
	user := UserInfo{
		ID:       1,
		Username: "test_user",
		Email:    "test@example.com",
		Avatar:   "",
		Bio:      "",
	}
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data:    user,
	})
}

func UpdateUser(c *gin.Context) {
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// TODO: 更新用户信息
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "更新成功",
		Data:    nil,
	})
}
