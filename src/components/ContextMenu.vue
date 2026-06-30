<script setup lang="ts">
/**
 * ContextMenu.vue - 桌面右键菜单
 *
 * 功能：
 * - 在右键点击位置弹出
 * - 支持"新建文章"等快捷操作
 * - 点击菜单项或点击外部关闭
 * - 自动吸附屏幕边缘，避免溢出
 */
import { computed, onMounted, onUnmounted } from 'vue'

const props = defineProps<{
  x: number
  y: number
}>()

const emit = defineEmits<{
  close: []
  openApp: [appId: string]
}>()

const MENU_WIDTH = 180
const MENU_HEIGHT = 160

/** 计算菜单位置，防止溢出屏幕 */
const menuStyle = computed(() => {
  const maxX = window.innerWidth - MENU_WIDTH - 8
  const maxY = window.innerHeight - MENU_HEIGHT - 8
  return {
    left: Math.min(props.x, maxX) + 'px',
    top: Math.min(props.y, maxY) + 'px',
  }
})

function handleClickOutside() {
  emit('close')
}

function handleMenuItem(action: string) {
  switch (action) {
    case 'new-article':
      emit('openApp', 'admin')
      break
    case 'refresh':
      window.location.reload()
      break
  }
  emit('close')
}

onMounted(() => {
  setTimeout(() => {
    document.addEventListener('click', handleClickOutside)
  }, 0)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<template>
  <div
    class="fixed z-[9999]"
    :style="menuStyle"
    @click.stop
  >
    <div
      class="
        min-w-[180px] py-1
        rounded-xl overflow-hidden
        bg-white/90 backdrop-blur-2xl
        border border-gray-200/60
        shadow-2xl
        text-sm text-gray-700
      "
    >
      <button
        class="w-full flex items-center gap-3 px-4 py-2 hover:bg-blue-500 hover:text-white transition-colors text-left"
        @click="handleMenuItem('new-article')"
      >
        <span>📝</span>
        <span>新建文章</span>
      </button>
      <button
        class="w-full flex items-center gap-3 px-4 py-2 hover:bg-blue-500 hover:text-white transition-colors text-left"
        @click="handleMenuItem('refresh')"
      >
        <span>🔄</span>
        <span>刷新桌面</span>
      </button>
      <div class="border-t border-gray-200/60 my-1" />
      <div class="px-4 py-1.5 text-xs text-gray-400">
        TUDO's blog v0.1
      </div>
    </div>
  </div>
</template>
