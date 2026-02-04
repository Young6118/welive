package handler

import (
	"net/http"
	"strconv"
	"strings"

	"ai-egg/app-service/internal/config"
	"ai-egg/app-service/internal/model"

	"github.com/gin-gonic/gin"
)

type CreateNoteRequest struct {
	Title    string   `json:"title" binding:"required"`
	Content  string   `json:"content" binding:"required"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

func CreateNote(c *gin.Context) {
	db := config.GetDB()

	var req CreateNoteRequest
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

	// 将tags数组转换为逗号分隔的字符串
	tagsStr := strings.Join(req.Tags, ",")

	note := model.Note{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Tags:     tagsStr,
		AuthorID: userID.(uint),
		Status:   1,
	}

	result := db.Create(&note)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "创建笔记失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "发布成功",
		Data: gin.H{
			"id": note.ID,
		},
	})
}

func GetNotes(c *gin.Context) {
	db := config.GetDB()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	category := c.Query("category")

	var notes []model.Note
	query := db.Model(&model.Note{}).Where("status = ?", 1)

	// 根据分类筛选
	if category != "" {
		query = query.Where("category = ?", category)
	}

	var total int64
	query.Count(&total)

	offset := (page - 1) * pageSize
	result := query.Preload("Author").Order("created_at DESC").Limit(pageSize).Offset(offset).Find(&notes)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "获取笔记列表失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data: gin.H{
			"list":  notes,
			"total": total,
		},
	})
}

func GetNote(c *gin.Context) {
	db := config.GetDB()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的笔记ID",
			Data:    nil,
		})
		return
	}

	var note model.Note
	result := db.Preload("Author").First(&note, id)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "笔记不存在",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data:    note,
	})
}

func GetNoteCategories(c *gin.Context) {
	db := config.GetDB()

	// 获取所有不重复的分类
	var categories []string
	result := db.Model(&model.Note{}).Where("status = ?", 1).Distinct().Pluck("category", &categories)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "获取分类列表失败",
			Data:    nil,
		})
		return
	}

	// 转换为响应格式
	type CategoryResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	var response []CategoryResponse
	for i, cat := range categories {
		if cat != "" {
			response = append(response, CategoryResponse{
				ID:   uint(i + 1),
				Name: cat,
			})
		}
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data:    response,
	})
}

func GetNotesByCategory(c *gin.Context) {
	db := config.GetDB()

	category := c.Param("id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	var notes []model.Note
	query := db.Model(&model.Note{}).Where("status = ? AND category = ?", 1, category)

	var total int64
	query.Count(&total)

	offset := (page - 1) * pageSize
	result := query.Preload("Author").Order("created_at DESC").Limit(pageSize).Offset(offset).Find(&notes)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "获取笔记列表失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data: gin.H{
			"list":  notes,
			"total": total,
		},
	})
}
