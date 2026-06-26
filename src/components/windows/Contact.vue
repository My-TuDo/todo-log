<script setup lang="ts">
/**
 * Contact.vue - "联系我.app" 窗口内容
 * 从后端 API 获取联系方式并渲染
 */
import { ref, onMounted } from 'vue'

interface ContactItem {
  icon: string
  label: string
  value: string
}

const contacts = ref<ContactItem[]>([])

onMounted(async () => {
  const res = await fetch('http://localhost:8080/api/contact')
  contacts.value = await res.json()
})
</script>

<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold text-gray-800">📬 联系我</h2>
    <p class="text-sm text-gray-500">欢迎通过以下方式与我取得联系</p>

    <div class="space-y-3">
      <div
        v-for="(item, i) in contacts"
        :key="i"
        class="flex items-center gap-3 p-3 rounded-lg bg-gray-50 hover:bg-gray-100 transition-colors border border-gray-100"
      >
        <span class="text-xl">{{ item.icon }}</span>
        <div>
          <p class="text-xs text-gray-400">{{ item.label }}</p>
          <p class="text-sm text-gray-700">{{ item.value }}</p>
        </div>
      </div>
    </div>
  </div>
</template>
