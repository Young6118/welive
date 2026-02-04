package main

import (
	"ai-egg-backend/internal/config"
	"ai-egg-backend/internal/router"
	"ai-egg-backend/pkg/database"
	"ai-egg-backend/pkg/logger"
	"ai-egg-backend/pkg/redis"
	"log"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化日志
	logger.Init(cfg.Log)

	// 初始化数据库
	db, err := database.Init(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to init database: %v", err)
	}

	// 初始化Redis
	rdb := redis.Init(cfg.Redis)

	// 设置路由
	r := router.SetupRouter(db, rdb, cfg)

	// 启动服务
	log.Printf("Server starting on %s", cfg.Server.Addr)
	if err := r.Run(cfg.Server.Addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
