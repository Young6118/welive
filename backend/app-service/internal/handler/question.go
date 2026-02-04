package handler

import (
	"net/http"
	"strconv"

	"ai-egg/app-service/internal/config"
	"ai-egg/app-service/internal/model"

	"github.com/gin-gonic/gin"
)

type CreateQuestionRequest struct {
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Tags    []string `json:"tags"`
}

type CreateAnswerRequest struct {
	QuestionID uint   `json:"questionId" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

// GetQuestions 获取问题列表
func GetQuestions(c *gin.Context) {
	db := config.GetDB()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	category := c.Query("category")

	var questions []model.Question
	query := db.Model(&model.Question{}).Where("status = ?", 1)

	// 根据分类筛选
	if category != "" && category != "recommend" {
		query = query.Where("tags LIKE ?", "%"+category+"%")
	}

	var total int64
	query.Count(&total)

	offset := (page - 1) * pageSize
	result := query.Preload("Author").Order("created_at DESC").Limit(pageSize).Offset(offset).Find(&questions)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "获取问题列表失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data: gin.H{
			"list":  questions,
			"total": total,
		},
	})
}

// GetQuestion 获取问题详情
func GetQuestion(c *gin.Context) {
	db := config.GetDB()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的问题ID",
			Data:    nil,
		})
		return
	}

	var question model.Question
	result := db.Preload("Author").First(&question, id)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "问题不存在",
			Data:    nil,
		})
		return
	}

	// 增加浏览量
	db.Model(&question).UpdateColumn("views", question.Views+1)

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data:    question,
	})
}

// CreateQuestion 创建问题
func CreateQuestion(c *gin.Context) {
	db := config.GetDB()

	var req CreateQuestionRequest
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

	question := model.Question{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: userID.(uint),
		Tags:     "",
		Likes:    0,
		Views:    0,
		Status:   1,
	}

	result := db.Create(&question)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "创建问题失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "发布成功",
		Data: gin.H{
			"id": question.ID,
		},
	})
}

// CreateAnswer 创建回答
func CreateAnswer(c *gin.Context) {
	db := config.GetDB()

	var req CreateAnswerRequest
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

	// 检查问题是否存在
	var question model.Question
	result := db.First(&question, req.QuestionID)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "问题不存在",
			Data:    nil,
		})
		return
	}

	answer := model.Answer{
		QuestionID: req.QuestionID,
		Content:    req.Content,
		AuthorID:   userID.(uint),
		Likes:      0,
		Status:     1,
	}

	result = db.Create(&answer)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "创建回答失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "回答成功",
		Data: gin.H{
			"id": answer.ID,
		},
	})
}

// LikeQuestion 点赞问题
func LikeQuestion(c *gin.Context) {
	db := config.GetDB()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的问题ID",
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

	// 检查问题是否存在
	var question model.Question
	result := db.First(&question, id)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "问题不存在",
			Data:    nil,
		})
		return
	}

	// 检查是否已经点赞
	var existingLike model.QuestionLike
	result = db.Where("question_id = ? AND user_id = ?", id, userID.(uint)).First(&existingLike)
	if result.Error == nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "已经点赞过了",
			Data:    nil,
		})
		return
	}

	// 创建点赞记录
	like := model.QuestionLike{
		QuestionID: uint(id),
		UserID:     userID.(uint),
	}

	// 使用事务
	tx := db.Begin()
	if err := tx.Create(&like).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "点赞失败",
			Data:    nil,
		})
		return
	}

	// 增加点赞数
	if err := tx.Model(&question).UpdateColumn("likes", question.Likes+1).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "点赞失败",
			Data:    nil,
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "点赞成功",
		Data:    nil,
	})
}

// UnlikeQuestion 取消点赞问题
func UnlikeQuestion(c *gin.Context) {
	db := config.GetDB()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的问题ID",
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

	// 检查问题是否存在
	var question model.Question
	result := db.First(&question, id)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "问题不存在",
			Data:    nil,
		})
		return
	}

	// 检查是否已经点赞
	var existingLike model.QuestionLike
	result = db.Where("question_id = ? AND user_id = ?", id, userID.(uint)).First(&existingLike)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "还没有点赞",
			Data:    nil,
		})
		return
	}

	// 使用事务
	tx := db.Begin()
	if err := tx.Delete(&existingLike).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "取消点赞失败",
			Data:    nil,
		})
		return
	}

	// 减少点赞数
	if err := tx.Model(&question).UpdateColumn("likes", question.Likes-1).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "取消点赞失败",
			Data:    nil,
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "取消点赞成功",
		Data:    nil,
	})
}
