package handler

import (
	"net/http"
	"strconv"

	"ai-egg/app-service/internal/config"
	"ai-egg/app-service/internal/model"

	"github.com/gin-gonic/gin"
)

type SendMessageRequest struct {
	ReceiverID uint   `json:"receiverId" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Type       string `json:"type"`
}

func SendMessage(c *gin.Context) {
	db := config.GetDB()

	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// 从上下文获取当前用户ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, Response{
			Code:    401,
			Message: "未登录",
			Data:    nil,
		})
		return
	}

	// 设置默认消息类型
	msgType := req.Type
	if msgType == "" {
		msgType = "text"
	}

	// 查找或创建聊天会话
	var chat model.Chat
	chatResult := db.Where(
		"(user_id = ? AND receiver_id = ?) OR (user_id = ? AND receiver_id = ?)",
		userID.(uint), req.ReceiverID, req.ReceiverID, userID.(uint),
	).First(&chat)

	if chatResult.Error != nil {
		// 创建新会话
		chat = model.Chat{
			UserID:     userID.(uint),
			ReceiverID: req.ReceiverID,
			Type:       "user",
		}
		if err := db.Create(&chat).Error; err != nil {
			c.JSON(http.StatusOK, Response{
				Code:    500,
				Message: "创建聊天会话失败",
				Data:    nil,
			})
			return
		}
	}

	// 创建消息
	message := model.Message{
		ChatID:   chat.ID,
		SenderID: userID.(uint),
		Content:  req.Content,
		Type:     msgType,
		Status:   1,
	}

	if err := db.Create(&message).Error; err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "保存消息失败",
			Data:    nil,
		})
		return
	}

	// 更新会话最后消息
	db.Model(&chat).UpdateColumn("last_message", req.Content)

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "发送成功",
		Data: gin.H{
			"id": message.ID,
		},
	})
}

func GetChatHistory(c *gin.Context) {
	db := config.GetDB()

	chatID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的聊天ID",
			Data:    nil,
		})
		return
	}

	// 从上下文获取当前用户ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, Response{
			Code:    401,
			Message: "未登录",
			Data:    nil,
		})
		return
	}

	// 验证用户是否有权限查看此聊天
	var chat model.Chat
	if result := db.First(&chat, chatID); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "聊天不存在",
			Data:    nil,
		})
		return
	}

	// 检查用户是否是聊天参与者
	if chat.UserID != userID.(uint) && chat.ReceiverID != userID.(uint) {
		c.JSON(http.StatusOK, Response{
			Code:    403,
			Message: "无权查看此聊天",
			Data:    nil,
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	var messages []model.Message
	var total int64

	db.Model(&model.Message{}).Where("chat_id = ?", chatID).Count(&total)

	offset := (page - 1) * pageSize
	result := db.Where("chat_id = ?", chatID).
		Order("created_at DESC").
		Limit(pageSize).Offset(offset).
		Find(&messages)

	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "获取聊天记录失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data: gin.H{
			"list":  messages,
			"total": total,
		},
	})
}
