package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const PasswordHash = "$2a$10$wIQRM2F4Th8a.HLtU667GeDDnAn2pr60WS.wK48dE6Vz57DwOgkdu" // 预先生成的密码哈希

// 功能实现：用户登录验证
func Login(c *gin.Context) {
	// 获取请求中的密码
	var req struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "请求参数错误"})
	}

	// 进行密码比对
	if err := bcrypt.CompareHashAndPassword([]byte("$2a$10$wIQRM2F4Th8a.HLtU667GeDDnAn2pr60WS.wK48dE6Vz57DwOgkdu"), []byte(req.Password)); err != nil {
		c.JSON(401, gin.H{"error": "密码错误"})
		return
	}

	c.JSON(200, gin.H{"message": "登录成功"})
}
