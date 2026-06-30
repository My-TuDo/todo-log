<script setup lang="ts">
/**
 * App.vue - 根组件
 *
 * 提供桌面应用配置，壁纸由全局 store 管理
 */
import Desktop from './components/Desktop.vue'
import { appSettings } from './store'
import type { AppDefinition } from './types'

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
</script>

<template>
  <Desktop :apps="apps" :wallpaper="appSettings.wallpaper" />
</template>
