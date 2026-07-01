package handler

import (
	"todo-blog/database"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	var profile = []map[string]any{}
	rows, err := database.DB.Query("SELECT * FROM profile LIMIT 1")
	if err != nil {
		c.JSON(500, gin.H{"error": "查询失败"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, title, avatar_emoji, bio, tags string

		err := rows.Scan(&id, &name, &title, &avatar_emoji, &bio, &tags)
		if err != nil {
			c.JSON(500, gin.H{"error": "数据解析失败"})
			return
		}

		profile = append(profile, map[string]any{
			"id":           id,
			"name":         name,
			"title":        title,
			"avatar_emoji": avatar_emoji,
			"bio":          bio,
			"tags":         tags,
		})
	}
	c.JSON(200, profile)
}
