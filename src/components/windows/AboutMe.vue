<script setup lang="ts">
/**
 * AboutMe.vue - "关于我.app" 窗口内容
 * 从后端 API 获取个人简介并渲染
 */
import { ref, onMounted } from 'vue'

interface Profile {
  name: string
  title: string
  avatar_emoji: string
  bio: string
  tags: string[]
}

const profile = ref<Profile | null>(null)

onMounted(async () => {
  const res = await fetch('http://localhost:8080/api/profile')
  profile.value = await res.json()
})
</script>

<template>
  <div class="space-y-4" v-if="profile">
    <h2 class="text-xl font-bold text-gray-800">👤 关于我</h2>

    <div class="flex flex-col items-center py-4 space-y-3">
      <!-- 头像占位 -->
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
      <p class="text-sm text-gray-700 leading-relaxed">
        {{ profile.bio }}
      </p>
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
</template>
