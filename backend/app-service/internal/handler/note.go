package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateNoteRequest struct {
	Title    string   `json:"title" binding:"required"`
	Content  string   `json:"content" binding:"required"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

type Note struct {
	ID       uint     `json:"id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
	AuthorID uint     `json:"authorId"`
}

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateNote(c *gin.Context) {
	var req CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// TODO: 保存笔记到数据库
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "发布成功",
		Data: gin.H{
			"id": 1,
		},
	})
}

func GetNotes(c *gin.Context) {
	// TODO: 从数据库获取笔记列表
	notes := []Note{
		{ID: 1, Title: "笔记1", Content: "内容1", Category: "学习", Tags: []string{"AI"}, AuthorID: 1},
		{ID: 2, Title: "笔记2", Content: "内容2", Category: "工作", Tags: []string{"项目"}, AuthorID: 1},
	}
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data: gin.H{
			"list":  notes,
			"total": 2,
		},
	})
}

func GetNote(c *gin.Context) {
	// id := c.Param("id")
	// TODO: 从数据库获取笔记详情
	note := Note{
		ID:       1,
		Title:    "笔记标题",
		Content:  "笔记内容",
		Category: "学习",
		Tags:     []string{"AI"},
		AuthorID: 1,
	}
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data:    note,
	})
}

func GetNoteCategories(c *gin.Context) {
	// TODO: 从数据库获取分类列表
	categories := []Category{
		{ID: 1, Name: "学习"},
		{ID: 2, Name: "工作"},
		{ID: 3, Name: "生活"},
	}
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data:    categories,
	})
}

func GetNotesByCategory(c *gin.Context) {
	// categoryID := c.Param("id")
	// TODO: 从数据库获取分类下的笔记
	notes := []Note{
		{ID: 1, Title: "分类笔记1", Content: "内容1", Category: "学习", Tags: []string{"AI"}, AuthorID: 1},
	}
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data: gin.H{
			"list":  notes,
			"total": 1,
		},
	})
}
