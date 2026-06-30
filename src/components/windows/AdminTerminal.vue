<script setup lang="ts">
/**
 * AdminTerminal.vue - "后台管理.app" 窗口内容
 *
 * 简洁风格图形化后台管理界面
 * - 文章管理：列表查看、创建、删除
 * - 项目管理：列表查看、创建、删除
 * - 个人资料：查看与编辑
 * - 联系方式：查看与编辑
 */
import { ref, onMounted, computed } from 'vue'
import {
  login as apiLogin,
  fetchArticles,
  createArticle,
  deleteArticle,
  fetchProjects,
  createProject,
  deleteProject,
  fetchProfile,
  updateProfile,
  fetchContacts,
  updateContacts,
} from '../../services/api'
import { appSettings, setWallpaper, setAppIcon } from '../../store'
import type { Article, Project, Profile, ContactItem, AppDefinition } from '../../types'

/** 壁纸预设 */
const wallpapers = [
  { url: 'https://images.unsplash.com/photo-1506744038136-46273834b3fb?w=1920&q=80', name: '山水' },
  { url: 'https://images.unsplash.com/photo-1519681393784-d120267933ba?w=1920&q=80', name: '星空' },
  { url: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=1920&q=80', name: '极光' },
  { url: 'https://images.unsplash.com/photo-1470071459604-7b8ec44ffd4b?w=1920&q=80', name: '森林' },
  { url: 'https://images.unsplash.com/photo-1447752875215-b2761acb3c5d?w=1920&q=80', name: '湖畔' },
  { url: 'https://images.unsplash.com/photo-1469474968028-56623f02e42e?w=1920&q=80', name: '山川' },
]

/** 快捷方式列表（用于外观设置中的图标配置） */
const apps: AppDefinition[] = [
  { id: 'articles', title: '我的文章.app', icon: 'FileText', componentName: 'ArticleList' },
  { id: 'about', title: '关于我.app', icon: 'User', componentName: 'AboutMe' },
  { id: 'projects', title: '项目展示.app', icon: 'FolderGit2', componentName: 'Projects' },
  { id: 'contact', title: '联系我.app', icon: 'Mail', componentName: 'Contact' },
  { id: 'admin', title: '后台管理.app', icon: 'Terminal', componentName: 'AdminTerminal' },
]

// ============================================================
// 认证
// ============================================================

const isLoggedIn = ref(false)
const token = ref('')
const loginPassword = ref('')
const loginError = ref('')
const loginLoading = ref(false)

async function handleLogin() {
  if (!loginPassword.value) return
  loginLoading.value = true
  loginError.value = ''
  try {
    const res = await apiLogin(loginPassword.value)
    token.value = res.token
    isLoggedIn.value = true
    loginPassword.value = ''
    loadAllData()
  } catch (e: unknown) {
    loginError.value = e instanceof Error ? e.message : '登录失败'
  } finally {
    loginLoading.value = false
  }
}

function handleLogout() {
  isLoggedIn.value = false
  token.value = ''
}

// ============================================================
// 导航
// ============================================================

type Tab = 'articles' | 'projects' | 'profile' | 'contacts' | 'appearance'
const activeTab = ref<Tab>('articles')

const tabs: { id: Tab; label: string }[] = [
  { id: 'articles', label: '文章管理' },
  { id: 'projects', label: '项目管理' },
  { id: 'profile', label: '个人资料' },
  { id: 'contacts', label: '联系方式' },
  { id: 'appearance', label: '外观设置' },
]

// ============================================================
// 数据
// ============================================================

const articles = ref<Article[]>([])
const projects = ref<Project[]>([])
const profile = ref<Profile | null>(null)
const contacts = ref<ContactItem[]>([])

const loading = ref(false)
const toastMessage = ref('')
const toastType = ref<'success' | 'error'>('success')
let toastTimer: ReturnType<typeof setTimeout> | null = null

function showToast(msg: string, type: 'success' | 'error' = 'success') {
  toastMessage.value = msg
  toastType.value = type
  if (toastTimer) clearTimeout(toastTimer)
  toastTimer = setTimeout(() => { toastMessage.value = '' }, 2500)
}

async function loadAllData() {
  loading.value = true
  try {
    const [a, p, pr, c] = await Promise.all([
      fetchArticles(),
      fetchProjects(),
      fetchProfile(),
      fetchContacts(),
    ])
    articles.value = a
    projects.value = p
    profile.value = pr
    contacts.value = c
  } catch {
    showToast('数据加载失败', 'error')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  if (isLoggedIn.value) loadAllData()
})

// ============================================================
// 文章管理
// ============================================================

const showAddArticle = ref(false)
const newArticle = ref({ title: '', summary: '' })
const articleLoading = ref(false)

async function handleCreateArticle() {
  if (!newArticle.value.title || !newArticle.value.summary) return
  articleLoading.value = true
  try {
    await createArticle(newArticle.value.title, newArticle.value.summary, token.value)
    showToast('文章创建成功')
    newArticle.value = { title: '', summary: '' }
    showAddArticle.value = false
    articles.value = await fetchArticles()
  } catch {
    showToast('创建失败', 'error')
  } finally {
    articleLoading.value = false
  }
}

async function handleDeleteArticle(id: number) {
  if (!confirm('确定要删除这篇文章吗？')) return
  try {
    await deleteArticle(id, token.value)
    showToast('文章已删除')
    articles.value = await fetchArticles()
  } catch {
    showToast('删除失败', 'error')
  }
}

// ============================================================
// 项目管理
// ============================================================

const showAddProject = ref(false)
const newProject = ref({ name: '', description: '', tech: '' })
const projectLoading = ref(false)

async function handleCreateProject() {
  if (!newProject.value.name) return
  projectLoading.value = true
  try {
    await createProject(newProject.value.name, newProject.value.description, newProject.value.tech, token.value)
    showToast('项目创建成功')
    newProject.value = { name: '', description: '', tech: '' }
    showAddProject.value = false
    projects.value = await fetchProjects()
  } catch {
    showToast('创建失败', 'error')
  } finally {
    projectLoading.value = false
  }
}

async function handleDeleteProject(id: number) {
  if (!confirm('确定要删除这个项目吗？')) return
  try {
    await deleteProject(id, token.value)
    showToast('项目已删除')
    projects.value = await fetchProjects()
  } catch {
    showToast('删除失败', 'error')
  }
}

// ============================================================
// 个人资料
// ============================================================

const profileForm = ref<Profile>({
  name: '', title: '', avatar_emoji: '', bio: '', tags: [],
})
const profileLoading = ref(false)

function initProfileForm() {
  if (profile.value) {
    profileForm.value = { ...profile.value, tags: [...profile.value.tags] }
  }
}

async function handleSaveProfile() {
  profileLoading.value = true
  try {
    await updateProfile(profileForm.value, token.value)
    showToast('个人资料已更新')
    profile.value = { ...profileForm.value }
  } catch {
    showToast('保存失败', 'error')
  } finally {
    profileLoading.value = false
  }
}

/** 取姓名首字作为头像占位 */
const avatarInitial = computed(() => {
  return profileForm.value.name ? profileForm.value.name.charAt(0).toUpperCase() : '?'
})

// ============================================================
// 联系方式
// ============================================================

const contactsEdit = ref<ContactItem[]>([])
const contactsLoading = ref(false)

function initContactsForm() {
  contactsEdit.value = contacts.value.map(c => ({ ...c }))
}

async function handleSaveContacts() {
  contactsLoading.value = true
  try {
    await updateContacts(contactsEdit.value, token.value)
    showToast('联系方式已更新')
    contacts.value = contactsEdit.value.map(c => ({ ...c }))
  } catch {
    showToast('保存失败', 'error')
  } finally {
    contactsLoading.value = false
  }
}

// Tab 切换时初始化表单
const activeTabComputed = computed({
  get: () => activeTab.value,
  set: (val: Tab) => {
    activeTab.value = val
    if (val === 'profile') initProfileForm()
    if (val === 'contacts') initContactsForm()
  },
})
</script>

<template>
  <!-- ===================== 登录页 ===================== -->
  <div v-if="!isLoggedIn" class="h-full w-full flex items-center justify-center bg-gray-50">
    <div class="w-80 p-8 rounded-xl bg-white shadow-sm border border-gray-100">
      <div class="text-center mb-6">
        <div class="w-12 h-12 mx-auto mb-3 rounded-full bg-gray-100 flex items-center justify-center text-gray-400 text-lg font-mono">/</div>
        <h2 class="text-base font-semibold text-gray-800">后台管理</h2>
        <p class="text-xs text-gray-400 mt-1">请输入密码登录</p>
      </div>
      <form @submit.prevent="handleLogin" class="space-y-4">
        <div>
          <input
            v-model="loginPassword"
            type="password"
            placeholder="管理员密码"
            class="w-full px-3 py-2 rounded-lg border border-gray-200 bg-white text-sm text-gray-700 placeholder-gray-400 outline-none focus:ring-1 focus:ring-blue-400 focus:border-blue-400 transition-all"
            :class="{ 'border-red-300': loginError }"
          />
          <p v-if="loginError" class="text-xs text-red-500 mt-1">{{ loginError }}</p>
        </div>
        <button
          type="submit"
          class="w-full py-2 rounded-lg bg-gray-900 hover:bg-gray-800 text-white text-sm font-medium transition-all disabled:opacity-50"
          :disabled="loginLoading"
        >
          {{ loginLoading ? '登录中...' : '登录' }}
        </button>
      </form>
      <p class="text-center text-[10px] text-gray-300 mt-4">默认密码: admin123</p>
    </div>
  </div>

  <!-- ===================== 主界面 ===================== -->
  <div v-else class="h-full w-full flex flex-col bg-white">
    <!-- 顶栏 -->
    <div class="flex items-center justify-between px-4 py-2.5 border-b border-gray-100 shrink-0">
      <div class="flex items-center gap-2">
        <span class="text-xs font-semibold text-gray-800 tracking-wide">管理后台</span>
        <span class="text-[10px] px-1.5 py-0.5 rounded bg-gray-100 text-gray-500">admin</span>
      </div>
      <button class="text-xs text-gray-400 hover:text-gray-600 transition-colors" @click="handleLogout">退出</button>
    </div>

    <!-- 主体 -->
    <div class="flex-1 flex overflow-hidden">
      <!-- 侧边栏 -->
      <div class="w-36 shrink-0 border-r border-gray-100 p-1.5 space-y-0.5">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          class="w-full px-3 py-2 rounded-md text-xs text-left transition-all"
          :class="activeTab === tab.id
            ? 'bg-gray-100 text-gray-900 font-medium'
            : 'text-gray-500 hover:text-gray-700 hover:bg-gray-50'"
          @click="activeTabComputed = tab.id"
        >
          {{ tab.label }}
        </button>
      </div>

      <!-- 内容 -->
      <div class="flex-1 overflow-y-auto p-5 bg-gray-50/50">
        <div v-if="loading" class="flex items-center justify-center h-full text-xs text-gray-400">
          <span class="animate-pulse">加载中...</span>
        </div>

        <!-- ===== 文章管理 ===== -->
        <div v-else-if="activeTab === 'articles'" class="max-w-2xl space-y-3">
          <div class="flex items-center justify-between">
            <h2 class="text-sm font-semibold text-gray-800">文章管理</h2>
            <button
              v-if="!showAddArticle"
              class="px-3 py-1.5 rounded-md bg-gray-900 hover:bg-gray-800 text-white text-xs transition-all"
              @click="showAddArticle = true"
            >+ 新建</button>
          </div>

          <div v-if="showAddArticle" class="p-4 rounded-lg bg-white border border-gray-200 space-y-3">
            <input v-model="newArticle.title" placeholder="文章标题" class="w-full px-3 py-1.5 rounded border border-gray-200 bg-gray-50 text-xs outline-none focus:border-blue-400 transition-colors" />
            <textarea v-model="newArticle.summary" placeholder="文章摘要" rows="2" class="w-full px-3 py-1.5 rounded border border-gray-200 bg-gray-50 text-xs outline-none focus:border-blue-400 transition-colors resize-none"></textarea>
            <div class="flex justify-end gap-2">
              <button class="px-3 py-1 rounded text-xs text-gray-500 hover:bg-gray-100 transition-colors" @click="showAddArticle = false">取消</button>
              <button class="px-3 py-1 rounded text-xs bg-gray-900 hover:bg-gray-800 text-white transition-all disabled:opacity-50" :disabled="articleLoading || !newArticle.title" @click="handleCreateArticle">{{ articleLoading ? '创建中...' : '创建' }}</button>
            </div>
          </div>

          <div v-if="articles.length === 0" class="py-10 text-center text-xs text-gray-300">暂无文章</div>
          <div v-else class="space-y-1.5">
            <div v-for="article in articles" :key="article.id" class="flex items-start gap-2 p-3 rounded-lg bg-white border border-gray-100 hover:border-gray-200 transition-all group">
              <div class="flex-1 min-w-0">
                <div class="text-sm font-medium text-gray-800 truncate">{{ article.title }}</div>
                <div class="text-xs text-gray-400 mt-0.5 line-clamp-1">{{ article.summary }}</div>
                <div class="text-[10px] text-gray-300 mt-1">{{ article.created_at }} · {{ article.views }} 次阅读</div>
              </div>
              <button class="shrink-0 px-2 py-0.5 rounded text-[10px] text-red-400 hover:bg-red-50 hover:text-red-600 transition-all opacity-0 group-hover:opacity-100" @click="handleDeleteArticle(article.id)">删除</button>
            </div>
          </div>
        </div>

        <!-- ===== 项目管理 ===== -->
        <div v-else-if="activeTab === 'projects'" class="max-w-2xl space-y-3">
          <div class="flex items-center justify-between">
            <h2 class="text-sm font-semibold text-gray-800">项目管理</h2>
            <button v-if="!showAddProject" class="px-3 py-1.5 rounded-md bg-gray-900 hover:bg-gray-800 text-white text-xs transition-all" @click="showAddProject = true">+ 新建</button>
          </div>

          <div v-if="showAddProject" class="p-4 rounded-lg bg-white border border-gray-200 space-y-3">
            <input v-model="newProject.name" placeholder="项目名称" class="w-full px-3 py-1.5 rounded border border-gray-200 bg-gray-50 text-xs outline-none focus:border-blue-400" />
            <input v-model="newProject.description" placeholder="项目描述" class="w-full px-3 py-1.5 rounded border border-gray-200 bg-gray-50 text-xs outline-none focus:border-blue-400" />
            <input v-model="newProject.tech" placeholder="技术栈" class="w-full px-3 py-1.5 rounded border border-gray-200 bg-gray-50 text-xs outline-none focus:border-blue-400" />
            <div class="flex justify-end gap-2">
              <button class="px-3 py-1 rounded text-xs text-gray-500 hover:bg-gray-100" @click="showAddProject = false">取消</button>
              <button class="px-3 py-1 rounded text-xs bg-gray-900 hover:bg-gray-800 text-white disabled:opacity-50" :disabled="projectLoading || !newProject.name" @click="handleCreateProject">{{ projectLoading ? '创建中...' : '创建' }}</button>
            </div>
          </div>

          <div v-if="projects.length === 0" class="py-10 text-center text-xs text-gray-300">暂无项目</div>
          <div v-else class="grid gap-2">
            <div v-for="project in projects" :key="project.id" class="flex items-start gap-2 p-3 rounded-lg bg-white border border-gray-100 hover:border-gray-200 transition-all group">
              <div class="flex-1 min-w-0">
                <div class="text-sm font-medium text-gray-800">{{ project.name }}</div>
                <div class="text-xs text-gray-400 mt-0.5">{{ project.description }}</div>
                <span class="inline-block mt-1 text-[10px] px-1.5 py-0.5 rounded bg-gray-100 text-gray-500">{{ project.tech }}</span>
              </div>
              <button class="shrink-0 px-2 py-0.5 rounded text-[10px] text-red-400 hover:bg-red-50 hover:text-red-600 transition-all opacity-0 group-hover:opacity-100" @click="handleDeleteProject(project.id)">删除</button>
            </div>
          </div>
        </div>

        <!-- ===== 个人资料 ===== -->
        <div v-else-if="activeTab === 'profile'" class="max-w-lg space-y-4">
          <h2 class="text-sm font-semibold text-gray-800">个人资料</h2>

          <div class="p-5 rounded-lg bg-white border border-gray-200 space-y-4">
            <div class="flex flex-col items-center py-2">
              <div class="w-14 h-14 rounded-full bg-gray-800 flex items-center justify-center text-white text-lg font-semibold shadow-sm">
                {{ avatarInitial }}
              </div>
            </div>
            <div class="grid grid-cols-2 gap-3">
              <div>
                <label class="block text-[10px] text-gray-400 mb-0.5">姓名</label>
                <input v-model="profileForm.name" class="w-full px-2.5 py-1.5 rounded border border-gray-200 bg-gray-50 text-xs outline-none focus:border-blue-400" />
              </div>
              <div>
                <label class="block text-[10px] text-gray-400 mb-0.5">头衔</label>
                <input v-model="profileForm.title" class="w-full px-2.5 py-1.5 rounded border border-gray-200 bg-gray-50 text-xs outline-none focus:border-blue-400" />
              </div>
            </div>
            <div>
              <label class="block text-[10px] text-gray-400 mb-0.5">个人简介</label>
              <textarea v-model="profileForm.bio" rows="3" class="w-full px-2.5 py-1.5 rounded border border-gray-200 bg-gray-50 text-xs outline-none focus:border-blue-400 resize-none"></textarea>
            </div>
            <div>
              <label class="block text-[10px] text-gray-400 mb-0.5">标签（逗号分隔）</label>
              <input :value="profileForm.tags.join(', ')" @input="(e: any) => { profileForm.tags = e.target.value.split(',').map((t: string) => t.trim()) }" placeholder="Vue 3, TypeScript, Go" class="w-full px-2.5 py-1.5 rounded border border-gray-200 bg-gray-50 text-xs outline-none focus:border-blue-400" />
            </div>
            <button class="w-full py-1.5 rounded bg-gray-900 hover:bg-gray-800 text-white text-xs font-medium transition-all disabled:opacity-50" :disabled="profileLoading" @click="handleSaveProfile">
              {{ profileLoading ? '保存中...' : '保存修改' }}
            </button>
          </div>
        </div>

        <!-- ===== 联系方式 ===== -->
        <div v-else-if="activeTab === 'contacts'" class="max-w-lg space-y-4">
          <h2 class="text-sm font-semibold text-gray-800">联系方式</h2>
          <div class="p-5 rounded-lg bg-white border border-gray-200 space-y-3">
            <div v-for="(item, i) in contactsEdit" :key="i" class="grid grid-cols-2 gap-2">
              <input v-model="item.label" placeholder="标签" class="px-2.5 py-1.5 rounded border border-gray-200 bg-gray-50 text-xs outline-none focus:border-blue-400" />
              <input v-model="item.value" placeholder="内容" class="px-2.5 py-1.5 rounded border border-gray-200 bg-gray-50 text-xs outline-none focus:border-blue-400" />
            </div>
            <button class="w-full py-1.5 rounded bg-gray-900 hover:bg-gray-800 text-white text-xs font-medium transition-all disabled:opacity-50" :disabled="contactsLoading" @click="handleSaveContacts">
              {{ contactsLoading ? '保存中...' : '保存修改' }}
            </button>
          </div>
        </div>

        <!-- ===== 外观设置 ===== -->
        <div v-else-if="activeTab === 'appearance'" class="max-w-lg space-y-4">
          <h2 class="text-sm font-semibold text-gray-800">外观设置</h2>

          <!-- 桌面壁纸 -->
          <div class="p-4 rounded-lg bg-white border border-gray-200 space-y-3">
            <label class="block text-xs text-gray-500 font-medium">桌面壁纸</label>
            <div class="flex gap-2">
              <input
                :value="appSettings.wallpaper"
                @input="(e: any) => setWallpaper(e.target.value)"
                placeholder="输入图片 URL"
                class="flex-1 px-2.5 py-1.5 rounded border border-gray-200 bg-gray-50 text-xs outline-none focus:border-blue-400 font-mono"
              />
            </div>
            <div class="w-full h-24 rounded-lg overflow-hidden bg-gray-100 border border-gray-100">
              <img
                :src="appSettings.wallpaper"
                class="w-full h-full object-cover"
                @error="(e: any) => e.target.style.display = 'none'"
              />
            </div>
            <div class="flex gap-1.5 flex-wrap">
              <button
                v-for="preset in wallpapers"
                :key="preset.url"
                class="w-12 h-8 rounded border-2 overflow-hidden transition-all"
                :class="appSettings.wallpaper === preset.url ? 'border-blue-400' : 'border-transparent hover:border-gray-300'"
                @click="setWallpaper(preset.url)"
              >
                <img :src="preset.url" class="w-full h-full object-cover" />
              </button>
            </div>
          </div>

          <!-- 应用图标 -->
          <div class="p-4 rounded-lg bg-white border border-gray-200 space-y-3">
            <label class="block text-xs text-gray-500 font-medium">应用图标</label>
            <p class="text-[10px] text-gray-300">设置图片 URL 后桌面图标将替换为自定义图片</p>
            <div v-for="app in apps" :key="app.id" class="flex items-center gap-2">
              <span class="text-xs text-gray-500 w-16 shrink-0 truncate">{{ app.title.replace('.app', '') }}</span>
              <input
                :value="appSettings.appIcons[app.id] || ''"
                @input="(e: any) => setAppIcon(app.id, e.target.value)"
                placeholder="图片 URL（留空使用默认图标）"
                class="flex-1 px-2 py-1 rounded border border-gray-200 bg-gray-50 text-[10px] outline-none focus:border-blue-400 font-mono"
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Toast -->
    <Transition name="toast">
      <div v-if="toastMessage" class="absolute bottom-5 left-1/2 -translate-x-1/2 px-3 py-1.5 rounded text-xs text-white shadow-sm z-50" :class="toastType === 'success' ? 'bg-gray-800' : 'bg-red-500'">
        {{ toastMessage }}
      </div>
    </Transition>
  </div>
</template>

<style scoped>
.toast-enter-active, .toast-leave-active { transition: all 0.25s ease; }
.toast-enter-from { opacity: 0; transform: translate(-50%, 8px); }
.toast-leave-to { opacity: 0; transform: translate(-50%, -8px); }
</style>
