<script setup lang="ts">
/**
 * ContextMenu.vue - 桌面右键菜单
 *
 * 功能：
 * - 在右键点击位置弹出
 * - 支持"新建文章"等快捷操作
 * - 点击菜单项或点击外部关闭
 */
import { onMounted, onUnmounted } from 'vue'

const props = defineProps<{
  x: number
  y: number
}>()

const emit = defineEmits<{
  close: []
}>()

function handleClickOutside() {
  emit('close')
}

function handleMenuItem(action: string) {
  switch (action) {
    case 'new-article':
      window.open('http://localhost:5173', '_blank')
      break
    case 'refresh':
      window.location.reload()
      break
  }
  emit('close')
}

onMounted(() => {
  // 点击菜单外部关闭
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
    :style="{ left: x + 'px', top: y + 'px' }"
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
