<script setup lang="ts">
/**
 * AppIcon.vue - 桌面图标组件（支持拖拽）
 *
 * 功能：
 * - 显示应用图标（Lucide）+ 文字标签
 * - 使用 CSS3 transform: translate() 实现位置渲染（硬件加速 60fps）
 * - 原生 mousedown/mousemove/mouseup 拖拽
 * - 拖拽时禁用 transition，结束后恢复
 * - 双击打开窗口（拖拽距离 < 5px 才触发）
 */
import type { AppDefinition, Position } from '../types'
import * as LucideIcons from 'lucide-vue-next'
import { computed, ref, h } from 'vue'

const props = defineProps<{
  app: AppDefinition
  position: Position
}>()

const emit = defineEmits<{
  open: []
  positionChange: [pos: Position]
}>()

// ============ 图标渲染 ============

// 根据 icon 名称动态获取 Lucide 图标组件
const IconComponent = computed(() => {
  const iconName = props.app.icon as keyof typeof LucideIcons
  return (LucideIcons[iconName] as ReturnType<typeof h>) || LucideIcons.File
})

// ============ 拖拽状态 ============

const isDragging = ref(false)
const isDragMoving = ref(false)
const dragOffset = ref<Position>({ x: 0, y: 0 })
const currentPos = ref<Position>({ ...props.position })

// 拖拽时的临时位置（用于实时 transform）
const dragPos = ref<Position | null>(null)

/** 拖拽移动阈值（px），超过此值视为拖拽而非点击 */
const DRAG_THRESHOLD = 5

function onDragStart(e: MouseEvent) {
  // 只响应左键
  if (e.button !== 0) return
  e.preventDefault()

  isDragging.value = true
  isDragMoving.value = false
  currentPos.value = { ...props.position }
  dragOffset.value = {
    x: e.clientX - props.position.x,
    y: e.clientY - props.position.y,
  }
  dragPos.value = { ...props.position }

  document.addEventListener('mousemove', onDragMove)
  document.addEventListener('mouseup', onDragEnd)
}

function onDragMove(e: MouseEvent) {
  if (!isDragging.value) return

  const newX = e.clientX - dragOffset.value.x
  const newY = e.clientY - dragOffset.value.y

  // 限制在屏幕范围内（不超出桌面）
  const clampedX = Math.max(0, Math.min(newX, window.innerWidth - 96))
  const clampedY = Math.max(0, Math.min(newY, window.innerHeight - 112))

  dragPos.value = { x: clampedX, y: clampedY }

  // 判断是否超过拖拽阈值
  const dx = newX - currentPos.value.x
  const dy = newY - currentPos.value.y
  if (Math.sqrt(dx * dx + dy * dy) > DRAG_THRESHOLD) {
    isDragMoving.value = true
  }
}

function onDragEnd() {
  if (!isDragging.value) return

  document.removeEventListener('mousemove', onDragMove)
  document.removeEventListener('mouseup', onDragEnd)

  if (isDragMoving.value && dragPos.value) {
    // 真正发生了拖拽 → 提交最终位置
    emit('positionChange', { ...dragPos.value })
  }

  isDragging.value = false
  dragPos.value = null
  isDragMoving.value = false
}

// ============ 双击打开（拖拽时不触发） ==========

function onDblClick() {
  // 如果刚刚拖拽过，不触发打开
  if (isDragMoving.value) return
  emit('open')
}

// ============ 计算样式 ============

const iconStyle = computed(() => {
  // 使用 transform: translate() 渲染位置，硬件加速
  const pos = dragPos.value || props.position
  return {
    transform: `translate(${pos.x}px, ${pos.y}px)`,
  }
})

const transitionClass = computed(() => {
  // 拖拽过程中禁用 transition，确保即时响应
  // 拖拽结束后恢复过渡动画
  return isDragging.value ? '' : 'transition-all duration-200 ease-out'
})
</script>

<template>
  <div
    class="absolute top-0 left-0 will-change-transform"
    :class="transitionClass"
    :style="iconStyle"
  >
    <div
      class="
        flex flex-col items-center justify-center
        w-24 h-28
        cursor-grab active:cursor-grabbing
        rounded-xl
        hover:scale-105
        hover:bg-white/10
        hover:shadow-lg hover:shadow-white/10
        active:scale-95
        p-2
        select-none
      "
      :class="{ 'scale-105 bg-white/10 shadow-lg shadow-white/10': isDragging }"
      @mousedown="onDragStart"
      @dblclick="onDblClick"
    >
      <!-- 图标 -->
      <div class="w-14 h-14 flex items-center justify-center text-white drop-shadow-lg pointer-events-none">
        <component :is="IconComponent" :size="40" :stroke-width="1.5" />
      </div>
      <!-- 文字标签 -->
      <span
        class="
          mt-1.5 text-xs text-center text-white
          drop-shadow-[0_1px_3px_rgba(0,0,0,0.8)]
          leading-tight
          line-clamp-2
          px-1
          pointer-events-none
        "
      >
        {{ app.title }}
      </span>
    </div>
  </div>
</template>
