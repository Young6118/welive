package router

import (
	"ai-egg/app-service/internal/handler"
	"ai-egg/app-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 公开路由
	public := r.Group("/api/v1")
	{
		public.POST("/register", handler.Register)
		public.POST("/login", handler.Login)
		public.GET("/check-login", handler.CheckLogin)
	}

	// 需要认证的路由
	authorized := r.Group("/api/v1")
	authorized.Use(middleware.Auth())
	{
		authorized.POST("/logout", handler.Logout)

		// 用户模块
		authorized.GET("/user", handler.GetUser)
		authorized.PUT("/user", handler.UpdateUser)

		// 问答模块
		authorized.POST("/question", handler.CreateQuestion)
		authorized.POST("/answer", handler.CreateAnswer)
		authorized.GET("/questions", handler.GetQuestions)
		authorized.GET("/question/:id", handler.GetQuestion)
		authorized.POST("/question/:id/like", handler.LikeQuestion)
		authorized.POST("/question/:id/unlike", handler.UnlikeQuestion)

		// 评论模块
		authorized.POST("/comment", handler.CreateComment)
		authorized.GET("/comments", handler.GetComments)
		authorized.GET("/comment/:id", handler.GetComment)
		authorized.POST("/comment/:id/like", handler.LikeComment)
		authorized.POST("/comment/:id/unlike", handler.UnlikeComment)
		authorized.DELETE("/comment/:id", handler.DeleteComment)
		authorized.POST("/comment/:id/reply", handler.ReplyComment)

		// 笔记模块
		authorized.POST("/note", handler.CreateNote)
		authorized.GET("/notes", handler.GetNotes)
		authorized.GET("/note/:id", handler.GetNote)
		authorized.GET("/note/categories", handler.GetNoteCategories)
		authorized.GET("/note/category/:id", handler.GetNotesByCategory)

		// 聊天模块
		authorized.POST("/chat", handler.SendMessage)
		authorized.GET("/chat/:id", handler.GetChatHistory)

		// 地球村模块
		authorized.POST("/earth-village/join", handler.JoinVillage)
		authorized.POST("/earth-village/leave", handler.LeaveVillage)
		authorized.GET("/earth-villages", handler.GetVillages)
		authorized.GET("/earth-village/:id", handler.GetVillage)
		authorized.POST("/earth-village/:id/post", handler.CreatePost)
		authorized.GET("/earth-village/:id/posts", handler.GetPosts)
		authorized.POST("/earth-village/:id/post/:postId/like", handler.LikePost)
		authorized.POST("/earth-village/:id/post/:postId/unlike", handler.UnlikePost)
		authorized.DELETE("/earth-village/:id/post/:postId", handler.DeletePost)
		authorized.POST("/earth-village/:id/post/:postId/reply", handler.ReplyPost)
		authorized.GET("/earth-village/:id/post/:postId/replies", handler.GetReplies)

		// 搜索模块
		authorized.GET("/search/questions", handler.SearchQuestions)
		authorized.GET("/search/notes", handler.SearchNotes)
	}

	return r
}
