<script setup lang="ts">
/**
 * Projects.vue - "项目展示.app" 窗口内容
 * 通过 API 服务层获取项目列表，支持 Mock 数据
 */
import { ref, onMounted } from 'vue'
import { fetchProjects } from '../../services/api'
import type { Project } from '../../types'

const projects = ref<Project[]>([])
const loading = ref(true)
const error = ref<string | null>(null)

onMounted(async () => {
  try {
    projects.value = await fetchProjects()
  } catch (e) {
    error.value = '加载项目列表失败'
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold text-gray-800">项目展示</h2>
    <p class="text-sm text-gray-500">这里展示你的个人项目</p>

    <!-- 加载状态 -->
    <div v-if="loading" class="flex items-center justify-center py-8 text-gray-400">
      <span class="animate-pulse">加载中...</span>
    </div>

    <!-- 错误状态 -->
    <div v-else-if="error" class="p-4 rounded-lg bg-red-50 border border-red-100 text-red-500 text-sm text-center">
      {{ error }}
    </div>

    <!-- 空状态 -->
    <div v-else-if="projects.length === 0" class="py-8 text-center text-gray-400 text-sm">
      暂无项目
    </div>

    <!-- 项目列表 -->
    <div v-else class="grid gap-3">
      <div
        v-for="project in projects"
        :key="project.id"
        class="p-3 rounded-lg bg-gray-50 hover:bg-gray-100 transition-colors border border-gray-100"
      >
        <h3 class="font-medium text-gray-800">{{ project.name }}</h3>
        <p class="text-xs text-gray-500 mt-1">{{ project.description }}</p>
        <span class="inline-block mt-2 text-xs px-2 py-0.5 rounded-full bg-blue-50 text-blue-600 border border-blue-100">
          {{ project.tech }}
        </span>
      </div>
    </div>
  </div>
</template>
