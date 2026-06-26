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
- [x] **Step 7**: 创建 `src/components/AppIcon.vue` — 桌面图标组件
- [x] **Step 8**: 创建 `src/components/Window.vue` — macOS 风格窗口（拖拽、层级、关闭动画）
- [x] **Step 9**: 创建 `src/components/windows/` 目录下的演示窗口内容组件
  - [x] `ArticleList.vue` — "我的文章.app"
  - [x] `AboutMe.vue` — "关于我.app"
  - [x] `Projects.vue` — "项目展示.app"
  - [x] `Contact.vue` — "联系我.app"

## ✅ 阶段四：根组件与入口

- [x] **Step 10**: 修改 `src/App.vue` — 集成 Desktop 并传入演示数据
- [x] **Step 11**: 更新 `index.html` 与 `src/main.ts`

## ✅ 阶段五：验证

- [x] **Step 12**: `npm run dev` 启动并验证 8 项检查点

---

## 📋 验证清单

- [x] 1. `npm run dev` 启动无报错 ✅
- [x] 2. 桌面背景图全屏铺满，图标网格可见 ✅
- [x] 3. 单击 AppIcon 打开对应窗口（窗口从屏幕中央弹出）✅
- [x] 4. 窗口可拖拽（仅标题栏区域），可正常拖动 ✅
- [x] 5. 多窗口点击切换 z-index 层级 ✅
- [x] 6. 红色关闭按钮 → opacity + scale 动画 → 窗口移除 ✅
- [x] 7. `openWindows` 数组随打开/关闭正确增删 ✅
- [x] 8. TypeScript 类型检查通过（`vue-tsc --noEmit`）
