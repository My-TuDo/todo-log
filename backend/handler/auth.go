package handler

import (
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// 处理 CORS 预检请求
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// 设置响应头为 JSON
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 只接受 POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求体内的 password
	var body struct {
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// 检查密码是否正确
	if body.Password == "admin123" {

		// 返回 token
		w.Write([]byte(`{"token":"admin-token"}`))
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
