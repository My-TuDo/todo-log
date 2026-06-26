<script setup lang="ts">
/**
 * ArticleList.vue - "我的文章.app" 窗口内容
 * 从后端 API 获取文章列表并渲染
 */
import { ref, onMounted } from 'vue'

// 定义文章数据的类型
interface Article {
    id: number  // 文章唯一标识
    title: string // 文章标题
    summary: string // 文章摘要
    created_at: string  // 文章创建时间
    views: number   // 文章阅读量
}

// 响应式数据
const articles = ref<Article[]>([])

// 页面加载时自动调用后端API
onMounted(async () =>{
    const response = await fetch('http://localhost:8080/api/articles')
    articles.value = await response.json()
})
</script>

<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold text-gray-800">📝 我的文章</h2>
    <p class="text-sm text-gray-500">这里将展示你的博客文章列表</p>

    <!-- 文章列表（从后端 API 获取） -->
    <div class="space-y-3">
      <div
        v-for="article in articles"
        :key="article.id"
        class="p-3 rounded-lg bg-gray-50 hover:bg-gray-100 transition-colors cursor-pointer border border-gray-100"
      >
        <h3 class="font-medium text-gray-800">{{ article.title }}</h3>
        <p class="text-xs text-gray-500 mt-1">
          {{ article.summary }}
        </p>
        <div class="flex items-center gap-3 mt-2 text-xs text-gray-400">
          <span>📅 {{ article.created_at }}</span>
          <span>👁️ {{ article.views }} 次阅读</span>
        </div>
      </div>
    </div>
  </div>
</template>
