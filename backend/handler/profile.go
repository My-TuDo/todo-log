package handler

import "github.com/gin-gonic/gin"

func GetProfile(c *gin.Context) {
	c.JSON(200, map[string]any{
		"name":         "TUDO",
		"title":        "全栈开发者 / 学生",
		"avatar_emoji": "🐱",
		"bio":          "热爱编程、设计和一切有趣的事物。这个博客将记录我的学习笔记、项目经验和生活思考。🚀",
		"tags":         []string{"Vue", "Go", "TypeScript", "Tailwind CSS", "Docker"},
	})
}
