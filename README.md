# TUDO's Blog

> 仿 macOS 桌面系统（Web OS）的个人博客 — Vue 3 前端 + Go 后端

## 项目简介

一个以 **仿 macOS 桌面** 为交互核心的全栈个人博客。进入主页看到的是全屏电脑桌面，图标双击弹出多任务窗口，窗口可拖拽/缩放/最小化/最大化。底部 Dock 显示已打开的窗口，桌面支持右键菜单。

后端使用 Go + Gin 框架，提供 RESTful API，当前使用 Mock 数据层，可无缝切换到真实后端。

## 核心特性

### 前端

| 特性 | 说明 |
|------|------|
| **全屏桌面壁纸** | 支持 URL 自定义 + 6 种预设，实时切换 |
| **桌面快捷方式** | 网格吸附布局，自由拖拽（幽灵占位预览） |
| **macOS 风格窗口** | 红黄绿控制按钮，拖拽/缩放/最大最小化 |
| **多任务管理** | z-index 层级，Dock 任务栏，最小化指示 |
| **右键菜单** | 屏幕边缘自适应，快速打开后台管理 |
| **后台管理界面** | 文章/项目管理，个人资料/联系方式编辑 |
| **外观设置** | 壁纸 + 应用图标自定义（localStorage 持久化） |
| **Mock 数据层** | 无需后端即可独立运行开发 |

### 后端

| 特性 | 说明 |
|------|------|
| **Go + Gin** | 轻量高性能 HTTP 框架 |
| **RESTful API** | 文章/项目/简介/联系方式 CRUD |
| **JWT 认证** | 登录验证 + 中间件保护 |
| **SQLite** | 嵌入式数据库，零配置 |

## 技术栈

| 领域 | 技术 |
|------|------|
| 前端框架 | Vue 3 (Composition API + `<script setup lang="ts">`) |
| 构建工具 | Vite |
| 样式 | Tailwind CSS v4 |
| 图标 | Lucide Vue Next |
| 后端语言 | Go |
| Web 框架 | Gin |
| 数据库 | SQLite (via GORM / 原生驱动) |
| 认证 | JWT |

## 快速开始

### 前端

```bash
npm install
npm run dev      # http://localhost:5173
```

### 后端

```bash
cd backend
go mod tidy
go run .         # http://localhost:8080
```

> 前端默认使用 Mock 数据，无需启动后端即可开发。切换到后端数据请修改 `src/services/api.ts` 中的 `USE_MOCK = false`。

## 项目结构

```
├── index.html
├── package.json
├── vite.config.ts
├── tsconfig.json
│
├── src/                          # 前端源码
│   ├── main.ts                   # 应用入口
│   ├── App.vue                   # 根组件
│   ├── style.css                 # 全局样式 + Tailwind
│   ├── types/index.ts            # 所有 TypeScript 类型定义
│   ├── mock/data.ts              # Mock 数据层
│   ├── services/api.ts           # API 服务层（USE_MOCK 开关）
│   ├── store/index.ts            # 全局状态管理
│   │
│   └── components/
│       ├── Desktop.vue           # 桌面容器（状态管理中枢）
│       ├── AppIcon.vue           # 桌面图标（自由拖拽 + 幽灵占位）
│       ├── Window.vue            # 窗口管理器
│       ├── Dock.vue              # 底部任务栏
│       ├── ContextMenu.vue       # 右键菜单
│       └── windows/              # 窗口内容
│           ├── ArticleList.vue
│           ├── AboutMe.vue
│           ├── Projects.vue
│           ├── Contact.vue
│           └── AdminTerminal.vue
│
├── backend/                      # 后端源码
│   ├── main.go                   # 入口
│   ├── go.mod / go.sum
│   ├── config/config.go          # 配置
│   ├── database/db.go            # 数据库初始化
│   ├── middleware/auth.go        # JWT 认证中间件
│   ├── models/                   # 数据模型
│   ├── handlers/                 # HTTP handler
│   │   ├── articles.go
│   │   ├── projects.go
│   │   ├── profile.go
│   │   ├── contact.go
│   │   └── auth.go
│   └── routes/routes.go          # 路由注册
│
└── .github/agents/               # VS Code Agent 配置（已 gitignore）
```

## API 接口

| 方法 | 路径 | 说明 | 需认证 |
|------|------|------|--------|
| GET | `/api/articles` | 获取文章列表 | 否 |
| POST | `/api/articles` | 创建文章 | 是 |
| DELETE | `/api/articles/:id` | 删除文章 | 是 |
| GET | `/api/projects` | 获取项目列表 | 否 |
| POST | `/api/projects` | 创建项目 | 是 |
| DELETE | `/api/projects/:id` | 删除项目 | 是 |
| GET | `/api/profile` | 获取个人简介 | 否 |
| PUT | `/api/profile` | 更新个人简介 | 是 |
| GET | `/api/contact` | 获取联系方式 | 否 |
| PUT | `/api/contact` | 更新联系方式 | 是 |
| POST | `/api/login` | 登录获取 JWT | 否 |
├── App.vue                  # 根组件（应用配置）
├── types/
│   └── index.ts             # 核心类型定义
└── components/
    ├── Desktop.vue          # 桌面布局 + 窗口状态管理
    ├── AppIcon.vue          # 桌面图标组件
    ├── Window.vue           # macOS 风格窗口
    └── windows/             # 窗口内容组件
        ├── ArticleList.vue  # 文章列表
        ├── AboutMe.vue      # 关于我
        ├── Projects.vue     # 项目展示
        └── Contact.vue      # 联系方式
```

## 🧩 核心类型

```typescript
// 桌面应用定义
interface AppDefinition {
  id: string
  title: string       // 例如 "我的文章.app"
  icon: string        // Lucide 图标名
  componentName: string // 窗口组件名
}

// 窗口运行状态
interface WindowState {
  id: string
  title: string
  icon: string
  componentName: string
  zIndex: number
  position: Position
  size: Size
  isClosing: boolean
}
```

## 📌 待办

- [ ] 对接 Go 后端 API
- [ ] 窗口最小化/最大化功能
- [ ] Dock 任务栏
- [ ] 右键菜单
- [ ] 桌面文件拖拽

---

<p align="center">Made with ❤️ by TUDO</p>
