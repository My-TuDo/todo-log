<script setup lang="ts">
/**
 * Dock.vue - macOS 风格 Dock 任务栏
 *
 * 功能：
 * - 固定在屏幕底部中央
 * - 显示所有已打开窗口的图标
 * - 最小化（minimized）的窗口图标显示为半透明 + 无指示点
 * - 点击图标还原/聚焦窗口
 * - 毛玻璃背景，悬浮放大效果
 */
import { computed } from 'vue'
import type { WindowState, AppDefinition } from '../types'
import * as LucideIcons from 'lucide-vue-next'

const props = defineProps<{
  openWindows: WindowState[]
  apps: AppDefinition[]
}>()

const emit = defineEmits<{
  restore: [id: string]
}>()

/** 合并窗口与 app 信息，显示在 Dock 中 */
const dockItems = computed(() => {
  return props.openWindows.map((win) => {
    const app = props.apps.find((a) => a.id === win.id)
    return {
      id: win.id,
      title: win.title,
      icon: win.icon,
      isMinimized: win.mode === 'minimized',
    }
  })
})

/** 获取 Lucide 图标组件 */
function getIcon(iconName: string) {
  const icons = LucideIcons as Record<string, unknown>
  return icons[iconName] || LucideIcons.File
}
</script>

<template>
  <div
    class="
      absolute bottom-3 left-1/2 -translate-x-1/2
      flex items-end gap-1
      px-3 py-2
      rounded-2xl
      bg-white/10 backdrop-blur-2xl
      border border-white/20
      shadow-2xl
    "
  >
    <button
      v-for="item in dockItems"
      :key="item.id"
      class="
        relative flex flex-col items-center
        px-2 py-1
        rounded-lg
        transition-all duration-150 ease-out
        hover:scale-110 hover:bg-white/10
        active:scale-95
      "
      :class="{ 'opacity-40': item.isMinimized }"
      @click="emit('restore', item.id)"
      :title="item.title"
    >
      <component :is="getIcon(item.icon)" :size="28" :stroke-width="1.5" class="text-white drop-shadow-md" />

      <!-- 指示点：未最小化的窗口显示小圆点 -->
      <span
        v-if="!item.isMinimized"
        class="absolute -bottom-0.5 w-1 h-1 rounded-full bg-white/80"
      />
    </button>
  </div>
</template>
