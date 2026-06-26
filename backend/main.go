package main

import (
	"net/http"

	"todo-blog/handler"
)

func main() {
	// 注册路由
	http.HandleFunc("/api/articles", handler.GetArticles)
	http.HandleFunc("/api/projects", handler.GetProjects)
	http.HandleFunc("/api/profile", handler.GetProfile)
	http.HandleFunc("/api/contact", handler.GetContact)
	http.HandleFunc("/api/login", handler.Login)

	// 处理请求
	http.ListenAndServe(":8080", nil)
}
