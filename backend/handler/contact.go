package handler

import "net/http"

func GetContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write([]byte(`[
		{"icon": "📧", "label": "Email", "value": "todo@example.com"},
		{"icon": "🐙", "label": "GitHub", "value": "github.com/My-TuDo"},
		{"icon": "📱", "label": "微信", "value": "TODO_WeChat"},
		{"icon": "🐦", "label": "Twitter / X", "value": "@TODO_dev"}
	]`))
}
