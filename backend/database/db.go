package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// Init 初始化 SQLite 数据库，创建表并插入种子数据
func Init() {
	dbPath := "blog.db"

	// 确保 backend 目录存在
	dir := filepath.Dir(dbPath)
	if dir != "." {
		os.MkdirAll(dir, 0755)
	}

	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("打开数据库失败: %v", err)
	}

	// 启用 WAL 模式，提升并发性能
	DB.Exec("PRAGMA journal_mode=WAL")
	// 启用外键约束
	DB.Exec("PRAGMA foreign_keys=ON")

	createTables()
	seedData()
}

// createTables 创建数据库表
func createTables() {
	schema := `
	CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		summary TEXT NOT NULL DEFAULT '',
		created_at TEXT NOT NULL DEFAULT (date('now')),
		views INTEGER NOT NULL DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL DEFAULT '',
		tech TEXT NOT NULL DEFAULT ''
	);

	CREATE TABLE IF NOT EXISTS profile (
		id INTEGER PRIMARY KEY CHECK (id = 1),
		name TEXT NOT NULL DEFAULT 'TUDO',
		title TEXT NOT NULL DEFAULT '',
		avatar_emoji TEXT NOT NULL DEFAULT '🐱',
		bio TEXT NOT NULL DEFAULT '',
		tags TEXT NOT NULL DEFAULT '[]'
	);

	CREATE TABLE IF NOT EXISTS contacts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		icon TEXT NOT NULL,
		label TEXT NOT NULL,
		value TEXT NOT NULL
	);
	`

	_, err := DB.Exec(schema)
	if err != nil {
		log.Fatalf("创建表失败: %v", err)
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
