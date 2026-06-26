# TUDO's Blog 🖥️

> 仿 macOS 桌面系统（Web OS）的个人博客前端骨架

## 📋 项目简介

一个基于 **Vue 3 + TypeScript + Tailwind CSS** 构建的个人博客系统，以仿 macOS 桌面为视觉核心。用户进入主页后看到的是一个铺满全屏的电脑桌面，桌面图标点击后弹出"多任务窗口"显示具体内容，无需页面跳转。

## ✨ 核心特性

- 🖼️ **全屏桌面壁纸** — 支持自定义背景图片
- 📱 **桌面快捷方式** — Lucide 图标 + 文字标签，悬浮放大效果
- 🪟 **macOS 风格窗口** — 红黄绿控制按钮、毛玻璃效果（backdrop-blur）
- 🔄 **多任务窗口管理** — 拖拽移动、点击切换 z-index 层级
- 🎬 **流畅动画** — 关闭窗口的 opacity + scale 过渡动画
- 🧩 **动态组件渲染** — 不同窗口内容通过 `<component :is>` 动态加载
- 📦 **响应式状态** — `openWindows[]` 数组管理，方便对接后端 API

## 🛠️ 技术栈

| 技术 | 用途 |
|------|------|
| **Vue 3** (Composition API + `<script setup>`) | 前端框架 |
| **TypeScript** | 类型安全 |
| **Vite** | 构建工具 |
| **Tailwind CSS v4** | 原子化样式 |
| **Lucide Vue Next** | 图标库 |

## 🚀 快速开始

```bash
# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 构建生产版本
npm run build

# 预览构建结果
npm run preview
```

## 📁 项目结构

```
src/
├── main.ts                  # 应用入口
├── style.css                # 全局样式 + Tailwind
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
