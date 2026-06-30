<script setup lang="ts">
/**
 * Desktop.vue - 桌面布局组件（状态管理中心）
 *
 * 功能：
 * - 全屏桌面容器，支持自定义壁纸
 * - 管理桌面图标网格
 * - 管理"当前打开了哪些窗口"（openWindows 响应式数组）
 * - 右键桌面弹出上下文菜单
 * - 底部 Dock 显示所有打开的窗口
 * - 提供打开/关闭/聚焦/最小化/最大化/拖拽/缩放窗口的方法
 *
 * 数据流：
 *   App.vue
 *     └─ Desktop.vue (管理 openWindows[], nextZIndex)
 *          ├─ AppIcon.vue × N       (点击 → openApp)
 *          ├─ Window.vue × N        (渲染自 openWindows[])
 *          ├─ Dock.vue              (显示所有窗口，点击聚焦/还原)
 *          └─ ContextMenu.vue       (右键菜单，点击外部关闭)
 */
import { ref } from 'vue'
import type { AppDefinition, WindowState, Position, Size, ComponentMap } from '../types'
import AppIcon from './AppIcon.vue'
import Window from './Window.vue'
import Dock from './Dock.vue'
import ContextMenu from './ContextMenu.vue'

// 导入所有窗口内容组件
import ArticleList from './windows/ArticleList.vue'
import AboutMe from './windows/AboutMe.vue'
import Projects from './windows/Projects.vue'
import Contact from './windows/Contact.vue'
import AdminTerminal from './windows/AdminTerminal.vue'

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
  AdminTerminal,
}

/** 右键菜单状态 */
const contextMenu = ref<{ x: number; y: number } | null>(null)

// ============ 桌面图标位置管理 ============

/** 每个桌面图标的位置（id -> Position），使用 transform 渲染 */
const iconPositions = ref<Record<string, Position>>({})

/**
 * 网格布局参数（Windows 风格）
 * 所有图标自动吸附到网格线，保证行列对齐
 */
const CELL_WIDTH = 100        // 网格单元宽度（px）
const CELL_HEIGHT = 120       // 网格单元高度（px）
const GRID_OFFSET_X = 24      // 网格起始 X 偏移
const GRID_OFFSET_Y = 24      // 网格起始 Y 偏移
const GRID_COLS = 8           // 每行图标数（超出自动折行）

/** 将任意位置吸附到最近的网格点 */
function snapToGrid(pos: Position): Position {
  return {
    x: Math.round((pos.x - GRID_OFFSET_X) / CELL_WIDTH) * CELL_WIDTH + GRID_OFFSET_X,
    y: Math.round((pos.y - GRID_OFFSET_Y) / CELL_HEIGHT) * CELL_HEIGHT + GRID_OFFSET_Y,
  }
}

/** 初始化图标位置为网格布局 */
function initIconPositions() {
  const positions: Record<string, Position> = {}
  props.apps.forEach((app, index) => {
    const col = index % GRID_COLS
    const row = Math.floor(index / GRID_COLS)
    positions[app.id] = {
      x: GRID_OFFSET_X + col * CELL_WIDTH,
      y: GRID_OFFSET_Y + row * CELL_HEIGHT,
    }
  })
  iconPositions.value = positions
}

/** 更新某个图标的位置（拖拽完成后调用，自动吸附到网格） */
function updateIconPosition(id: string, position: Position) {
  iconPositions.value[id] = snapToGrid(position)
}

// 初始化位置
initIconPositions()

// ============ 窗口操作方法 ============

/**
 * 打开一个应用
 * - 如果窗口已存在且最小化，则还原并聚焦
 * - 如果已存在且未最小化，直接聚焦
 * - 如果不存在，创建一个新窗口
 */
function openApp(app: AppDefinition) {
  const existing = openWindows.value.find((w) => w.id === app.id)
  if (existing) {
    // 如果是最小化状态，还原为 normal
    if (existing.mode === 'minimized') {
      existing.mode = 'normal'
      // 还原到之前保存的位置和尺寸
      if (existing.previousPosition) {
        existing.position = { ...existing.previousPosition }
      }
      if (existing.previousSize) {
        existing.size = { ...existing.previousSize }
      }
    }
    focusWindow(app.id)
    return
  }

  // 计算偏移位置：新窗口右下微移，形成层叠效果
  const count = openWindows.value.length
  const offset = (count % 5) * 30
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
    mode: 'normal',
    isClosing: false,
  }

  openWindows.value.push(newWindow)
}

/**
 * 关闭窗口
 */
function closeWindow(id: string) {
  openWindows.value = openWindows.value.filter((w) => w.id !== id)
}

/**
 * 聚焦窗口：提升到最高层级
 */
function focusWindow(id: string) {
  const win = openWindows.value.find((w) => w.id === id)
  if (win) {
    win.zIndex = nextZIndex.value++
  }
}

/**
 * 更新窗口位置
 */
function updateWindowPosition(id: string, position: Position) {
  const win = openWindows.value.find((w) => w.id === id)
  if (win) {
    win.position = position
  }
}

/**
 * 更新窗口尺寸（拖拽缩放时调用）
 */
function updateWindowSize(id: string, size: Size) {
  const win = openWindows.value.find((w) => w.id === id)
  if (win) {
    win.size = size
  }
}

/**
 * 最小化窗口
 */
function minimizeWindow(id: string) {
  const win = openWindows.value.find((w) => w.id === id)
  if (win && win.mode !== 'minimized') {
    // 保存当前状态以便还原
    win.previousPosition = { ...win.position }
    win.previousSize = { ...win.size }
    win.mode = 'minimized'
  }
}

/**
 * 最大化 / 还原窗口
 */
function maximizeWindow(id: string) {
  const win = openWindows.value.find((w) => w.id === id)
  if (!win) return

  if (win.mode === 'maximized') {
    // 还原
    win.mode = 'normal'
    if (win.previousPosition) win.position = { ...win.previousPosition }
    if (win.previousSize) win.size = { ...win.previousSize }
  } else {
    // 最大化
    win.previousPosition = { ...win.position }
    win.previousSize = { ...win.size }
    win.mode = 'maximized'
    win.position = { x: 0, y: 0 }
    win.size = { width: window.innerWidth, height: window.innerHeight }
  }
  focusWindow(id)
}

/**
 * 从 Dock 还原（聚焦）窗口
 * 如果窗口已最小化，先还原为 normal
 */
function restoreFromDock(id: string) {
  const win = openWindows.value.find((w) => w.id === id)
  if (!win) return

  if (win.mode === 'minimized') {
    win.mode = 'normal'
    if (win.previousPosition) win.position = { ...win.previousPosition }
    if (win.previousSize) win.size = { ...win.previousSize }
  }
  focusWindow(id)
}

// ============ 右键菜单 ============

function handleContextMenu(e: MouseEvent) {
  e.preventDefault()
  contextMenu.value = { x: e.clientX, y: e.clientY }
}

function closeContextMenu() {
  contextMenu.value = null
}

/** 右键菜单 → 打开指定应用 */
function handleContextMenuOpenApp(appId: string) {
  closeContextMenu()
  const app = props.apps.find((a) => a.id === appId)
  if (app) openApp(app)
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
    @contextmenu="handleContextMenu"
  >
    <!-- 桌面图标层（网格吸附，支持拖拽） -->
    <AppIcon
      v-for="app in apps"
      :key="app.id"
      :app="app"
      :position="iconPositions[app.id] ?? { x: 0, y: 0 }"
      :cell-width="CELL_WIDTH"
      :cell-height="CELL_HEIGHT"
      :grid-offset-x="GRID_OFFSET_X"
      :grid-offset-y="GRID_OFFSET_Y"
      @open="openApp(app)"
      @position-change="(pos: Position) => updateIconPosition(app.id, pos)"
    />

    <!-- 窗口层 -->
    <Window
      v-for="win in openWindows"
      :key="win.id"
      :window-state="win"
      :component-map="componentMap"
      :on-close="() => closeWindow(win.id)"
      :on-focus="() => focusWindow(win.id)"
      :on-minimize="() => minimizeWindow(win.id)"
      :on-maximize="() => maximizeWindow(win.id)"
      :on-position-change="(pos: Position) => updateWindowPosition(win.id, pos)"
      :on-size-change="(size: Size) => updateWindowSize(win.id, size)"
    />

    <!-- Dock 任务栏 -->
    <Dock
      :open-windows="openWindows"
      :apps="apps"
      @restore="restoreFromDock"
    />

    <!-- 右键菜单 -->
    <ContextMenu
      v-if="contextMenu"
      :x="contextMenu.x"
      :y="contextMenu.y"
      @close="closeContextMenu"
      @open-app="handleContextMenuOpenApp"
    />
  </div>
</template>
