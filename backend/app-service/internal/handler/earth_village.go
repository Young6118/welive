package handler

import (
	"net/http"
	"strconv"

	"ai-egg/app-service/internal/config"
	"ai-egg/app-service/internal/model"

	"github.com/gin-gonic/gin"
)

type CreatePostRequest struct {
	Content string `json:"content" binding:"required"`
	Images  string `json:"images"`
}

type ReplyPostRequest struct {
	Content string `json:"content" binding:"required"`
}

func JoinVillage(c *gin.Context) {
	db := config.GetDB()

	villageID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的村落ID",
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

	// 检查村落是否存在
	var village model.Village
	if result := db.First(&village, villageID); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "村落不存在",
			Data:    nil,
		})
		return
	}

	// 检查是否已加入
	var existingMember model.VillageMember
	if result := db.Where("village_id = ? AND user_id = ?", villageID, userID.(uint)).First(&existingMember); result.Error == nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "已加入该村落",
			Data:    nil,
		})
		return
	}

	// 创建成员记录
	member := model.VillageMember{
		VillageID: uint(villageID),
		UserID:    userID.(uint),
		Role:      0,
	}

	if err := db.Create(&member).Error; err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "加入村落失败",
			Data:    nil,
		})
		return
	}

	// 更新村落成员数
	db.Model(&village).UpdateColumn("member_count", village.MemberCount+1)

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "加入成功",
		Data:    nil,
	})
}

func LeaveVillage(c *gin.Context) {
	db := config.GetDB()

	villageID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的村落ID",
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

	// 检查村落是否存在
	var village model.Village
	if result := db.First(&village, villageID); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "村落不存在",
			Data:    nil,
		})
		return
	}

	// 查找成员记录
	var member model.VillageMember
	if result := db.Where("village_id = ? AND user_id = ?", villageID, userID.(uint)).First(&member); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "未加入该村落",
			Data:    nil,
		})
		return
	}

	// 删除成员记录
	if err := db.Delete(&member).Error; err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "退出村落失败",
			Data:    nil,
		})
		return
	}

	// 更新村落成员数
	if village.MemberCount > 0 {
		db.Model(&village).UpdateColumn("member_count", village.MemberCount-1)
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "退出成功",
		Data:    nil,
	})
}

func GetVillages(c *gin.Context) {
	db := config.GetDB()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	var villages []model.Village
	var total int64

	db.Model(&model.Village{}).Where("status = ?", 1).Count(&total)

	offset := (page - 1) * pageSize
	result := db.Where("status = ?", 1).Order("created_at DESC").Limit(pageSize).Offset(offset).Find(&villages)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "获取村落列表失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data: gin.H{
			"list":  villages,
			"total": total,
		},
	})
}

func GetVillage(c *gin.Context) {
	db := config.GetDB()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的村落ID",
			Data:    nil,
		})
		return
	}

	var village model.Village
	result := db.First(&village, id)
	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "村落不存在",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data:    village,
	})
}

func CreatePost(c *gin.Context) {
	db := config.GetDB()

	villageID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的村落ID",
			Data:    nil,
		})
		return
	}

	var req CreatePostRequest
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

	// 检查村落是否存在
	var village model.Village
	if result := db.First(&village, villageID); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "村落不存在",
			Data:    nil,
		})
		return
	}

	post := model.Post{
		VillageID: uint(villageID),
		AuthorID:  userID.(uint),
		Content:   req.Content,
		Images:    req.Images,
		Likes:     0,
		Comments:  0,
		Status:    1,
	}

	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "发布帖子失败",
			Data:    nil,
		})
		return
	}

	// 更新村落帖子数
	db.Model(&village).UpdateColumn("post_count", village.PostCount+1)

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "发布成功",
		Data: gin.H{
			"id": post.ID,
		},
	})
}

func GetPosts(c *gin.Context) {
	db := config.GetDB()

	villageID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的村落ID",
			Data:    nil,
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	var posts []model.Post
	var total int64

	db.Model(&model.Post{}).Where("village_id = ? AND status = ?", villageID, 1).Count(&total)

	offset := (page - 1) * pageSize
	result := db.Where("village_id = ? AND status = ?", villageID, 1).
		Preload("Author").
		Order("created_at DESC").
		Limit(pageSize).Offset(offset).
		Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "获取帖子列表失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data: gin.H{
			"list":  posts,
			"total": total,
		},
	})
}

func LikePost(c *gin.Context) {
	db := config.GetDB()

	postID, err := strconv.ParseUint(c.Param("postId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的帖子ID",
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

	// 检查帖子是否存在
	var post model.Post
	if result := db.First(&post, postID); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "帖子不存在",
			Data:    nil,
		})
		return
	}

	// 检查是否已点赞
	var existingLike model.PostLike
	if result := db.Where("post_id = ? AND user_id = ?", postID, userID.(uint)).First(&existingLike); result.Error == nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "已经点赞过了",
			Data:    nil,
		})
		return
	}

	// 创建点赞记录
	like := model.PostLike{
		PostID: uint(postID),
		UserID: userID.(uint),
	}
	if err := db.Create(&like).Error; err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "点赞失败",
			Data:    nil,
		})
		return
	}

	// 更新帖子点赞数
	db.Model(&post).UpdateColumn("likes", post.Likes+1)

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "点赞成功",
		Data:    nil,
	})
}

func UnlikePost(c *gin.Context) {
	db := config.GetDB()

	postID, err := strconv.ParseUint(c.Param("postId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的帖子ID",
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

	// 检查帖子是否存在
	var post model.Post
	if result := db.First(&post, postID); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "帖子不存在",
			Data:    nil,
		})
		return
	}

	// 检查是否已点赞
	var existingLike model.PostLike
	if result := db.Where("post_id = ? AND user_id = ?", postID, userID.(uint)).First(&existingLike); result.Error != nil {
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

	// 更新帖子点赞数
	if post.Likes > 0 {
		db.Model(&post).UpdateColumn("likes", post.Likes-1)
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "取消点赞成功",
		Data:    nil,
	})
}

func DeletePost(c *gin.Context) {
	db := config.GetDB()

	postID, err := strconv.ParseUint(c.Param("postId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的帖子ID",
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

	// 检查帖子是否存在
	var post model.Post
	if result := db.First(&post, postID); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "帖子不存在",
			Data:    nil,
		})
		return
	}

	// 检查是否是帖子作者
	if post.AuthorID != userID.(uint) {
		c.JSON(http.StatusOK, Response{
			Code:    403,
			Message: "无权删除此帖子",
			Data:    nil,
		})
		return
	}

	// 软删除帖子
	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "删除帖子失败",
			Data:    nil,
		})
		return
	}

	// 更新村落帖子数
	var village model.Village
	if result := db.First(&village, post.VillageID); result.Error == nil && village.PostCount > 0 {
		db.Model(&village).UpdateColumn("post_count", village.PostCount-1)
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "删除成功",
		Data:    nil,
	})
}

func ReplyPost(c *gin.Context) {
	db := config.GetDB()

	postID, err := strconv.ParseUint(c.Param("postId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的帖子ID",
			Data:    nil,
		})
		return
	}

	var req ReplyPostRequest
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

	// 检查帖子是否存在
	var post model.Post
	if result := db.First(&post, postID); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    404,
			Message: "帖子不存在",
			Data:    nil,
		})
		return
	}

	// 创建评论
	comment := model.Comment{
		TargetID:   uint(postID),
		TargetType: "post",
		Content:    req.Content,
		AuthorID:   userID.(uint),
		Likes:      0,
		Status:     1,
	}

	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "回复失败",
			Data:    nil,
		})
		return
	}

	// 更新帖子评论数
	db.Model(&post).UpdateColumn("comments", post.Comments+1)

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "回复成功",
		Data: gin.H{
			"id": comment.ID,
		},
	})
}

func GetReplies(c *gin.Context) {
	db := config.GetDB()

	postID, err := strconv.ParseUint(c.Param("postId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "无效的帖子ID",
			Data:    nil,
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	var comments []model.Comment
	var total int64

	db.Model(&model.Comment{}).
		Where("target_id = ? AND target_type = ? AND status = ?", postID, "post", 1).
		Count(&total)

	offset := (page - 1) * pageSize
	result := db.Where("target_id = ? AND target_type = ? AND status = ?", postID, "post", 1).
		Preload("Author").
		Order("created_at DESC").
		Limit(pageSize).Offset(offset).
		Find(&comments)

	if result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "获取回复列表失败",
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
