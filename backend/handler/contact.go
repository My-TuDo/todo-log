package handler

import "github.com/gin-gonic/gin"

func GetContact(c *gin.Context) {
	c.JSON(200, []map[string]any{
		{"icon": "📧", "label": "Email", "value": "todo@example.com"},
		{"icon": "🐙", "label": "GitHub", "value": "github.com/My-TuDo"},
		{"icon": "📱", "label": "微信", "value": "TODO_WeChat"},
		{"icon": "🐦", "label": "Twitter / X", "value": "@TODO_dev"},
	})
}
