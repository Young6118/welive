package handler

import (
	"net/http"
	"strconv"

	"ai-egg/app-service/internal/config"
	"ai-egg/app-service/internal/model"

	"github.com/gin-gonic/gin"
)

// SearchQuestions 搜索问题
func SearchQuestions(c *gin.Context) {
	db := config.GetDB()

	keyword := c.Query("keyword")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if keyword == "" {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "搜索关键词不能为空",
			Data:    nil,
		})
		return
	}

	var questions []model.Question
	query := db.Model(&model.Question{}).
		Where("status = ? AND (title LIKE ? OR content LIKE ?)",
			1, "%"+keyword+"%", "%"+keyword+"%")

	var total int64
	query.Count(&total)

	offset := (page - 1) * pageSize
	result := query.Preload("Author").
		Order("created_at DESC").
		Limit(pageSize).Offset(offset).
		Find(&questions)

	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "搜索问题失败",
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

// SearchNotes 搜索笔记
func SearchNotes(c *gin.Context) {
	db := config.GetDB()

	keyword := c.Query("keyword")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if keyword == "" {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "搜索关键词不能为空",
			Data:    nil,
		})
		return
	}

	var notes []model.Note
	query := db.Model(&model.Note{}).
		Where("status = ? AND (title LIKE ? OR content LIKE ?)",
			1, "%"+keyword+"%", "%"+keyword+"%")

	var total int64
	query.Count(&total)

	offset := (page - 1) * pageSize
	result := query.Preload("Author").
		Order("created_at DESC").
		Limit(pageSize).Offset(offset).
		Find(&notes)

	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "搜索笔记失败",
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
