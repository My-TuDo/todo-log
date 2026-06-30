# TODO's Blog — 仿 macOS Web OS

> 🎯 总体目标：Vue 3 + Tailwind CSS 仿 macOS 桌面个人博客

---

## ✅ 已完成 — 前端核心骨架

### 项目初始化

- [x] Vite + Vue 3 + TypeScript + Tailwind CSS v4 + Lucide Vue Next
- [x] 清理脚手架，重置入口

### 类型系统 `src/types/index.ts`

- [x] `AppDefinition` — 桌面应用快捷方式定义
- [x] `WindowState` / `Position` / `Size` / `WindowMode` — 窗口运行时状态
- [x] `ComponentMap` — 动态组件映射
- [x] `Article` / `Profile` / `Project` / `ContactItem` — 业务数据
- [x] `LoginResponse` / `ApiError` — API 响应

### 数据层

- [x] `src/mock/data.ts` — 本地 Mock 数据，无需后端即可运行
- [x] `src/services/api.ts` — 统一 API 服务层，`USE_MOCK` 一键切换
- [x] `src/store/index.ts` — 轻量级全局状态（壁纸/图标配置）

### 桌面系统组件 `src/components/`

| 组件 | 职责 | 功能 |
|------|------|------|
| `Desktop.vue` | **状态管理中心** | 窗口增删/聚焦/最小化/最大化，图标位置管理，右键菜单，Dock 集成 |
| `AppIcon.vue` | **桌面图标** | 自由拖拽+幽灵占位+网格吸附，双击打开，支持自定义图片 |
| `Window.vue` | **多任务窗口** | 标题栏拖拽，四边缩放，红黄绿按钮（关闭/最小化/最大化），动态内容 |
| `Dock.vue` | **底部 Dock** | 显示已打开窗口，点击聚焦/还原，最小化半透明 |
| `ContextMenu.vue` | **右键菜单** | 屏幕边缘吸附，打开后台管理/刷新桌面 |

### 窗口内容 `src/components/windows/`

| 组件 | 对应 App | 数据源 |
|------|----------|--------|
| `ArticleList.vue` | 我的文章.app | `fetchArticles()` + 配图 |
| `AboutMe.vue` | 关于我.app | `fetchProfile()` + 头像 |
| `Projects.vue` | 项目展示.app | `fetchProjects()` |
| `Contact.vue` | 联系我.app | `fetchContacts()` |
| `AdminTerminal.vue` | 后台管理.app | 登录 + CRUD + 外观设置 |

### 交互功能

- [x] 桌面图标自由拖拽 + 网格吸附 + 幽灵占位
- [x] 双击图标打开窗口（窗口层叠偏移）
- [x] 窗口标题栏拖拽移动
- [x] 窗口四边/四角拖拽缩放（最小 400×280）
- [x] 红黄绿按钮：关闭动画 / 最小化到 Dock / 最大化还原
- [x] 多窗口 z-index 层级管理（点击聚焦）
- [x] Dock 任务栏 + 最小化指示
- [x] 右键桌面菜单
- [x] 后台管理图形界面（登录 + 4 管理模块 + 外观设置）
- [x] 壁纸 / 图标自定义（localStorage 持久化）
- [x] 所有组件有 loading / error / empty 三态

---

## 🔧 待优化项

- [ ] **替换 `confirm()` 弹窗** — `AdminTerminal.vue` 中删除确认使用浏览器原生 `confirm()`，阻塞 UI，后续应替换为自定义模态框组件
- [ ] **窗口最小尺寸限制** — 当前硬编码在 `Window.vue` 中，可提取为配置常量
- [ ] **拖拽触屏支持** — 当前仅支持鼠标事件，可补充 touch 事件支持平板操作

---

## 📁 前端架构总览

```
src/
│
├── main.ts                        # 应用入口
├── App.vue                        # 根组件（提供应用列表 + 壁纸）
├── style.css                      # 全局样式 + Tailwind 入口
├── vite-env.d.ts                  # Vite 类型声明
│
├── types/
│   └── index.ts                   # ★ 所有 TypeScript 类型定义
│
├── mock/
│   └── data.ts                    # ★ Mock 数据（开发初期用）
│
├── services/
│   └── api.ts                     # ★ API 服务层（对接后端的唯一桥梁）
│
├── store/
│   └── index.ts                   # ★ 轻量级全局状态管理
│
└── components/
    │
    ├── Desktop.vue                # 桌面容器（状态管理中枢）
    ├── AppIcon.vue                # 桌面图标
    ├── Window.vue                 # 窗口管理器
    ├── Dock.vue                   # 底部任务栏
    ├── ContextMenu.vue            # 右键菜单
    │
    └── windows/                   # 窗口内容（每个 .app 对应一个）
        ├── ArticleList.vue
        ├── AboutMe.vue
        ├── Projects.vue
        ├── Contact.vue
        └── AdminTerminal.vue
```

### 数据流

```
App.vue
  └─ Desktop.vue (状态中枢)
       ├─ AppIcon.vue × N     → 双击 → openApp()
       ├─ Window.vue × N      → 关闭/最小化/最大化 → Desktop 方法
       ├─ Dock.vue            → 点击 → restoreFromDock()
       └─ ContextMenu.vue     → 右键 → openApp('admin')

src/store/index.ts  ← 壁纸 / 图标配置（reactive，全局响应）
src/services/api.ts ← 所有 fetch 调用集中在此（USE_MOCK 控制）
src/mock/data.ts    ← Mock 数据（USE_MOCK=true 时生效）
```

---

## 📋 验证清单（全部通过）

- [x] `npm run dev` 启动无报错
- [x] TypeScript 类型检查 `vue-tsc --noEmit` 零错误
- [x] 桌面壁纸全屏铺满，图标网格可见
- [x] 双击图标打开窗口（层叠偏移）
- [x] 窗口拖拽 / 缩放 / 关闭 / 最小化 / 最大化
- [x] 多窗口 z-index 层级切换
- [x] Dock 显示已打开窗口 + 最小化指示
- [x] 右键弹出菜单
- [x] 后台管理登录 + 4 个管理模块 + 外观设置
- [x] Mock 数据可独立运行，无需后端
- [x] 全局加载 / 错误 / 空状态覆盖
- [x] localStorage 持久化壁纸和图标设置
