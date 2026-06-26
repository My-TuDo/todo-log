<script setup lang="ts">
/**
 * ArticleList.vue - "我的文章.app" 窗口内容
 * 演示用：展示文章列表占位
 */
</script>

<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold text-gray-800">📝 我的文章</h2>
    <p class="text-sm text-gray-500">这里将展示你的博客文章列表</p>

    <!-- 文章列表占位 -->
    <div class="space-y-3">
      <div
        v-for="i in 5"
        :key="i"
        class="p-3 rounded-lg bg-gray-50 hover:bg-gray-100 transition-colors cursor-pointer border border-gray-100"
      >
        <h3 class="font-medium text-gray-800">示例文章标题 {{ i }}</h3>
        <p class="text-xs text-gray-500 mt-1">
          这是文章的摘要描述文字，后续将从后端 API 获取真实数据...
        </p>
        <div class="flex items-center gap-3 mt-2 text-xs text-gray-400">
          <span>📅 2026-06-{{ String(i).padStart(2, '0') }}</span>
          <span>👁️ {{ i * 10 }} 次阅读</span>
        </div>
      </div>
    </div>
  </div>
</template>
