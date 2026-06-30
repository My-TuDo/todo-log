package handler

import "github.com/gin-gonic/gin"

func GetArticles(c *gin.Context) {
	c.JSON(200, []map[string]any{
		{"id": 1, "title": "文章标题1", "content": "文章内容1"},
		{"id": 2, "title": "文章标题2", "content": "文章内容2"},
	})

}
