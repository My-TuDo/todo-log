<script setup lang="ts">
/**
 * ArticleList.vue - "我的文章.app" 窗口内容
 * 通过 API 服务层获取文章列表，支持 Mock 数据
 */
import { ref, onMounted } from 'vue'
import { fetchArticles } from '../../services/api'
import type { Article } from '../../types'

const articles = ref<Article[]>([])
const loading = ref(true)
const error = ref<string | null>(null)

onMounted(async () => {
  try {
    articles.value = await fetchArticles()
  } catch (e) {
    error.value = '加载文章失败，请稍后重试'
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold text-gray-800">📝 我的文章</h2>
    <p class="text-sm text-gray-500">这里展示你的博客文章列表</p>

    <!-- 加载状态 -->
    <div v-if="loading" class="flex items-center justify-center py-8 text-gray-400">
      <span class="animate-pulse">加载中...</span>
    </div>

    <!-- 错误状态 -->
    <div v-else-if="error" class="p-4 rounded-lg bg-red-50 border border-red-100 text-red-500 text-sm text-center">
      {{ error }}
    </div>

    <!-- 空状态 -->
    <div v-else-if="articles.length === 0" class="py-8 text-center text-gray-400 text-sm">
      暂无文章
    </div>

    <!-- 文章列表 -->
    <div v-else class="space-y-3">
      <div
        v-for="article in articles"
        :key="article.id"
        class="p-3 rounded-lg bg-gray-50 hover:bg-gray-100 transition-colors cursor-pointer border border-gray-100"
      >
        <h3 class="font-medium text-gray-800">{{ article.title }}</h3>
        <p class="text-xs text-gray-500 mt-1">{{ article.summary }}</p>
        <div class="flex items-center gap-3 mt-2 text-xs text-gray-400">
          <span>📅 {{ article.created_at }}</span>
          <span>👁️ {{ article.views }} 次阅读</span>
        </div>
      </div>
    </div>
  </div>
</template>
