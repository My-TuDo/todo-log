/**
 * services/api.ts - 统一 API 服务层
 *
 * 通过 USE_MOCK 开关控制使用 Mock 数据还是后端 API。
 * - true:  使用本地 Mock 数据（开发初期，无需后端）
 * - false: 调用后端 API（后端就绪后切换）
 *
 * 所有业务组件统一调用此文件中的函数，不再直接 fetch。
 */
import type { Article, Profile, Project, ContactItem, LoginResponse } from '../types'
import { mockArticles, mockProfile, mockProjects, mockContacts } from '../mock/data'

// ============================================================
// 开关：是否使用 Mock 数据
// ============================================================
const USE_MOCK = true

// ============================================================
// 基础 fetch 封装
// ============================================================

const BASE_URL = 'http://localhost:8080'

interface FetchOptions {
  method?: string
  body?: unknown
  token?: string
}

async function apiFetch<T>(path: string, options: FetchOptions = {}): Promise<T> {
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
  }
  if (options.token) {
    headers['Authorization'] = `Bearer ${options.token}`
  }

  const res = await fetch(`${BASE_URL}${path}`, {
    method: options.method || 'GET',
    headers,
    body: options.body ? JSON.stringify(options.body) : undefined,
  })

  if (!res.ok) {
    throw new Error(`API Error ${res.status}: ${res.statusText}`)
  }

  return res.json()
}

// ============================================================
// 模拟网络延迟（Mock 模式下使用）
// ============================================================

function delay(ms = 300): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

// ============================================================
// 文章 API
// ============================================================

export async function fetchArticles(): Promise<Article[]> {
  if (USE_MOCK) {
    await delay()
    return [...mockArticles]
  }
  return apiFetch<Article[]>('/api/articles')
}

export async function createArticle(
  title: string,
  summary: string,
  token: string,
): Promise<void> {
  if (USE_MOCK) {
    await delay()
    mockArticles.unshift({
      id: Date.now(),
      title,
      summary,
      created_at: new Date().toISOString().slice(0, 10),
      views: 0,
    })
    return
  }
  await apiFetch('/api/articles', {
    method: 'POST',
    body: { title, summary },
    token,
  })
}

export async function deleteArticle(id: number, token: string): Promise<void> {
  if (USE_MOCK) {
    await delay()
    const idx = mockArticles.findIndex((a) => a.id === id)
    if (idx !== -1) mockArticles.splice(idx, 1)
    return
  }
  await apiFetch(`/api/articles/${id}`, { method: 'DELETE', token })
}

// ============================================================
// 个人简介 API
// ============================================================

export async function fetchProfile(): Promise<Profile> {
  if (USE_MOCK) {
    await delay()
    return { ...mockProfile }
  }
  return apiFetch<Profile>('/api/profile')
}

export async function updateProfile(data: Partial<Profile>, token: string): Promise<void> {
  if (USE_MOCK) {
    await delay()
    Object.assign(mockProfile, data)
    return
  }
  await apiFetch('/api/profile', { method: 'PUT', body: data, token })
}

// ============================================================
// 项目 API
// ============================================================

export async function fetchProjects(): Promise<Project[]> {
  if (USE_MOCK) {
    await delay()
    return [...mockProjects]
  }
  return apiFetch<Project[]>('/api/projects')
}

export async function createProject(
  name: string,
  description: string,
  tech: string,
  token: string,
): Promise<void> {
  if (USE_MOCK) {
    await delay()
    mockProjects.push({ id: Date.now(), name, description, tech })
    return
  }
  await apiFetch('/api/projects', { method: 'POST', body: { name, description, tech }, token })
}

export async function deleteProject(id: number, token: string): Promise<void> {
  if (USE_MOCK) {
    await delay()
    const idx = mockProjects.findIndex((p) => p.id === id)
    if (idx !== -1) mockProjects.splice(idx, 1)
    return
  }
  await apiFetch(`/api/projects/${id}`, { method: 'DELETE', token })
}

// ============================================================
// 联系方式 API
// ============================================================

export async function fetchContacts(): Promise<ContactItem[]> {
  if (USE_MOCK) {
    await delay()
    return [...mockContacts]
  }
  return apiFetch<ContactItem[]>('/api/contact')
}

export async function updateContacts(contacts: ContactItem[], token: string): Promise<void> {
  if (USE_MOCK) {
    await delay()
    mockContacts.length = 0
    mockContacts.push(...contacts)
    return
  }
  await apiFetch('/api/contact', { method: 'PUT', body: { contacts }, token })
}

// ============================================================
// 登录 API
// ============================================================

export async function login(password: string): Promise<LoginResponse> {
  if (USE_MOCK) {
    await delay()
    if (password === 'admin123') {
      return { token: 'mock-token-' + Date.now() }
    }
    throw new Error('密码错误')
  }
  return apiFetch<LoginResponse>('/api/login', {
    method: 'POST',
    body: { password },
  })
}
