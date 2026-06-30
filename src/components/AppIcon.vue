<script setup lang="ts">
/**
 * AppIcon.vue - 桌面图标组件（自由拖拽 + 幽灵占位）
 *
 * 拖拽行为（仿 Windows）：
 * - 拖拽时图标自由跟随鼠标，60fps 无跳跃
 * - 目标网格位置显示半透明幽灵占位
 * - 松开鼠标后图标过渡动画飞到网格位置
 */
import type { AppDefinition, Position } from '../types'
import * as LucideIcons from 'lucide-vue-next'
import { computed, ref, h } from 'vue'

const props = defineProps<{
  app: AppDefinition
  position: Position
  cellWidth: number
  cellHeight: number
  gridOffsetX: number
  gridOffsetY: number
}>()

const emit = defineEmits<{
  open: []
  positionChange: [pos: Position]
}>()

// ============ 网格吸附工具函数 ============

function snapToGrid(pos: Position): Position {
  return {
    x: Math.round((pos.x - props.gridOffsetX) / props.cellWidth) * props.cellWidth + props.gridOffsetX,
    y: Math.round((pos.y - props.gridOffsetY) / props.cellHeight) * props.cellHeight + props.gridOffsetY,
  }
}

// ============ 图标渲染 ============

const IconComponent = computed(() => {
  const iconName = props.app.icon as keyof typeof LucideIcons
  return (LucideIcons[iconName] as ReturnType<typeof h>) || LucideIcons.File
})

// ============ 拖拽状态 ============

const isDragging = ref(false)
const isDragMoving = ref(false)
const dragOffset = ref<Position>({ x: 0, y: 0 })
const currentPos = ref<Position>({ ...props.position })

/** 鼠标实时位置（自由跟随，不吸附） */
const freePos = ref<Position | null>(null)
/** 网格目标位置（幽灵占位用） */
const ghostPos = ref<Position | null>(null)

/** 拖拽移动阈值 */
const DRAG_THRESHOLD = 5

function onDragStart(e: MouseEvent) {
  if (e.button !== 0) return
  e.preventDefault()

  isDragging.value = true
  isDragMoving.value = false
  currentPos.value = { ...props.position }
  dragOffset.value = {
    x: e.clientX - props.position.x,
    y: e.clientY - props.position.y,
  }
  freePos.value = { ...props.position }
  ghostPos.value = { ...props.position }

  document.addEventListener('mousemove', onDragMove)
  document.addEventListener('mouseup', onDragEnd)
}

function onDragMove(e: MouseEvent) {
  if (!isDragging.value) return

  const newX = e.clientX - dragOffset.value.x
  const newY = e.clientY - dragOffset.value.y

  // 限制在屏幕范围内
  const clampedX = Math.max(0, Math.min(newX, window.innerWidth - 96))
  const clampedY = Math.max(0, Math.min(newY, window.innerHeight - 112))

  // 自由位置：直接跟随鼠标，不吸附
  freePos.value = { x: clampedX, y: clampedY }

  // 幽灵位置：实时计算最近网格点
  ghostPos.value = snapToGrid({ x: clampedX, y: clampedY })

  // 判断拖拽阈值
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

  if (isDragMoving.value && ghostPos.value) {
    // 提交吸附后的网格位置
    emit('positionChange', { ...ghostPos.value })
  }

  isDragging.value = false
  freePos.value = null
  ghostPos.value = null
  isDragMoving.value = false
}

// ============ 双击打开 ==========

function onDblClick() {
  if (isDragMoving.value) return
  emit('open')
}

// ============ 计算样式 ============

/** 主图标位置 + 层级：拖拽中自由跟随，结束后回到 props.position */
const iconStyle = computed(() => {
  const pos = freePos.value || props.position
  return {
    transform: `translate(${pos.x}px, ${pos.y}px)`,
    zIndex: isDragging.value ? 9999 : undefined,
  }
})

/** 拖拽中禁用 transition，结束后恢复 */
const transitionClass = computed(() => {
  return isDragging.value ? '' : 'transition-all duration-200 ease-out'
})
</script>

<template>
  <!-- 幽灵占位（拖拽中显示，半透明） -->
  <div
    v-if="isDragging && ghostPos"
    class="absolute top-0 left-0 pointer-events-none transition-none"
    :style="{ transform: `translate(${ghostPos.x}px, ${ghostPos.y}px)` }"
  >
    <div
      class="
        flex flex-col items-center justify-center
        w-24 h-28 rounded-xl p-2
        border-2 border-dashed border-white/30
        bg-white/5
      "
    >
      <div class="w-14 h-14 flex items-center justify-center text-white/30">
        <component :is="IconComponent" :size="40" :stroke-width="1.5" />
      </div>
      <span class="mt-1.5 text-xs text-center text-white/30 leading-tight line-clamp-2 px-1">
        {{ app.title }}
      </span>
    </div>
  </div>

  <!-- 主图标 -->
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
      <div class="w-14 h-14 flex items-center justify-center text-white drop-shadow-lg pointer-events-none">
        <component :is="IconComponent" :size="40" :stroke-width="1.5" />
      </div>
      <span
        class="
          mt-1.5 text-xs text-center text-white
          drop-shadow-[0_1px_3px_rgba(0,0,0,0.8)]
          leading-tight line-clamp-2 px-1
          pointer-events-none
        "
      >
        {{ app.title }}
      </span>
    </div>
  </div>
</template>
