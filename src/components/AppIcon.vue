<script setup lang="ts">
/**
 * AppIcon.vue - 桌面图标组件
 *
 * 功能：
 * - 显示应用图标（Lucide）+ 文字标签
 * - 悬浮时放大 + 阴影效果
 * - 单击触发打开窗口逻辑
 */
import type { AppDefinition } from '../types'
import * as LucideIcons from 'lucide-vue-next'
import { computed } from 'vue'
import { h } from 'vue'

const props = defineProps<{
  app: AppDefinition
  onOpen: () => void
}>()

// 根据 icon 名称动态获取 Lucide 图标组件
const IconComponent = computed(() => {
  const iconName = props.app.icon as keyof typeof LucideIcons
  return (LucideIcons[iconName] as ReturnType<typeof h>) || LucideIcons.File
})
</script>

<template>
  <div
    class="
      flex flex-col items-center justify-center
      w-24 h-28
      cursor-pointer
      rounded-xl
      transition-all duration-200 ease-out
      hover:scale-105
      hover:bg-white/10
      hover:shadow-lg hover:shadow-white/10
      active:scale-95
      p-2
    "
    @click="onOpen"
  >
    <!-- 图标 -->
    <div class="w-14 h-14 flex items-center justify-center text-white drop-shadow-lg">
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
      "
    >
      {{ app.title }}
    </span>
  </div>
</template>
