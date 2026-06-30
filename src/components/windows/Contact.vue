<script setup lang="ts">
/**
 * Contact.vue - "联系我.app" 窗口内容
 * 通过 API 服务层获取联系方式，支持 Mock 数据
 */
import { ref, onMounted } from 'vue'
import { fetchContacts } from '../../services/api'
import type { ContactItem } from '../../types'

const contacts = ref<ContactItem[]>([])
const loading = ref(true)
const error = ref<string | null>(null)

onMounted(async () => {
  try {
    contacts.value = await fetchContacts()
  } catch (e) {
    error.value = '加载联系方式失败'
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold text-gray-800">联系我</h2>
    <p class="text-sm text-gray-500">欢迎通过以下方式与我取得联系</p>

    <!-- 加载状态 -->
    <div v-if="loading" class="flex items-center justify-center py-8 text-gray-400">
      <span class="animate-pulse">加载中...</span>
    </div>

    <!-- 错误状态 -->
    <div v-else-if="error" class="p-4 rounded-lg bg-red-50 border border-red-100 text-red-500 text-sm text-center">
      {{ error }}
    </div>

    <!-- 空状态 -->
    <div v-else-if="contacts.length === 0" class="py-8 text-center text-gray-400 text-sm">
      暂无联系方式
    </div>

    <!-- 联系列表 -->
    <div v-else class="space-y-3">
      <div
        v-for="(item, i) in contacts"
        :key="i"
        class="flex items-center gap-3 p-3 rounded-lg bg-gray-50 hover:bg-gray-100 transition-colors border border-gray-100"
      >
        <span class="text-xs text-gray-300 font-mono w-6 text-center shrink-0">{{ item.icon }}</span>
        <div>
          <p class="text-xs text-gray-400">{{ item.label }}</p>
          <p class="text-sm text-gray-700">{{ item.value }}</p>
        </div>
      </div>
    </div>
  </div>
</template>
