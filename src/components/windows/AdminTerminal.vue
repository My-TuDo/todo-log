<script setup lang="ts">
/**
 * AdminTerminal.vue - "后台管理.app" 窗口内容
 *
 * 图形化后台管理界面
 * - 文章管理：列表查看、创建、删除
 * - 项目管理：列表查看、创建、删除
 * - 个人资料：查看与编辑
 * - 联系方式：查看与编辑
 * - 所有数据通过 API 服务层操作，支持 Mock 数据
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
import type { Article, Project, Profile, ContactItem } from '../../types'

// ============================================================
// 认证状态
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
// 标签导航
// ============================================================

type Tab = 'articles' | 'projects' | 'profile' | 'contacts'
const activeTab = ref<Tab>('articles')

const tabs: { id: Tab; label: string; icon: string }[] = [
  { id: 'articles', label: '文章管理', icon: '📝' },
  { id: 'projects', label: '项目管理', icon: '💻' },
  { id: 'profile', label: '个人资料', icon: '👤' },
  { id: 'contacts', label: '联系方式', icon: '📬' },
]

// ============================================================
// 数据状态
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
    // 刷新列表
    articles.value = await fetchArticles()
  } catch {
    showToast('创建失败', 'error')
  } finally {
    articleLoading.value = false
  }
}

async function handleDeleteArticle(id: number) {
  if (!confirm('确定要删除这篇文文章吗？')) return
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
// 个人资料编辑
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

// ============================================================
// 联系方式编辑
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

// 切换到编辑相关的 tab 时初始化表单
const previousTab = ref<Tab>('articles')
const activeTabComputed = computed({
  get: () => activeTab.value,
  set: (val: Tab) => {
    previousTab.value = activeTab.value
    activeTab.value = val
    if (val === 'profile') initProfileForm()
    if (val === 'contacts') initContactsForm()
  },
})
</script>

<template>
  <!-- ============================================================ -->
  <!-- 登录页面 -->
  <!-- ============================================================ -->
  <div v-if="!isLoggedIn" class="h-full w-full flex items-center justify-center bg-gradient-to-br from-gray-50 to-gray-100">
    <div class="w-80 p-8 rounded-2xl bg-white shadow-xl border border-gray-100">
      <div class="text-center mb-6">
        <div class="text-4xl mb-2">🔐</div>
        <h2 class="text-lg font-bold text-gray-800">后台管理</h2>
        <p class="text-xs text-gray-400 mt-1">请输入密码登录</p>
      </div>

      <form @submit.prevent="handleLogin" class="space-y-4">
        <div>
          <input
            v-model="loginPassword"
            type="password"
            placeholder="管理员密码"
            class="w-full px-4 py-2.5 rounded-xl border border-gray-200 bg-gray-50 text-sm text-gray-700 placeholder-gray-400 outline-none focus:ring-2 focus:ring-blue-400 focus:border-transparent transition-all"
            :class="{ 'border-red-300 ring-1 ring-red-300': loginError }"
          />
          <p v-if="loginError" class="text-xs text-red-500 mt-1.5">{{ loginError }}</p>
        </div>
        <button
          type="submit"
          class="w-full py-2.5 rounded-xl bg-blue-500 hover:bg-blue-600 active:bg-blue-700 text-white text-sm font-medium transition-all disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="loginLoading"
        >
          {{ loginLoading ? '登录中...' : '登 录' }}
        </button>
      </form>

      <p class="text-center text-xs text-gray-300 mt-4">Mock 模式默认密码: admin123</p>
    </div>
  </div>

  <!-- ============================================================ -->
  <!-- 管理主界面 -->
  <!-- ============================================================ -->
  <div v-else class="h-full w-full flex flex-col bg-gray-50">
    <!-- 顶部栏 -->
    <div class="flex items-center justify-between px-5 py-3 bg-white border-b border-gray-200 shrink-0">
      <div class="flex items-center gap-3">
        <span class="text-lg">⚙️</span>
        <h1 class="text-sm font-bold text-gray-800">后台管理</h1>
        <span class="text-[10px] px-2 py-0.5 rounded-full bg-green-50 text-green-600 border border-green-200">已登录</span>
      </div>
      <div class="flex items-center gap-3">
        <span class="text-xs text-gray-400">admin</span>
        <button
          class="px-3 py-1.5 rounded-lg text-xs text-gray-500 hover:bg-gray-100 hover:text-gray-700 transition-colors"
          @click="handleLogout"
        >
          退出
        </button>
      </div>
    </div>

    <!-- 主体区域 -->
    <div class="flex-1 flex overflow-hidden">
      <!-- 侧边导航 -->
      <div class="w-44 shrink-0 bg-white border-r border-gray-200 p-2 space-y-1">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          class="w-full flex items-center gap-2.5 px-3 py-2.5 rounded-lg text-sm text-left transition-all"
          :class="activeTab === tab.id
            ? 'bg-blue-50 text-blue-700 font-medium'
            : 'text-gray-600 hover:bg-gray-50 hover:text-gray-800'"
          @click="activeTabComputed = tab.id"
        >
          <span>{{ tab.icon }}</span>
          <span>{{ tab.label }}</span>
        </button>
      </div>

      <!-- 内容区 -->
      <div class="flex-1 overflow-y-auto p-5">
        <!-- 加载中 -->
        <div v-if="loading" class="flex items-center justify-center h-full text-gray-400 text-sm">
          <span class="animate-pulse">加载中...</span>
        </div>

        <!-- ======================== 文章管理 ======================== -->
        <div v-else-if="activeTab === 'articles'" class="space-y-4">
          <div class="flex items-center justify-between">
            <h2 class="text-base font-bold text-gray-800">📝 文章管理</h2>
            <button
              v-if="!showAddArticle"
              class="flex items-center gap-1.5 px-3 py-1.5 rounded-lg bg-blue-500 hover:bg-blue-600 text-white text-xs font-medium transition-all"
              @click="showAddArticle = true"
            >
              <span class="text-base leading-none">＋</span>
              新建文章
            </button>
          </div>

          <!-- 新建文章表单 -->
          <div v-if="showAddArticle" class="p-4 rounded-xl bg-white border border-gray-200 shadow-sm space-y-3">
            <input
              v-model="newArticle.title"
              placeholder="文章标题"
              class="w-full px-3 py-2 rounded-lg border border-gray-200 bg-gray-50 text-sm outline-none focus:ring-2 focus:ring-blue-400 focus:border-transparent"
            />
            <textarea
              v-model="newArticle.summary"
              placeholder="文章摘要"
              rows="3"
              class="w-full px-3 py-2 rounded-lg border border-gray-200 bg-gray-50 text-sm outline-none focus:ring-2 focus:ring-blue-400 focus:border-transparent resize-none"
            ></textarea>
            <div class="flex justify-end gap-2">
              <button
                class="px-3 py-1.5 rounded-lg text-xs text-gray-500 hover:bg-gray-100 transition-colors"
                @click="showAddArticle = false"
              >取消</button>
              <button
                class="px-4 py-1.5 rounded-lg bg-blue-500 hover:bg-blue-600 text-white text-xs font-medium transition-all disabled:opacity-50"
                :disabled="articleLoading || !newArticle.title"
                @click="handleCreateArticle"
              >{{ articleLoading ? '创建中...' : '创建' }}</button>
            </div>
          </div>

          <!-- 文章列表 -->
          <div v-if="articles.length === 0" class="py-12 text-center text-gray-400 text-sm">
            暂无文章，点击"新建文章"开始写作
          </div>
          <div v-else class="space-y-2">
            <div
              v-for="article in articles"
              :key="article.id"
              class="flex items-start gap-3 p-3 rounded-xl bg-white border border-gray-100 hover:border-gray-200 hover:shadow-sm transition-all group"
            >
              <div class="flex-1 min-w-0">
                <h3 class="text-sm font-medium text-gray-800 truncate">{{ article.title }}</h3>
                <p class="text-xs text-gray-400 mt-0.5 line-clamp-1">{{ article.summary }}</p>
                <div class="flex items-center gap-3 mt-1.5 text-[10px] text-gray-300">
                  <span>📅 {{ article.created_at }}</span>
                  <span>👁️ {{ article.views }} 次阅读</span>
                </div>
              </div>
              <button
                class="shrink-0 px-2 py-1 rounded-md text-[10px] text-red-400 hover:bg-red-50 hover:text-red-600 transition-all opacity-0 group-hover:opacity-100"
                @click="handleDeleteArticle(article.id)"
              >删除</button>
            </div>
          </div>
        </div>

        <!-- ======================== 项目管理 ======================== -->
        <div v-else-if="activeTab === 'projects'" class="space-y-4">
          <div class="flex items-center justify-between">
            <h2 class="text-base font-bold text-gray-800">💻 项目管理</h2>
            <button
              v-if="!showAddProject"
              class="flex items-center gap-1.5 px-3 py-1.5 rounded-lg bg-blue-500 hover:bg-blue-600 text-white text-xs font-medium transition-all"
              @click="showAddProject = true"
            >
              <span class="text-base leading-none">＋</span>
              新建项目
            </button>
          </div>

          <!-- 新建项目表单 -->
          <div v-if="showAddProject" class="p-4 rounded-xl bg-white border border-gray-200 shadow-sm space-y-3">
            <input v-model="newProject.name" placeholder="项目名称" class="w-full px-3 py-2 rounded-lg border border-gray-200 bg-gray-50 text-sm outline-none focus:ring-2 focus:ring-blue-400 focus:border-transparent" />
            <input v-model="newProject.description" placeholder="项目描述" class="w-full px-3 py-2 rounded-lg border border-gray-200 bg-gray-50 text-sm outline-none focus:ring-2 focus:ring-blue-400 focus:border-transparent" />
            <input v-model="newProject.tech" placeholder="技术栈" class="w-full px-3 py-2 rounded-lg border border-gray-200 bg-gray-50 text-sm outline-none focus:ring-2 focus:ring-blue-400 focus:border-transparent" />
            <div class="flex justify-end gap-2">
              <button class="px-3 py-1.5 rounded-lg text-xs text-gray-500 hover:bg-gray-100 transition-colors" @click="showAddProject = false">取消</button>
              <button class="px-4 py-1.5 rounded-lg bg-blue-500 hover:bg-blue-600 text-white text-xs font-medium transition-all disabled:opacity-50" :disabled="projectLoading || !newProject.name" @click="handleCreateProject">{{ projectLoading ? '创建中...' : '创建' }}</button>
            </div>
          </div>

          <div v-if="projects.length === 0" class="py-12 text-center text-gray-400 text-sm">暂无项目</div>
          <div v-else class="grid gap-3">
            <div v-for="project in projects" :key="project.id" class="flex items-start gap-3 p-3 rounded-xl bg-white border border-gray-100 hover:border-gray-200 hover:shadow-sm transition-all group">
              <div class="flex-1 min-w-0">
                <h3 class="text-sm font-medium text-gray-800">{{ project.name }}</h3>
                <p class="text-xs text-gray-400 mt-0.5">{{ project.description }}</p>
                <span class="inline-block mt-1.5 text-[10px] px-2 py-0.5 rounded-full bg-blue-50 text-blue-500 border border-blue-100">{{ project.tech }}</span>
              </div>
              <button class="shrink-0 px-2 py-1 rounded-md text-[10px] text-red-400 hover:bg-red-50 hover:text-red-600 transition-all opacity-0 group-hover:opacity-100" @click="handleDeleteProject(project.id)">删除</button>
            </div>
          </div>
        </div>

        <!-- ======================== 个人资料 ======================== -->
        <div v-else-if="activeTab === 'profile'" class="max-w-lg space-y-4">
          <h2 class="text-base font-bold text-gray-800">👤 个人资料</h2>

          <div class="p-5 rounded-xl bg-white border border-gray-200 shadow-sm space-y-4">
            <div>
              <label class="block text-xs text-gray-400 mb-1">头像 Emoji</label>
              <input v-model="profileForm.avatar_emoji" class="w-full px-3 py-2 rounded-lg border border-gray-200 bg-gray-50 text-sm outline-none focus:ring-2 focus:ring-blue-400" />
            </div>
            <div>
              <label class="block text-xs text-gray-400 mb-1">姓名</label>
              <input v-model="profileForm.name" class="w-full px-3 py-2 rounded-lg border border-gray-200 bg-gray-50 text-sm outline-none focus:ring-2 focus:ring-blue-400" />
            </div>
            <div>
              <label class="block text-xs text-gray-400 mb-1">头衔</label>
              <input v-model="profileForm.title" class="w-full px-3 py-2 rounded-lg border border-gray-200 bg-gray-50 text-sm outline-none focus:ring-2 focus:ring-blue-400" />
            </div>
            <div>
              <label class="block text-xs text-gray-400 mb-1">个人简介</label>
              <textarea v-model="profileForm.bio" rows="4" class="w-full px-3 py-2 rounded-lg border border-gray-200 bg-gray-50 text-sm outline-none focus:ring-2 focus:ring-blue-400 resize-none"></textarea>
            </div>
            <div>
              <label class="block text-xs text-gray-400 mb-1">标签（逗号分隔）</label>
              <input
                :value="profileForm.tags.join(', ')"
                @input="(e: any) => { profileForm.tags = e.target.value.split(',').map((t: string) => t.trim()) }"
                placeholder="Vue 3, TypeScript, Go"
                class="w-full px-3 py-2 rounded-lg border border-gray-200 bg-gray-50 text-sm outline-none focus:ring-2 focus:ring-blue-400"
              />
            </div>
            <button class="w-full py-2 rounded-lg bg-blue-500 hover:bg-blue-600 text-white text-sm font-medium transition-all disabled:opacity-50" :disabled="profileLoading" @click="handleSaveProfile">
              {{ profileLoading ? '保存中...' : '保存修改' }}
            </button>
          </div>

          <!-- 预览 -->
          <div class="p-5 rounded-xl bg-white border border-gray-200 shadow-sm">
            <h3 class="text-xs text-gray-400 mb-3">👁️ 预览</h3>
            <div class="flex flex-col items-center py-3 space-y-2">
              <div class="w-16 h-16 rounded-full bg-gradient-to-br from-blue-400 to-purple-500 flex items-center justify-center text-2xl text-white shadow-md">
                {{ profileForm.avatar_emoji || '?' }}
              </div>
              <div class="text-center">
                <p class="text-sm font-semibold text-gray-800">{{ profileForm.name || '(未填写)' }}</p>
                <p class="text-xs text-gray-400">{{ profileForm.title || '(未填写)' }}</p>
              </div>
            </div>
            <p class="text-xs text-gray-500 text-center leading-relaxed">{{ profileForm.bio || '(未填写)' }}</p>
          </div>
        </div>

        <!-- ======================== 联系方式 ======================== -->
        <div v-else-if="activeTab === 'contacts'" class="max-w-lg space-y-4">
          <h2 class="text-base font-bold text-gray-800">📬 联系方式</h2>

          <div class="p-5 rounded-xl bg-white border border-gray-200 shadow-sm space-y-4">
            <div v-for="(item, i) in contactsEdit" :key="i" class="flex items-center gap-3">
              <span class="text-xl shrink-0">{{ item.icon }}</span>
              <div class="flex-1 grid grid-cols-2 gap-2">
                <input v-model="item.label" placeholder="标签" class="px-3 py-2 rounded-lg border border-gray-200 bg-gray-50 text-xs outline-none focus:ring-2 focus:ring-blue-400" />
                <input v-model="item.value" placeholder="内容" class="px-3 py-2 rounded-lg border border-gray-200 bg-gray-50 text-xs outline-none focus:ring-2 focus:ring-blue-400" />
              </div>
            </div>
            <button class="w-full py-2 rounded-lg bg-blue-500 hover:bg-blue-600 text-white text-sm font-medium transition-all disabled:opacity-50" :disabled="contactsLoading" @click="handleSaveContacts">
              {{ contactsLoading ? '保存中...' : '保存修改' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Toast 消息 -->
    <Transition name="toast">
      <div
        v-if="toastMessage"
        class="absolute bottom-6 left-1/2 -translate-x-1/2 px-4 py-2 rounded-lg text-xs text-white shadow-lg z-50"
        :class="toastType === 'success' ? 'bg-green-500' : 'bg-red-500'"
      >
        {{ toastMessage }}
      </div>
    </Transition>
  </div>
</template>

<style scoped>
.toast-enter-active,
.toast-leave-active {
  transition: all 0.3s ease;
}
.toast-enter-from {
  opacity: 0;
  transform: translate(-50%, 10px);
}
.toast-leave-to {
  opacity: 0;
  transform: translate(-50%, -10px);
}
</style>
