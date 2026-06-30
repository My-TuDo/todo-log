package handler

import "github.com/gin-gonic/gin"

func GetProjects(c *gin.Context) {
	c.JSON(200, []map[string]any{
		{"id": 1, "name": "TODO Blog", "description": "仿 macOS 桌面的个人博客系统", "tech": "Vue 3 + Go"},
		{"id": 2, "name": "Task Manager", "description": "简洁高效的待办事项管理工具", "tech": "React + Node.js"},
		{"id": 3, "name": "Weather App", "description": "实时天气查询与预报应用", "tech": "Vue 3 + API"},
	})
}
