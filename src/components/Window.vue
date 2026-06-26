<script setup lang="ts">
/**
 * Window.vue - macOS 风格多任务窗口组件
 *
 * 功能：
 * - 仿 macOS 标题栏（红黄绿三色控制按钮）
 * - 标题栏拖拽移动窗口
 * - 点击窗口聚焦（提升 z-index）
 * - 红色关闭按钮带动画（opacity + scale）
 * - macOS 默认主题（白色背景，无毛玻璃）
 * - 动态内容渲染（通过 componentMap + componentName）
 */
import { ref, computed, onMounted, onUnmounted } from 'vue'
import type { WindowState, Position, ComponentMap } from '../types'
import { X, Minus, Square } from 'lucide-vue-next'

const props = defineProps<{
  windowState: WindowState
  componentMap: ComponentMap
  onClose: () => void
  onFocus: () => void
  onPositionChange: (pos: Position) => void
}>()

// ---- 拖拽逻辑 ----
const isDragging = ref(false)
const dragOffset = ref<Position>({ x: 0, y: 0 })
const windowRef = ref<HTMLElement | null>(null)

/** 开始拖拽（仅在标题栏上触发） */
function onDragStart(e: MouseEvent) {
  // 只响应左键
  if (e.button !== 0) return
  isDragging.value = true
  dragOffset.value = {
    x: e.clientX - props.windowState.position.x,
    y: e.clientY - props.windowState.position.y,
  }
  // 拖拽时自动聚焦
  props.onFocus()
  document.addEventListener('mousemove', onDragMove)
  document.addEventListener('mouseup', onDragEnd)
}

/** 拖拽移动 */
function onDragMove(e: MouseEvent) {
  if (!isDragging.value) return
  props.onPositionChange({
    x: e.clientX - dragOffset.value.x,
    y: e.clientY - dragOffset.value.y,
  })
}

/** 结束拖拽 */
function onDragEnd() {
  isDragging.value = false
  document.removeEventListener('mousemove', onDragMove)
  document.removeEventListener('mouseup', onDragEnd)
}

// ---- 关闭动画 ----
const isAnimating = ref(false)

function handleClose() {
  isAnimating.value = true
  // 等 CSS transition 结束后再调用父组件的关闭逻辑
  setTimeout(() => {
    props.onClose()
  }, 200)
}

// ---- 窗口点击聚焦 ----
function handleWindowClick() {
  props.onFocus()
}

// 组件卸载时清理事件
onUnmounted(() => {
  document.removeEventListener('mousemove', onDragMove)
  document.removeEventListener('mouseup', onDragEnd)
})
</script>

<template>
  <div
    ref="windowRef"
    class="window-container absolute"
    :class="{
      'opacity-0 scale-95 transition-all duration-200 ease-in': isAnimating,
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
    <!-- 窗口主体：macOS 默认主题（白色背景 + 圆角 + 阴影） -->
    <div
      class="
        flex flex-col w-full h-full
        rounded-2xl overflow-hidden
        bg-white
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
          bg-gray-50
          border-b border-gray-200/60
          cursor-grab active:cursor-grabbing
        "
        @mousedown="onDragStart"
      >
        <!-- 三色控制按钮 -->
        <div class="flex items-center gap-[6px]">
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
          >
            <X :size="8" class="text-[#4d1c1a] opacity-0 group-hover:opacity-100 transition-opacity" />
          </button>
          <!-- 黄色 - 最小化（预留） -->
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
          >
            <Minus :size="8" class="text-[#5c4a1a] opacity-0 group-hover:opacity-100 transition-opacity" />
          </button>
          <!-- 绿色 - 最大化（预留） -->
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
          >
            <Square :size="7" class="text-[#1a4a1a] opacity-0 group-hover:opacity-100 transition-opacity" />
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
      <div class="flex-1 overflow-auto p-4 text-gray-700">
        <component :is="componentMap[windowState.componentName]" />
      </div>
    </div>
  </div>
</template>
