package handler

import (
	"net/http"

	"ai-egg/app-service/internal/config"
	"ai-egg/app-service/internal/model"

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
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Bio    string `json:"bio"`
}

func GetUser(c *gin.Context) {
	// 从上下文中获取当前用户ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, Response{
			Code:    401,
			Message: "未登录",
			Data:    nil,
		})
		return
	}

	db := config.GetDB()
	var user model.User
	if result := db.First(&user, userID); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "用户不存在",
			Data:    nil,
		})
		return
	}

	userInfo := UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Bio:      user.Bio,
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data:    userInfo,
	})
}

func UpdateUser(c *gin.Context) {
	// 从上下文中获取当前用户ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, Response{
			Code:    401,
			Message: "未登录",
			Data:    nil,
		})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	db := config.GetDB()
	var user model.User
	if result := db.First(&user, userID); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "用户不存在",
			Data:    nil,
		})
		return
	}

	// 更新用户信息
	updates := make(map[string]interface{})
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}
	if req.Bio != "" {
		updates["bio"] = req.Bio
	}

	if result := db.Model(&user).Updates(updates); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "更新用户信息失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "更新成功",
		Data:    nil,
	})
}
