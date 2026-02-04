package main

import (
	"ai-egg/app-service/internal/config"
	"ai-egg/app-service/internal/model"
	"ai-egg/app-service/internal/router"
	"log"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库
	config.InitDB(cfg)

	// 自动迁移数据库表结构
	if err := config.AutoMigrate(
		&model.User{},
		&model.Question{},
		&model.Answer{},
		&model.QuestionLike{},
		&model.Note{},
		&model.Comment{},
		&model.CommentLike{},
		&model.Chat{},
		&model.Message{},
		&model.Village{},
		&model.VillageMember{},
		&model.Post{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully")

	// 初始化Redis
	config.InitRedis(cfg)

	// 设置路由
	r := router.SetupRouter()

	// 启动服务 - 监听所有网卡，支持局域网访问
	addr := "0.0.0.0:" + cfg.Server.Port
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
