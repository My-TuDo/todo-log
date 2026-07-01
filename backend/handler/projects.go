package handler

import (
	"todo-blog/database"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	var projects = []map[string]any{}
	rows, err := database.DB.Query("SELECT * FROM projects")
	if err != nil {
		c.JSON(500, gin.H{"error": "查询失败"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, description, tech string

		err := rows.Scan(&id, &name, &description, &tech)
		if err != nil {
			c.JSON(500, gin.H{"error": "数据解析失败"})
			return
		}

		projects = append(projects, map[string]any{
			"id":          id,
			"name":        name,
			"description": description,
			"tech":        tech,
		})
	}
	c.JSON(200, projects)
}
