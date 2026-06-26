<script setup lang="ts">
/**
 * App.vue - 根组件
 *
 * 负责提供桌面应用的配置数据（快捷方式列表、壁纸）
 * 并渲染 Desktop 组件
 */
import Desktop from './components/Desktop.vue'
import type { AppDefinition } from './types'

// ============ 应用配置 ============

/** 桌面快捷方式列表 */
const apps: AppDefinition[] = [
  {
    id: 'articles',
    title: '我的文章.app',
    icon: 'FileText',
    componentName: 'ArticleList',
  },
  {
    id: 'about',
    title: '关于我.app',
    icon: 'User',
    componentName: 'AboutMe',
  },
  {
    id: 'projects',
    title: '项目展示.app',
    icon: 'FolderGit2',
    componentName: 'Projects',
  },
  {
    id: 'contact',
    title: '联系我.app',
    icon: 'Mail',
    componentName: 'Contact',
  },
  {
    id: 'admin',
    title: '后台管理.app',
    icon: 'Terminal',
    componentName: 'AdminTerminal',
  },
]

/** 桌面壁纸 URL（使用 Unsplash 的高质量图片） */
const wallpaper = 'https://images.unsplash.com/photo-1506744038136-46273834b3fb?w=1920&q=80'
</script>

<template>
  <Desktop :apps="apps" :wallpaper="wallpaper" />
</template>
