package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendMessageRequest struct {
	ReceiverID uint   `json:"receiverId" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Type       string `json:"type"`
}

type Message struct {
	ID         uint   `json:"id"`
	SenderID   uint   `json:"senderId"`
	ReceiverID uint   `json:"receiverId"`
	Content    string `json:"content"`
	Type       string `json:"type"`
}

func SendMessage(c *gin.Context) {
	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// TODO: 保存消息到数据库
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "发送成功",
		Data: gin.H{
			"id": 1,
		},
	})
}

func GetChatHistory(c *gin.Context) {
	// chatID := c.Param("id")
	// TODO: 从数据库获取聊天记录
	messages := []Message{
		{ID: 1, SenderID: 1, ReceiverID: 2, Content: "你好", Type: "text"},
		{ID: 2, SenderID: 2, ReceiverID: 1, Content: "你好！", Type: "text"},
	}
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data: gin.H{
			"list":  messages,
			"total": 2,
		},
	})
}
