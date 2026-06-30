<script setup lang="ts">
/**
 * Window.vue - macOS 风格多任务窗口组件
 *
 * 功能：
 * - 仿 macOS 标题栏（红黄绿三色控制按钮）
 * - 标题栏拖拽移动窗口
 * - 点击窗口聚焦（提升 z-index）
 * - 红色关闭按钮 → opacity 动画 → 关闭
 * - 黄色最小化按钮 → 隐藏到 Dock
 * - 绿色最大化/还原按钮
 * - 四边/四角拖拽缩放
 * - 动态内容渲染（通过 componentMap + componentName）
 */
import { ref, computed, onUnmounted } from 'vue'
import type { WindowState, Position, Size, ComponentMap } from '../types'
import { X, Minus, Square } from 'lucide-vue-next'

const props = defineProps<{
  windowState: WindowState
  componentMap: ComponentMap
  onClose: () => void
  onFocus: () => void
  onMinimize: () => void
  onMaximize: () => void
  onPositionChange: (pos: Position) => void
  onSizeChange: (size: Size) => void
}>()

// ---- 窗口聚焦 ----
function handleWindowClick() {
  props.onFocus()
}

// ============================================================
// 拖拽移动（仅标题栏）
// ============================================================

const isDragging = ref(false)
const dragOffset = ref<Position>({ x: 0, y: 0 })

function onDragStart(e: MouseEvent) {
  if (e.button !== 0) return
  // 最大化状态下不允许拖拽
  if (props.windowState.mode === 'maximized') return
  isDragging.value = true
  dragOffset.value = {
    x: e.clientX - props.windowState.position.x,
    y: e.clientY - props.windowState.position.y,
  }
  props.onFocus()
  document.addEventListener('mousemove', onDragMove)
  document.addEventListener('mouseup', onDragEnd)
}

function onDragMove(e: MouseEvent) {
  if (!isDragging.value) return
  props.onPositionChange({
    x: e.clientX - dragOffset.value.x,
    y: e.clientY - dragOffset.value.y,
  })
}

function onDragEnd() {
  isDragging.value = false
  document.removeEventListener('mousemove', onDragMove)
  document.removeEventListener('mouseup', onDragEnd)
}

// ============================================================
// 边缘拖拽缩放
// ============================================================

type ResizeDirection = 'n' | 's' | 'e' | 'w' | 'ne' | 'nw' | 'se' | 'sw' | null

const isResizing = ref(false)
const resizeDirection = ref<ResizeDirection>(null)
const resizeStart = ref<{ pos: Position; size: Size; mouse: Position }>({
  pos: { x: 0, y: 0 },
  size: { width: 0, height: 0 },
  mouse: { x: 0, y: 0 },
})

const MIN_WIDTH = 400
const MIN_HEIGHT = 280

function onResizeStart(e: MouseEvent, direction: ResizeDirection) {
  if (e.button !== 0) return
  if (props.windowState.mode === 'maximized') return
  e.stopPropagation()
  isResizing.value = true
  resizeDirection.value = direction
  resizeStart.value = {
    pos: { ...props.windowState.position },
    size: { ...props.windowState.size },
    mouse: { x: e.clientX, y: e.clientY },
  }
  props.onFocus()
  document.addEventListener('mousemove', onResizeMove)
  document.addEventListener('mouseup', onResizeEnd)
}

function onResizeMove(e: MouseEvent) {
  if (!isResizing.value || !resizeDirection.value) return
  const dir = resizeDirection.value
  const dx = e.clientX - resizeStart.value.mouse.x
  const dy = e.clientY - resizeStart.value.mouse.y
  let { x, y } = resizeStart.value.pos
  let { width, height } = resizeStart.value.size

  // 右边界
  if (dir.includes('e')) {
    width = Math.max(MIN_WIDTH, resizeStart.value.size.width + dx)
  }
  // 左边界
  if (dir.includes('w')) {
    const newWidth = Math.max(MIN_WIDTH, resizeStart.value.size.width - dx)
    x = resizeStart.value.pos.x + resizeStart.value.size.width - newWidth
    width = newWidth
  }
  // 下边界
  if (dir.includes('s')) {
    height = Math.max(MIN_HEIGHT, resizeStart.value.size.height + dy)
  }
  // 上边界
  if (dir.includes('n')) {
    const newHeight = Math.max(MIN_HEIGHT, resizeStart.value.size.height - dy)
    y = resizeStart.value.pos.y + resizeStart.value.size.height - newHeight
    height = newHeight
  }

  props.onPositionChange({ x, y })
  props.onSizeChange({ width, height })
}

function onResizeEnd() {
  isResizing.value = false
  resizeDirection.value = null
  document.removeEventListener('mousemove', onResizeMove)
  document.removeEventListener('mouseup', onResizeEnd)
}

// ============================================================
// 三色按钮功能
// ============================================================

/** 红色 - 关闭窗口（带动画） */
const isAnimating = ref(false)

function handleClose() {
  isAnimating.value = true
  setTimeout(() => {
    props.onClose()
  }, 200)
}

/** 黄色 - 最小化 */
function handleMinimize() {
  props.onMinimize()
}

/** 绿色 - 最大化/还原 */
function handleMaximize() {
  props.onMaximize()
}

// ---- 计算最大化/还原图标 ----
const maximizeIcon = computed(() => {
  return props.windowState.mode === 'maximized' ? Minus : Square
})

// ============================================================
// 清理
// ============================================================

onUnmounted(() => {
  document.removeEventListener('mousemove', onDragMove)
  document.removeEventListener('mouseup', onDragEnd)
  document.removeEventListener('mousemove', onResizeMove)
  document.removeEventListener('mouseup', onResizeEnd)
})
</script>

<template>
  <div
    ref="windowRef"
    class="window-container absolute"
    :class="{
      'opacity-0 scale-95 transition-all duration-200 ease-in': isAnimating,
      'invisible': windowState.mode === 'minimized',
    }"
    :style="{
      left: windowState.position.x + 'px',
      top: windowState.position.y + 'px',
      width: windowState.size.width + 'px',
      height: windowState.size.height + 'px',
      zIndex: windowState.zIndex,
    }"
    @mousedown="handleWindowClick"
  >
    <!-- 窗口主体：macOS 默认主题 -->
    <div
      class="
        relative flex flex-col w-full h-full
        rounded-2xl overflow-hidden
        bg-white/95 backdrop-blur-xl
        border border-gray-200/70
        shadow-2xl shadow-black/25
      "
    >
      <!-- 标题栏 -->
      <div
        class="
          drag-handle
          relative flex items-center
          h-10 min-h-10
          px-3
          bg-gray-50/90
          border-b border-gray-200/60
          cursor-grab active:cursor-grabbing
          select-none
        "
        @mousedown="onDragStart"
      >
        <!-- 三色控制按钮 -->
        <div class="flex items-center gap-[6px] z-10">
          <!-- 红色 - 关闭 -->
          <button
            class="
              w-3 h-3 rounded-full
              bg-[#FF5F57]
              hover:brightness-90
              active:brightness-75
              flex items-center justify-center
              transition-all duration-150
              group
            "
            @mousedown.stop
            @click.stop="handleClose"
            title="关闭"
          >
            <X :size="8" class="text-[#4d1c1a] opacity-0 group-hover:opacity-100 transition-opacity" />
          </button>
          <!-- 黄色 - 最小化 -->
          <button
            class="
              w-3 h-3 rounded-full
              bg-[#FEBC2E]
              hover:brightness-90
              active:brightness-75
              flex items-center justify-center
              transition-all duration-150
              group
            "
            @mousedown.stop
            @click.stop="handleMinimize"
            title="最小化"
          >
            <Minus :size="8" class="text-[#5c4a1a] opacity-0 group-hover:opacity-100 transition-opacity" />
          </button>
          <!-- 绿色 - 最大化/还原 -->
          <button
            class="
              w-3 h-3 rounded-full
              bg-[#28C840]
              hover:brightness-90
              active:brightness-75
              flex items-center justify-center
              transition-all duration-150
              group
            "
            @mousedown.stop
            @click.stop="handleMaximize"
            :title="windowState.mode === 'maximized' ? '还原' : '最大化'"
          >
            <component :is="maximizeIcon" :size="7" class="text-[#1a4a1a] opacity-0 group-hover:opacity-100 transition-opacity" />
          </button>
        </div>

        <!-- 窗口标题 -->
        <div class="absolute inset-0 flex items-center justify-center pointer-events-none">
          <span class="text-xs text-gray-600 font-semibold truncate max-w-[60%]">
            {{ windowState.title }}
          </span>
        </div>
      </div>

      <!-- 窗口内容区 -->
      <div class="flex-1 overflow-auto p-4 text-gray-700 bg-white">
        <component :is="componentMap[windowState.componentName]" />
      </div>

      <!-- 四边缩放热区 -->
      <!-- 上 -->
      <div
        class="absolute top-0 left-3 right-3 h-1.5 cursor-n-resize z-20 hover:bg-blue-400/20"
        @mousedown.stop="onResizeStart($event, 'n')"
      />
      <!-- 下 -->
      <div
        class="absolute bottom-0 left-3 right-3 h-1.5 cursor-s-resize z-20 hover:bg-blue-400/20"
        @mousedown.stop="onResizeStart($event, 's')"
      />
      <!-- 左 -->
      <div
        class="absolute left-0 top-3 bottom-3 w-1.5 cursor-w-resize z-20 hover:bg-blue-400/20"
        @mousedown.stop="onResizeStart($event, 'w')"
      />
      <!-- 右 -->
      <div
        class="absolute right-0 top-3 bottom-3 w-1.5 cursor-e-resize z-20 hover:bg-blue-400/20"
        @mousedown.stop="onResizeStart($event, 'e')"
      />
      <!-- 四角 -->
      <div
        class="absolute top-0 left-0 w-3 h-3 cursor-nw-resize z-20 hover:bg-blue-400/20 rounded-tl-xl"
        @mousedown.stop="onResizeStart($event, 'nw')"
      />
      <div
        class="absolute top-0 right-0 w-3 h-3 cursor-ne-resize z-20 hover:bg-blue-400/20 rounded-tr-xl"
        @mousedown.stop="onResizeStart($event, 'ne')"
      />
      <div
        class="absolute bottom-0 left-0 w-3 h-3 cursor-sw-resize z-20 hover:bg-blue-400/20 rounded-bl-xl"
        @mousedown.stop="onResizeStart($event, 'sw')"
      />
      <div
        class="absolute bottom-0 right-0 w-3 h-3 cursor-se-resize z-20 hover:bg-blue-400/20 rounded-br-xl"
        @mousedown.stop="onResizeStart($event, 'se')"
      />
    </div>
  </div>
</template>
