package routes

import (
	"todo-blog/handler"

	"github.com/gin-gonic/gin"
)

// CORS 中间件：允许前端跨域请求
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 若是 OPTIONS 请求，直接返回 200
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func SetupRoutes(r *gin.Engine) {
	r.Use(CORSMiddleware()) // 使用 CORS 中间件
	r.GET("/api/articles", handler.GetArticles)
	r.GET("/api/projects", handler.GetProjects)
	r.GET("/api/profile", handler.GetProfile)
	r.GET("/api/contact", handler.GetContact)
	r.POST("/api/login", handler.Login)

}
