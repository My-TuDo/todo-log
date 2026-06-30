package routes

import (
	"todo-blog/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/api/articles", handler.GetArticles)
	r.GET("/api/projects", handler.GetProjects)
	r.GET("/api/profile", handler.GetProfile)
	r.GET("/api/contact", handler.GetContact)
	r.POST("/api/login", handler.Login)

}
