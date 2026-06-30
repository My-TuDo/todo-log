import type { Component } from 'vue'

// ============================================================
// 桌面应用定义
// ============================================================

/** 桌面应用定义（快捷方式元数据） */
export interface AppDefinition {
  /** 唯一标识 */
  id: string
  /** 显示名称（例如："我的文章.app"） */
  title: string
  /** Lucide 图标名称（例如："FileText"） */
  icon: string
  /** 对应的窗口组件名称，用于动态组件查找 */
  componentName: string
  /** 自定义图标图片 URL（优先级高于 Lucide 图标） */
  imageUrl?: string
}

// ============================================================
// 窗口系统类型
// ============================================================

export interface Position {
  x: number
  y: number
}

export interface Size {
  width: number
  height: number
}

/** 窗口显示模式 */
export type WindowMode = 'normal' | 'maximized' | 'minimized'

/** 窗口运行时状态 */
export interface WindowState {
  /** 窗口唯一标识 */
  id: string
  /** 窗口标题 */
  title: string
  /** 窗口图标 */
  icon: string
  /** 对应的动态组件名称 */
  componentName: string
  /** 层级（越大越靠前） */
  zIndex: number
  /** 窗口位置（相对于桌面左上角） */
  position: Position
  /** 窗口尺寸 */
  size: Size
  /** 显示模式 */
  mode: WindowMode
  /** 是否正在播放关闭动画 */
  isClosing: boolean
  /** 最大化前保存的位置（用于还原） */
  previousPosition?: Position
  /** 最大化前保存的尺寸（用于还原） */
  previousSize?: Size
}

/** 动态组件映射表：componentName -> Vue 组件 */
export type ComponentMap = Record<string, Component>

// ============================================================
// 业务数据类型（统一集中管理）
// ============================================================

/** 文章 */
export interface Article {
  id: number
  title: string
  summary: string
  created_at: string
  views: number
  /** 文章配图 URL */
  image_url?: string
}

/** 个人简介 */
export interface Profile {
  name: string
  title: string
  avatar_emoji: string
  /** 头像图片 URL（优先级高于 avatar_emoji） */
  avatar_url?: string
  bio: string
  tags: string[]
}

/** 项目 */
export interface Project {
  id: number
  name: string
  description: string
  tech: string
}

/** 联系方式 */
export interface ContactItem {
  icon: string
  label: string
  value: string
}

/** 登录响应 */
export interface LoginResponse {
  token: string
}

/** API 通用错误响应 */
export interface ApiError {
  error: string
  message?: string
}
