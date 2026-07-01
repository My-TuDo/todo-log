package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Init 初始化 MySQL 数据库，创建表并插入种子数据
func Init() {

	var err error
	DB, err = sql.Open("mysql", "xzh:1234567@tcp(127.0.0.1:3306)/todo_blog?charset=utf8mb4&parseTime=True")
	if err != nil {
		log.Fatalf("打开数据库失败: %v", err)
	}

	createTables()
	seedData()
}

// createTables 创建数据库表
func createTables() {
	queries := []string{
		// 创建文章表
		`CREATE TABLE IF NOT EXISTS articles (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title TEXT NOT NULL,
			summary VARCHAR(500) NOT NULL DEFAULT '',
			created_at DATE NOT NULL DEFAULT (CURDATE()),
			views INT NOT NULL DEFAULT 0
		)`,
		// 创建项目表
		`CREATE TABLE IF NOT EXISTS projects (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name TEXT NOT NULL,
			description VARCHAR(500) NOT NULL DEFAULT '',
			tech VARCHAR(100) NOT NULL DEFAULT ''
		)`,
		// 创建个人资料表
		`CREATE TABLE IF NOT EXISTS profile (
			id INT PRIMARY KEY,
			name VARCHAR(100) NOT NULL DEFAULT 'TUDO',
			title VARCHAR(100) NOT NULL DEFAULT '',
			avatar_emoji VARCHAR(10) NOT NULL DEFAULT '',
			bio TEXT NOT NULL,
			tags TEXT NOT NULL
		)`,
		// 创建联系方式表
		`CREATE TABLE IF NOT EXISTS contacts (
			id INT AUTO_INCREMENT PRIMARY KEY,
			icon TEXT NOT NULL,
			label TEXT NOT NULL,
			value TEXT NOT NULL
		)`,
	}
	for _, q := range queries {
		if _, err := DB.Exec(q); err != nil {
			log.Fatalf("创建表失败: %v", err)
		}
	}
}

// seedData 插入种子数据（仅在表为空时）
func seedData() {
	// 插入种子文章
	var count int
	DB.QueryRow("SELECT COUNT(*) FROM articles").Scan(&count)
	if count == 0 {
		DB.Exec(`
			INSERT INTO articles (title, summary, created_at, views) VALUES
			('你好，能看见吗', '这是文章摘要，后续将从数据库读取真实数据', '2026-06-01', 10),
			('Go 语言入门指南', '从零开始学习 Go 语言的基础语法与工程实践', '2026-06-02', 20),
			('Vue 3 Composition API 实战', '深入理解 Vue 3 的组合式 API 与响应式系统', '2026-06-03', 30),
			('Tailwind CSS 最佳实践', '如何高效使用原子化 CSS 框架构建现代化界面', '2026-06-04', 40),
			('从零搭建 Web OS 博客', '用 Vue 3 + Go 从零实现一个仿 macOS 桌面的博客系统', '2026-06-05', 50)
		`)
		log.Println("已插入文章种子数据")
	}

	// 插入种子项目
	DB.QueryRow("SELECT COUNT(*) FROM projects").Scan(&count)
	if count == 0 {
		DB.Exec(`
			INSERT INTO projects (name, description, tech) VALUES
			('TODO Blog', '仿 macOS 桌面的个人博客系统', 'Vue 3 + Go'),
			('Task Manager', '简洁高效的待办事项管理工具', 'React + Node.js'),
			('Weather App', '实时天气查询与预报应用', 'Vue 3 + API')
		`)
		log.Println("已插入项目种子数据")
	}

	// 插入种子个人资料
	DB.QueryRow("SELECT COUNT(*) FROM profile").Scan(&count)
	if count == 0 {
		DB.Exec(`
			INSERT INTO profile (id, name, title, avatar_emoji, bio, tags) VALUES
			(1, 'TUDO', '全栈开发者 / 学生', '🐱',
			 '热爱编程、设计和一切有趣的事物。这个博客将记录我的学习笔记、项目经验和生活思考。🚀',
			 '["Vue","Go","TypeScript","Tailwind CSS","Docker"]')
		`)
		log.Println("已插入个人资料种子数据")
	}

	// 插入种子联系方式
	DB.QueryRow("SELECT COUNT(*) FROM contacts").Scan(&count)
	if count == 0 {
		DB.Exec(`
			INSERT INTO contacts (icon, label, value) VALUES
			('📧', 'Email', 'todo@example.com'),
			('🐙', 'GitHub', 'github.com/My-TuDo'),
			('📱', '微信', 'TODO_WeChat'),
			('🐦', 'Twitter / X', '@TODO_dev')
		`)
		log.Println("已插入联系方式种子数据")
	}
}
