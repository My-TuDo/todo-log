import type { Component } from 'vue'

/**
 * 桌面应用定义（快捷方式元数据）
 * 用于配置桌面上的图标和对应的窗口内容
 */
export interface AppDefinition {
  /** 唯一标识 */
  id: string
  /** 显示名称（例如："我的文章.app"） */
  title: string
  /** Lucide 图标名称（例如："FileText"） */
  icon: string
  /** 对应的窗口组件名称，用于动态组件查找 */
  componentName: string
}

/** 窗口在桌面上的位置 */
export interface Position {
  x: number
  y: number
}

/** 窗口尺寸 */
export interface Size {
  width: number
  height: number
}

/**
 * 窗口运行时状态
 * 每个打开的窗口对应一个 WindowState 实例
 */
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
  /** 是否正在播放关闭动画 */
  isClosing: boolean
}

/** 动态组件映射表：componentName -> Vue 组件 */
export type ComponentMap = Record<string, Component>
