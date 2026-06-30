/**
 * mock/data.ts - 本地 Mock 数据
 *
 * 开发初期使用本地 Mock 数据，无需依赖后端 API。
 * 后端就绪后只需修改 src/services/api.ts 中的开关即可切换。
 */
import type { Article, Profile, Project, ContactItem } from '../types'

/** Mock 文章列表 */
export const mockArticles: Article[] = [
  {
    id: 1,
    title: 'Vue 3 Composition API 入门指南',
    summary: '全面了解 Vue 3 的 Composition API，包括 ref、reactive、computed、watch 等核心 API 的使用方法和最佳实践。',
    created_at: '2026-06-15',
    views: 128,
  },
  {
    id: 2,
    title: 'TypeScript 高级类型体操',
    summary: '深入讲解 TypeScript 中的条件类型、映射类型、模板字面量类型等高级特性，提升你的类型编程能力。',
    created_at: '2026-06-12',
    views: 96,
  },
  {
    id: 3,
    title: 'Tailwind CSS v4 新特性解析',
    summary: 'Tailwind CSS v4 带来了全新的引擎和 API，本文带你快速上手新版本的核心变化。',
    created_at: '2026-06-08',
    views: 215,
  },
  {
    id: 4,
    title: '从零搭建 WebOS 桌面应用',
    summary: '使用 Vue 3 + TypeScript + Tailwind CSS 实现一个仿 macOS 的 Web 桌面操作系统。',
    created_at: '2026-06-01',
    views: 342,
  },
]

/** Mock 个人简介 */
export const mockProfile: Profile = {
  name: 'TUDO',
  title: '全栈开发者 / 大三学生',
  avatar_emoji: '🧑‍💻',
  bio: '热爱编程、设计和一切有趣的技术。目前正在学习企业级 Web 开发，专注于 Vue 3、Go 和云原生技术。喜欢将创意转化为实际的产品，追求代码质量与用户体验的平衡。',
  tags: ['Vue 3', 'TypeScript', 'Go', 'Tailwind CSS', 'Docker', 'Linux'],
}

/** Mock 项目列表 */
export const mockProjects: Project[] = [
  {
    id: 1,
    name: 'TODO Blog',
    description: '仿 macOS 风格的个人博客系统，前端使用 Vue 3 + Tailwind CSS，后端使用 Go + Gin。',
    tech: 'Vue 3 / Go / Tailwind CSS',
  },
  {
    id: 2,
    name: 'WebOS Desktop',
    description: '基于 Web 技术的虚拟桌面环境，支持窗口管理、多任务处理和丰富的交互体验。',
    tech: 'Vue 3 / TypeScript / CSS3',
  },
  {
    id: 3,
    name: 'API Gateway',
    description: '轻量级 API 网关，支持路由转发、负载均衡、限流和认证中间件。',
    tech: 'Go / Gin / Redis',
  },
]

/** Mock 联系信息 */
export const mockContacts: ContactItem[] = [
  { icon: '📧', label: '电子邮箱', value: 'tudo@example.com' },
  { icon: '💬', label: '微信', value: 'TUDO_dev' },
  { icon: '🐙', label: 'GitHub', value: '@tudo-dev' },
  { icon: '📱', label: '手机', value: '+86 138-0000-0000' },
]
