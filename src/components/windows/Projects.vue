<script setup lang="ts">
/**
 * Projects.vue - "项目展示.app" 窗口内容
 * 从后端 API 获取项目列表并渲染
 */
import { ref, onMounted } from 'vue'

interface Project {
  id: number
  name: string
  description: string
  tech: string
}

const projects = ref<Project[]>([])

onMounted(async () => {
  const res = await fetch('http://localhost:8080/api/projects')
  projects.value = await res.json()
})
</script>

<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold text-gray-800">💻 项目展示</h2>
    <p class="text-sm text-gray-500">这里将展示你的个人项目</p>

    <div class="grid gap-3">
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
