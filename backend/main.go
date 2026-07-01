package main

import (
	"todo-blog/database"
	"todo-blog/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()       // 初始化数据库
	r := gin.Default()    // 创建引擎
	routes.SetupRoutes(r) // 注册路由
	r.Run(":8080")        // 启动服务
}
