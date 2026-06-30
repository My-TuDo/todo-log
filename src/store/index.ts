/**
 * store/index.ts - 全局响应式状态管理
 *
 * 使用 Vue 的 reactive/ref 实现轻量级全局状态，
 * 无需引入 Pinia/Vuex，所有组件 inject 后自动响应。
 */
import { reactive, readonly } from 'vue'
import type { Position } from '../types'

// ============================================================
// 桌面外观设置
// ============================================================

export interface AppSettings {
  /** 桌面壁纸 URL */
  wallpaper: string
  /** 应用自定义图标（appId -> 图片 URL，优先级高于 Lucide 图标） */
  appIcons: Record<string, string>
}

const defaultWallpaper = 'https://images.unsplash.com/photo-1506744038136-46273834b3fb?w=1920&q=80'

const settings = reactive<AppSettings>({
  wallpaper: localStorage.getItem('todo-wallpaper') || defaultWallpaper,
  appIcons: JSON.parse(localStorage.getItem('todo-appIcons') || '{}'),
})

/** 更新壁纸 */
export function setWallpaper(url: string) {
  settings.wallpaper = url
  localStorage.setItem('todo-wallpaper', url)
}

/** 设置某个应用的图标图片 */
export function setAppIcon(appId: string, imageUrl: string) {
  settings.appIcons[appId] = imageUrl
  localStorage.setItem('todo-appIcons', JSON.stringify(settings.appIcons))
}

/** 获取某个应用的自定义图标（无则返回 undefined） */
export function getAppIcon(appId: string): string | undefined {
  return settings.appIcons[appId]
}

/** 只读的 settings 引用（组件用） */
export const appSettings = readonly(settings)

// ============================================================
// 图片占位符
// ============================================================

/** 生成随机占位图 URL（用于文章配图等） */
export function placeholderImg(width = 800, height = 400, text = ''): string {
  return `https://placehold.co/${width}x${height}/e2e8f0/64748b?text=${encodeURIComponent(text || 'Image')}`
}
