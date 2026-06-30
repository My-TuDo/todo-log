<script setup lang="ts">
/**
 * AboutMe.vue - "关于我.app" 窗口内容
 * 通过 API 服务层获取个人简介，支持 Mock 数据
 */
import { ref, onMounted } from 'vue'
import { fetchProfile } from '../../services/api'
import type { Profile } from '../../types'

const profile = ref<Profile | null>(null)
const loading = ref(true)
const error = ref<string | null>(null)

onMounted(async () => {
  try {
    profile.value = await fetchProfile()
  } catch (e) {
    error.value = '加载个人资料失败'
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold text-gray-800">👤 关于我</h2>

    <!-- 加载状态 -->
    <div v-if="loading" class="flex items-center justify-center py-8 text-gray-400">
      <span class="animate-pulse">加载中...</span>
    </div>

    <!-- 错误状态 -->
    <div v-else-if="error" class="p-4 rounded-lg bg-red-50 border border-red-100 text-red-500 text-sm text-center">
      {{ error }}
    </div>

    <!-- 内容 -->
    <div v-else-if="profile" class="space-y-4">
      <div class="flex flex-col items-center py-4 space-y-3">
        <!-- 头像 -->
        <div
          class="
            w-20 h-20 rounded-full
            bg-gradient-to-br from-blue-400 to-purple-500
            flex items-center justify-center
            text-3xl text-white
            shadow-lg
          "
        >
          {{ profile.avatar_emoji }}
        </div>
        <div class="text-center">
          <h3 class="text-lg font-semibold text-gray-800">{{ profile.name }}</h3>
          <p class="text-sm text-gray-500">{{ profile.title }}</p>
        </div>
      </div>

      <div class="p-3 rounded-lg bg-gray-50 border border-gray-100">
        <p class="text-sm text-gray-700 leading-relaxed">{{ profile.bio }}</p>
      </div>

      <div class="flex flex-wrap gap-2">
        <span
          v-for="tag in profile.tags"
          :key="tag"
          class="px-2 py-1 text-xs rounded-full bg-gray-100 text-gray-600"
        >
          {{ tag }}
        </span>
      </div>
    </div>
  </div>
</template>
