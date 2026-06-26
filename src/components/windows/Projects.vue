<script setup lang="ts">
/**
 * Projects.vue - "项目展示.app" 窗口内容
 * 演示用：展示项目列表占位
 */
</script>

<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold text-gray-800">💻 项目展示</h2>
    <p class="text-sm text-gray-500">这里将展示你的个人项目</p>

    <div class="grid gap-3">
      <div
        v-for="(project, i) in [
          { name: 'TODO Blog', desc: '仿 macOS 桌面的个人博客系统', tech: 'Vue 3 + Go' },
          { name: 'Task Manager', desc: '简洁高效的待办事项管理工具', tech: 'React + Node.js' },
          { name: 'Weather App', desc: '实时天气查询与预报应用', tech: 'Vue 3 + API' },
        ]"
        :key="i"
        class="p-3 rounded-lg bg-gray-50 hover:bg-gray-100 transition-colors border border-gray-100"
      >
        <h3 class="font-medium text-gray-800">{{ project.name }}</h3>
        <p class="text-xs text-gray-500 mt-1">{{ project.desc }}</p>
        <span class="inline-block mt-2 text-xs px-2 py-0.5 rounded-full bg-blue-50 text-blue-600 border border-blue-100">
          {{ project.tech }}
        </span>
      </div>
    </div>
  </div>
</template>
