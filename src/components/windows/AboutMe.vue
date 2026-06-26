<script setup lang="ts">
/**
 * AboutMe.vue - "关于我.app" 窗口内容
 * 演示用：展示个人简介占位
 */
</script>

<template>
  <div class="space-y-4">
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
        🐱
      </div>
      <div class="text-center">
        <h3 class="text-lg font-semibold text-gray-800">TUDO</h3>
        <p class="text-sm text-gray-500">全栈开发者 / 学生</p>
      </div>
    </div>

    <div class="p-3 rounded-lg bg-gray-50 border border-gray-100">
      <p class="text-sm text-gray-700 leading-relaxed">
        热爱编程、设计和一切有趣的事物。这个博客将记录我的学习笔记、项目经验和生活思考。🚀
      </p>
    </div>

    <div class="flex flex-wrap gap-2">
      <span
        v-for="tag in ['Vue', 'Go', 'TypeScript', 'Tailwind CSS', 'Docker']"
        :key="tag"
        class="px-2 py-1 text-xs rounded-full bg-gray-100 text-gray-600"
      >
        {{ tag }}
      </span>
    </div>
  </div>
</template>
