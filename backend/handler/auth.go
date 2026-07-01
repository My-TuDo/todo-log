package handler

import (
	"time"
	"todo-blog/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

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
	if err := bcrypt.CompareHashAndPassword([]byte(config.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(401, gin.H{"error": "密码错误"})
		return
	}

	IssueToken(c) // 签发 JWT Token
}

// 签发 JWT Token
func IssueToken(c *gin.Context) {
	// 创建 claims
	claims := jwt.MapClaims{
		"role": "admin",                               // 用户角色
		"exp":  time.Now().Add(24 * time.Hour).Unix(), // 24h timeout
	}

	// 用 HS256 算法签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 用密钥加密成字符串
	tokenString, _ := token.SignedString([]byte(config.JWTSecret))

	c.JSON(200, gin.H{"token": tokenString})
}
