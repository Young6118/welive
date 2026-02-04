package handler

import (
	"net/http"
	"strconv"

	"ai-egg/app-service/internal/config"
	"ai-egg/app-service/internal/model"

	"github.com/gin-gonic/gin"
)

type CreateCommentRequest struct {
	TargetID   uint   `json:"targetId" binding:"required"`
	TargetType string `json:"targetType" binding:"required"`
	Content    string `json:"content" binding:"required"`
	ParentID   *uint  `json:"parentId"`
}

type ReplyCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	db := config.GetDB()

	var req CreateCommentRequest
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

	comment := model.Comment{
		TargetID:   req.TargetID,
		TargetType: req.TargetType,
		Content:    req.Content,
		AuthorID:   userID.(uint),
		ParentID:   req.ParentID,
		Likes:      0,
		Status:     1,
	}

	result := db.Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "创建评论失败",
			Data:    nil,
		})
		return
	}

	// 预加载作者信息
	db.Preload("Author").First(&comment, comment.ID)

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "评论成功",
		Data:    comment,
	})
}

// GetComments 获取评论列表
func GetComments(c *gin.Context) {
	db := config.GetDB()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	targetID := c.Query("targetId")
	targetType := c.Query("targetType")

	var comments []model.Comment
	query := db.Model(&model.Comment{}).Where("status = ?", 1)

	// 根据目标筛选
	if targetID != "" {
		if id, err := strconv.ParseUint(targetID, 10, 64); err == nil {
			query = query.Where("target_id = ?", id)
		}
	}
	if targetType != "" {
		query = query.Where("target_type = ?", targetType)
	}

	var total int64
	query.Count(&total)

	offset := (page - 1) * pageSize
	result := query.Preload("Author").Order("created_at DESC").Limit(pageSize).Offset(offset).Find(&comments)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "获取评论列表失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data: gin.H{
			"list":  comments,
			"total": total,
		},
	})
}

// GetComment 获取评论详情
func GetComment(c *gin.Context) {
	db := config.GetDB()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的评论ID",
			Data:    nil,
		})
		return
	}

	var comment model.Comment
	result := db.Preload("Author").First(&comment, id)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "评论不存在",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data:    comment,
	})
}

// LikeComment 点赞评论
func LikeComment(c *gin.Context) {
	db := config.GetDB()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的评论ID",
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

	// 检查评论是否存在
	var comment model.Comment
	result := db.First(&comment, id)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "评论不存在",
			Data:    nil,
		})
		return
	}

	// 检查是否已经点赞
	var existingLike model.CommentLike
	likeResult := db.Where("comment_id = ? AND user_id = ?", id, userID.(uint)).First(&existingLike)
	if likeResult.Error == nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "已经点赞过了",
			Data:    nil,
		})
		return
	}

	// 创建点赞记录
	like := model.CommentLike{
		CommentID: uint(id),
		UserID:    userID.(uint),
	}
	if err := db.Create(&like).Error; err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "点赞失败",
			Data:    nil,
		})
		return
	}

	// 更新评论点赞数
	db.Model(&comment).UpdateColumn("likes", comment.Likes+1)

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "点赞成功",
		Data:    nil,
	})
}

// UnlikeComment 取消点赞评论
func UnlikeComment(c *gin.Context) {
	db := config.GetDB()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的评论ID",
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

	// 检查评论是否存在
	var comment model.Comment
	result := db.First(&comment, id)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "评论不存在",
			Data:    nil,
		})
		return
	}

	// 检查是否已点赞
	var existingLike model.CommentLike
	likeResult := db.Where("comment_id = ? AND user_id = ?", id, userID.(uint)).First(&existingLike)
	if likeResult.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "还没有点赞",
			Data:    nil,
		})
		return
	}

	// 删除点赞记录
	if err := db.Delete(&existingLike).Error; err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "取消点赞失败",
			Data:    nil,
		})
		return
	}

	// 更新评论点赞数
	if comment.Likes > 0 {
		db.Model(&comment).UpdateColumn("likes", comment.Likes-1)
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "取消点赞成功",
		Data:    nil,
	})
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	db := config.GetDB()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的评论ID",
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

	// 检查评论是否存在
	var comment model.Comment
	result := db.First(&comment, id)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "评论不存在",
			Data:    nil,
		})
		return
	}

	// 检查是否是评论作者
	if comment.AuthorID != userID.(uint) {
		c.JSON(http.StatusOK, Response{
			Code:    403,
			Message: "无权删除此评论",
			Data:    nil,
		})
		return
	}

	// 软删除评论
	if err := db.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "删除评论失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "删除成功",
		Data:    nil,
	})
}

// ReplyComment 回复评论
func ReplyComment(c *gin.Context) {
	db := config.GetDB()

	parentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的评论ID",
			Data:    nil,
		})
		return
	}

	var req ReplyCommentRequest
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

	// 检查父评论是否存在
	var parentComment model.Comment
	result := db.First(&parentComment, parentID)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "评论不存在",
			Data:    nil,
		})
		return
	}

	// 创建回复评论
	parentIDUint := uint(parentID)
	comment := model.Comment{
		TargetID:   parentComment.TargetID,
		TargetType: parentComment.TargetType,
		Content:    req.Content,
		AuthorID:   userID.(uint),
		ParentID:   &parentIDUint,
		Likes:      0,
		Status:     1,
	}

	result = db.Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "回复失败",
			Data:    nil,
		})
		return
	}

	// 预加载作者信息
	db.Preload("Author").First(&comment, comment.ID)

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "回复成功",
		Data:    comment,
	})
}
