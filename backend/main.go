package main

import (
	"net/http"
)

func main() {
	// 创建路由并注册路由
	http.HandleFunc("/api/articles", func(w http.ResponseWriter, r *http.Request) {
		// 设置响应头为 JSON
		w.Header().Set("Content-Type", "application/json")
		// 挂 CORS 中间件
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// 设置响应体
		w.Write([]byte(`[
			{"id": 1, "title": "你好，能看见吗", "summary": "这是文章的摘要描述文字，后续将从后端 API 获取真实数据...", "created_at": "2026-06-01", "views": 10},
			{"id": 2, "title": "示例文章标题 2", "summary": "这是文章的摘要描述文字，后续将从后端 API 获取真实数据...", "created_at": "2026-06-02", "views": 20},
			{"id": 3, "title": "示例文章标题 3", "summary": "这是文章的摘要描述文字，后续将从后端 API 获取真实数据...", "created_at": "2026-06-03", "views": 30},
			{"id": 4, "title": "示例文章标题 4", "summary": "这是文章的摘要描述文字，后续将从后端 API 获取真实数据...", "created_at": "2026-06-04", "views": 40},
			{"id": 5, "title": "示例文章标题 5", "summary": "这是文章的摘要描述文字，后续将从后端 API 获取真实数据...", "created_at": "2026-06-05", "views": 50}
		]`))
	})
	// 处理请求
	http.ListenAndServe(":8080", nil)
}
