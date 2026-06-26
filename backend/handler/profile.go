package handler

import "net/http"

func GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write([]byte(`{
		"name": "TUDO",
		"title": "全栈开发者 / 学生",
		"avatar_emoji": "🐱",
		"bio": "热爱编程、设计和一切有趣的事物。这个博客将记录我的学习笔记、项目经验和生活思考。🚀",
		"tags": ["Vue", "Go", "TypeScript", "Tailwind CSS", "Docker"]
	}`))
}
