<script setup lang="ts">
/**
 * Desktop.vue - 桌面布局组件（状态管理中心）
 *
 * 功能：
 * - 全屏桌面容器，支持自定义壁纸
 * - 管理桌面图标网格
 * - 管理"当前打开了哪些窗口"（openWindows 响应式数组）
 * - 提供打开/关闭/聚焦/拖拽更新窗口的方法
 *
 * 数据流：
 *   App.vue
 *     └─ Desktop.vue (管理 openWindows[])
 *          ├─ AppIcon.vue × N    (点击 → openApp)
 *          └─ Window.vue × N     (渲染自 openWindows[])
 */
import { ref } from 'vue'
import type { AppDefinition, WindowState, Position, ComponentMap } from '../types'
import AppIcon from './AppIcon.vue'
import Window from './Window.vue'

// 导入所有窗口内容组件
import ArticleList from './windows/ArticleList.vue'
import AboutMe from './windows/AboutMe.vue'
import Projects from './windows/Projects.vue'
import Contact from './windows/Contact.vue'

const props = defineProps<{
  /** 壁纸图片 URL */
  wallpaper: string
  /** 桌面快捷方式列表 */
  apps: AppDefinition[]
}>()

// ============ 状态管理 ============

/** 当前已打开的窗口列表 */
const openWindows = ref<WindowState[]>([])

/** 全局 z-index 计数器：每次打开/聚焦窗口时递增 */
const nextZIndex = ref(10)

/** 动态组件映射表：componentName -> Vue 组件 */
const componentMap: ComponentMap = {
  ArticleList,
  AboutMe,
  Projects,
  Contact,
}

// ============ 窗口操作方法 ============

/**
 * 打开一个应用
 * - 如果窗口已存在，则直接聚焦
 * - 如果不存在，创建一个新窗口并添加到列表
 * - 为新窗口分配一个随机偏移位置，避免多个窗口完全重叠
 */
function openApp(app: AppDefinition) {
  const existing = openWindows.value.find((w) => w.id === app.id)
  if (existing) {
    // 窗口已打开，只需聚焦
    focusWindow(app.id)
    return
  }

  // 计算偏移位置：每次新窗口向右下微移，形成层叠效果
  const offset = (openWindows.value.length % 5) * 30
  const centerX = (window.innerWidth - 640) / 2 + offset
  const centerY = (window.innerHeight - 440) / 2 + offset

  const newWindow: WindowState = {
    id: app.id,
    title: app.title,
    icon: app.icon,
    componentName: app.componentName,
    zIndex: nextZIndex.value++,
    position: { x: centerX, y: centerY },
    size: { width: 640, height: 440 },
    isClosing: false,
  }

  openWindows.value.push(newWindow)
}

/**
 * 关闭窗口（带动画）
 * 先标记 isClosing 触发关闭动画，再由 Window.vue 在动画结束后调用此方法
 */
function closeWindow(id: string) {
  // 从数组中移除
  openWindows.value = openWindows.value.filter((w) => w.id !== id)
}

/**
 * 聚焦窗口：将指定窗口的 z-index 提升到最高
 */
function focusWindow(id: string) {
  const window = openWindows.value.find((w) => w.id === id)
  if (window) {
    window.zIndex = nextZIndex.value++
  }
}

/**
 * 更新窗口位置（拖拽时调用）
 */
function updateWindowPosition(id: string, position: Position) {
  const window = openWindows.value.find((w) => w.id === id)
  if (window) {
    window.position = position
  }
}
</script>

<template>
  <!-- 全屏桌面容器 -->
  <div
    class="relative h-screen w-screen overflow-hidden"
    :style="{
      backgroundImage: `url(${wallpaper})`,
      backgroundSize: 'cover',
      backgroundPosition: 'center',
      backgroundRepeat: 'no-repeat',
    }"
  >
    <!-- 桌面图标网格 -->
    <div class="absolute top-6 left-6 flex flex-wrap gap-2 max-w-[200px]">
      <AppIcon
        v-for="app in apps"
        :key="app.id"
        :app="app"
        :on-open="() => openApp(app)"
      />
    </div>

    <!-- 窗口层：遍历 openWindows 渲染 Window 组件 -->
    <Window
      v-for="win in openWindows"
      :key="win.id"
      :window-state="win"
      :component-map="componentMap"
      :on-close="() => closeWindow(win.id)"
      :on-focus="() => focusWindow(win.id)"
      :on-position-change="(pos: Position) => updateWindowPosition(win.id, pos)"
    />

    <!-- 底部信息（可选） -->
    <div
      class="
        absolute bottom-4 left-1/2 -translate-x-1/2
        px-4 py-1.5 rounded-full
        bg-black/30 backdrop-blur-md
        text-xs text-white/50
        border border-white/10
      "
    >
      TUDO's blog · macOS Web OS
    </div>
  </div>
</template>
