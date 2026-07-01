package handler

import (
	"todo-blog/database"

	"github.com/gin-gonic/gin"
)

func GetContact(c *gin.Context) {
	var contactInfo = []map[string]any{}
	rows, err := database.DB.Query("SELECT * FROM contacts")
	if err != nil {
		c.JSON(500, gin.H{"error": "查询失败"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var icon, label, value string

		err := rows.Scan(&id, &icon, &label, &value)
		if err != nil {
			c.JSON(500, gin.H{"error": "数据解析失败"})
			return
		}

		contactInfo = append(contactInfo, map[string]any{
			"id":    id,
			"icon":  icon,
			"label": label,
			"value": value,
		})
	}
	c.JSON(200, contactInfo)
}
