# TODO's Blog — 仿 macOS Web OS 前端骨架

> 🎯 总体目标：用 Vue 3 + Tailwind CSS 实现仿 macOS 桌面的个人博客前端核心骨架

---

## ✅ 阶段一：项目初始化

- [x] **Step 1**: 使用 Vite 创建 Vue 3 + TypeScript 项目
- [x] **Step 2**: 安装并配置 Tailwind CSS v4
- [x] **Step 3**: 安装 Lucide Vue Next 图标库
- [x] **Step 4**: 清理脚手架默认文件，重置入口

## ✅ 阶段二：类型定义

- [x] **Step 5**: 创建 `src/types/index.ts` — 定义 AppDefinition、WindowState、Position 等核心类型

## ✅ 阶段三：核心组件实现

- [x] **Step 6**: 创建 `src/components/Desktop.vue` — 桌面布局 + 窗口状态管理
- [x] **Step 7**: 创建 `src/components/AppIcon.vue` — 桌面图标组件（双击打开）
- [x] **Step 8**: 创建 `src/components/Window.vue` — macOS 风格窗口（拖拽、层级、关闭动画 + 缩放 + 最大化/最小化）
- [x] **Step 9**: 创建 `src/components/windows/` 目录下的演示窗口内容组件
  - [x] `ArticleList.vue` — "我的文章.app"
  - [x] `AboutMe.vue` — "关于我.app"
  - [x] `Projects.vue` — "项目展示.app"
  - [x] `Contact.vue` — "联系我.app"

## ✅ 阶段四：根组件与入口

- [x] **Step 10**: 修改 `src/App.vue` — 集成 Desktop 并传入演示数据
- [x] **Step 11**: 更新 `index.html` 与 `src/main.ts`

## ✅ 阶段五：重构优化 (2026-06-30)

- [x] **R1**: 重构 `types/index.ts` — 添加 WindowMode、Article、Profile、Project、ContactItem 等类型
- [x] **R2**: 创建 `src/mock/data.ts` — 本地 Mock 数据层，开发初期无需后端
- [x] **R3**: 创建 `src/services/api.ts` — 统一 API 服务层，USE_MOCK 开关一键切换
- [x] **R4**: 重构 `Window.vue` — 添加四边缩放、最大化/最小化功能
- [x] **R5**: 重构 `Desktop.vue` — 集成 Dock + ContextMenu，窗口模式管理
- [x] **R6**: 重构 `Dock.vue` — 使用 mode 属性代替 isMinimized
- [x] **R7**: 重构 `AppIcon.vue` — 双击打开，emit 事件代替 prop callback
- [x] **R8**: 重构 `ContextMenu.vue` — 屏幕边缘吸附，emits openApp
- [x] **R9**: 重写窗口内容组件 — 使用 API 服务层 + 加载/错误/空状态
- [x] **R10**: 重写 `AdminTerminal.vue` — 内联交互输入代替 prompt()，使用 API 服务层

---

## 📋 项目结构

```
src/
├── types/index.ts          # 集中类型定义
├── mock/data.ts            # Mock 数据
├── services/api.ts         # API 服务层
├── components/
│   ├── Desktop.vue         # 桌面状态管理中心
│   ├── AppIcon.vue         # 桌面图标（双击打开）
│   ├── Window.vue          # 窗口（拖拽/缩放/最大最小化）
│   ├── Dock.vue            # Dock 任务栏
│   ├── ContextMenu.vue     # 右键菜单
│   └── windows/
│       ├── ArticleList.vue # 文章列表
│       ├── AboutMe.vue     # 个人简介
│       ├── Projects.vue    # 项目展示
│       ├── Contact.vue     # 联系方式
│       └── AdminTerminal.vue # 管理终端
```

## 📋 验证清单

- [x] 1. `npm run dev` 启动无报错 ✅
- [x] 2. 桌面背景图全屏铺满，图标网格可见 ✅
- [x] 3. 双击 AppIcon 打开对应窗口（窗口从屏幕中央弹出）✅
- [x] 4. 窗口可拖拽（仅标题栏区域），可正常拖动 ✅
- [x] 5. 多窗口点击切换 z-index 层级 ✅
- [x] 6. 红色关闭按钮 → opacity + scale 动画 → 窗口移除 ✅
- [x] 7. `openWindows` 数组随打开/关闭正确增删 ✅
- [x] 8. TypeScript 类型检查通过（`vue-tsc --noEmit`）✅
- [x] 9. 窗口四边/四角拖拽缩放 ✅
- [x] 10. 绿色最大化/还原按钮 ✅
- [x] 11. 黄色最小化 → Dock 显示半透明图标 ✅
- [x] 12. Dock 点击还原最小化窗口 ✅
- [x] 13. 右键桌面弹出菜单 ✅
- [x] 14. Mock 数据层支持无后端开发 ✅
- [x] 15. 所有组件加载/错误/空状态覆盖 ✅
- [x] 16. 内联终端输入代替 prompt() ✅
