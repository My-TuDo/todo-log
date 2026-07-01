package handler

import (
	"todo-blog/database"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	var articles = []map[string]any{}
	// 这里可以从数据库中获取文章列表，暂时返回模拟数据
	rows, err := database.DB.Query("SELECT id, title, summary, created_at, views FROM articles")
	if err != nil {
		c.JSON(500, gin.H{"error": "查询失败"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title, summary, createdAt string
		var views int

		err := rows.Scan(&id, &title, &summary, &createdAt, &views)
		if err != nil {
			c.JSON(500, gin.H{"error": "数据解析失败"})
			return
		}

		// 返回 JSON
		articles = append(articles, map[string]any{
			"id":         id,
			"title":      title,
			"summary":    summary,
			"created_at": createdAt,
			"views":      views,
		})
	}
	c.JSON(200, articles)

}
