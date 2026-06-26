<script setup lang="ts">
/**
 * AdminTerminal.vue - "后台管理.app" 窗口内容
 *
 * 仿 macOS Terminal 风格的管理控制台
 * 通过命令行的方式管理文章、项目、个人资料等内容
 *
 * 你需要在后端实现以下 API：
 *   POST   /api/login              → { "token": "..." }
 *   POST   /api/articles           → 创建文章
 *   DELETE /api/articles/:id       → 删除文章
 *   POST   /api/projects           → 创建项目
 *   DELETE /api/projects/:id       → 删除项目
 *   PUT    /api/profile            → 更新个人资料
 *   PUT    /api/contact            → 更新联系方式
 */
import { ref, nextTick, onMounted } from 'vue'

// ============ 终端状态 ============

/** 终端输出的所有行 */
const lines = ref<string[]>([])
/** 当前输入内容 */
const input = ref('')
/** 输入框引用（用于自动聚焦） */
const inputRef = ref<HTMLInputElement | null>(null)
/** 终端容器引用（用于自动滚动到底部） */
const terminalRef = ref<HTMLDivElement | null>(null)
/** 是否已登录 */
const isLoggedIn = ref(false)
/** 认证 token */
const token = ref('')
/** 命令历史（上下键切换） */
const history = ref<string[]>([])
const historyIndex = ref(-1)

// ============ 初始化 ============

onMounted(() => {
  printLine('⏣ TUDO\'s Blog 管理终端 v0.1')
  printLine(' type "help" 查看可用命令')
  printLine(' type "login <password>" 登录')
  printLine('─'.repeat(40))
  printPrompt()
})

// ============ 终端渲染辅助 ============

function printLine(text: string) {
  lines.value.push(text)
}

function printPrompt() {
  const user = isLoggedIn.value ? 'admin' : 'guest'
  lines.value.push(`[${user}@todo-blog ~]$ `)
}

function scrollToBottom() {
  nextTick(() => {
    if (terminalRef.value) {
      terminalRef.value.scrollTop = terminalRef.value.scrollHeight
    }
  })
}

// ============ 命令处理 ============

async function handleCommand() {
  const cmd = input.value.trim()
  input.value = ''
  if (!cmd) return

  // 记录命令到历史
  history.value.push(cmd)
  historyIndex.value = -1

  // 在终端中回显命令
  const lastLine = lines.value[lines.value.length - 1]
  lines.value[lines.value.length - 1] = lastLine + cmd

  // 解析命令
  const parts = cmd.split(/\s+/)
  const command = parts[0].toLowerCase()
  const args = parts.slice(1)

  await executeCommand(command, args)

  scrollToBottom()
}

async function executeCommand(command: string, args: string[]) {
  switch (command) {
    case 'help':     return cmdHelp()
    case 'clear':    return cmdClear()
    case 'login':    return cmdLogin(args)
    case 'logout':   return cmdLogout()
    case 'articles': return cmdArticles()
    case 'article':  return cmdArticle(args)
    case 'projects': return cmdProjects()
    case 'project':  return cmdProject(args)
    default:
      printLine(`未知命令: ${command}，输入 "help" 查看可用命令`)
      printPrompt()
  }
}

// ============ 命令实现 ============

function cmdHelp() {
  printLine('')
  printLine('  可用命令:')
  printLine('  ───────────────────────────────────────────')
  printLine('  login <password>    登录管理终端')
  printLine('  logout              退出登录')
  printLine('  articles            查看所有文章')
  printLine('  article create      创建文章 (交互式)')
  printLine('  article delete <id> 删除文章')
  printLine('  projects            查看所有项目')
  printLine('  project create      创建项目 (交互式)')
  printLine('  project delete <id> 删除项目')
  if (isLoggedIn.value) {
    printLine('  profile update      更新个人资料 (交互式)')
    printLine('  contact update      更新联系方式 (交互式)')
  }
  printLine('  clear               清屏')
  printLine('  help                显示此帮助')
  printLine('  ───────────────────────────────────────────')
  printLine('')
  printPrompt()
}

function cmdClear() {
  lines.value = []
  printPrompt()
}

async function cmdLogin(args: string[]) {
  if (args.length === 0) {
    printLine('用法: login <password>')
    printPrompt()
    return
  }

  try {
    const res = await fetch('http://localhost:8080/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ password: args[0] }),
    })

    if (!res.ok) {
      printLine('认证失败: 密码错误')
      printPrompt()
      return
    }

    const data = await res.json()
    token.value = data.token
    isLoggedIn.value = true
    printLine('✓ 登录成功，欢迎回来')
    printPrompt()
  } catch {
    printLine('错误: 无法连接到后端服务')
    printPrompt()
  }
}

function cmdLogout() {
  token.value = ''
  isLoggedIn.value = false
  printLine('已退出登录')
  printPrompt()
}

async function cmdArticles() {
  if (!isLoggedIn.value) { printLine('请先登录: login <password>'); printPrompt(); return }

  try {
    const res = await fetch('http://localhost:8080/api/articles')
    const articles = await res.json()

    if (articles.length === 0) {
      printLine('暂无文章')
      printPrompt()
      return
    }

    printLine('')
    printLine(`  ID  │ 标题                          │ 阅读量`)
    printLine(`  ────┼───────────────────────────────┼──────`)
    for (const a of articles) {
      const id = String(a.id).padEnd(4)
      const title = (a.title || '').substring(0, 30).padEnd(30)
      const views = String(a.views).padStart(4)
      printLine(`  ${id}│ ${title} │ ${views}`)
    }
    printLine('')
    printPrompt()
  } catch {
    printLine('错误: 无法获取文章列表')
    printPrompt()
  }
}

async function cmdArticle(args: string[]) {
  if (!isLoggedIn.value) { printLine('请先登录: login <password>'); printPrompt(); return }

  if (args.length === 0) {
    printLine('用法: article create 或 article delete <id>')
    printPrompt()
    return
  }

  const sub = args[0].toLowerCase()

  if (sub === 'create') {
    // 交互式创建 —— 提示用户输入
    printLine('请输入文章信息:')
    printLine('')

    // 用一个简单粗暴的方式：弹三个问题
    const title = prompt('标题: ')
    const summary = prompt('摘要: ')
    if (!title || !summary) {
      printLine('取消创建')
      printPrompt()
      return
    }

    try {
      const res = await fetch('http://localhost:8080/api/articles', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token.value}`,
        },
        body: JSON.stringify({ title, summary }),
      })

      if (res.ok) {
        printLine(`✓ 文章「${title}」创建成功`)
      } else {
        printLine('创建失败')
      }
    } catch {
      printLine('错误: 无法连接到后端')
    }
    printPrompt()
    return
  }

  if (sub === 'delete') {
    const id = args[1]
    if (!id) { printLine('用法: article delete <id>'); printPrompt(); return }

    try {
      const res = await fetch(`http://localhost:8080/api/articles/${id}`, {
        method: 'DELETE',
        headers: { 'Authorization': `Bearer ${token.value}` },
      })

      if (res.ok) {
        printLine(`✓ 文章 #${id} 已删除`)
      } else {
        printLine('删除失败: 文章不存在或无权操作')
      }
    } catch {
      printLine('错误: 无法连接到后端')
    }
    printPrompt()
    return
  }

  printLine(`未知子命令: ${sub}`)
  printPrompt()
}

async function cmdProjects() {
  if (!isLoggedIn.value) { printLine('请先登录: login <password>'); printPrompt(); return }

  try {
    const res = await fetch('http://localhost:8080/api/projects')
    const projects = await res.json()

    if (projects.length === 0) {
      printLine('暂无项目')
      printPrompt()
      return
    }

    printLine('')
    printLine(`  ID  │ 项目名称                       │ 技术栈`)
    printLine(`  ────┼───────────────────────────────┼──────────`)
    for (const p of projects) {
      const id = String(p.id).padEnd(4)
      const name = (p.name || '').substring(0, 30).padEnd(30)
      const tech = (p.tech || '').substring(0, 20)
      printLine(`  ${id}│ ${name} │ ${tech}`)
    }
    printLine('')
    printPrompt()
  } catch {
    printLine('错误: 无法获取项目列表')
    printPrompt()
  }
}

async function cmdProject(args: string[]) {
  if (!isLoggedIn.value) { printLine('请先登录: login <password>'); printPrompt(); return }

  if (args.length === 0) {
    printLine('用法: project create 或 project delete <id>')
    printPrompt()
    return
  }

  const sub = args[0].toLowerCase()

  if (sub === 'create') {
    const name = prompt('项目名称: ')
    const desc = prompt('项目描述: ')
    const tech = prompt('技术栈: ')
    if (!name) { printLine('取消创建'); printPrompt(); return }

    try {
      const res = await fetch('http://localhost:8080/api/projects', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token.value}`,
        },
        body: JSON.stringify({ name, description: desc, tech }),
      })

      if (res.ok) {
        printLine(`✓ 项目「${name}」创建成功`)
      } else {
        printLine('创建失败')
      }
    } catch {
      printLine('错误: 无法连接到后端')
    }
    printPrompt()
    return
  }

  if (sub === 'delete') {
    const id = args[1]
    if (!id) { printLine('用法: project delete <id>'); printPrompt(); return }

    try {
      const res = await fetch(`http://localhost:8080/api/projects/${id}`, {
        method: 'DELETE',
        headers: { 'Authorization': `Bearer ${token.value}` },
      })

      if (res.ok) {
        printLine(`✓ 项目 #${id} 已删除`)
      } else {
        printLine('删除失败')
      }
    } catch {
      printLine('错误: 无法连接到后端')
    }
    printPrompt()
    return
  }

  printLine(`未知子命令: ${sub}`)
  printPrompt()
}

// ============ 键盘事件 ============

function onKeydown(e: KeyboardEvent) {
  if (e.key === 'ArrowUp') {
    e.preventDefault()
    if (history.value.length === 0) return
    historyIndex.value = Math.max(0, historyIndex.value - 1)
    input.value = history.value[historyIndex.value] || ''
  } else if (e.key === 'ArrowDown') {
    e.preventDefault()
    if (historyIndex.value === -1) return
    historyIndex.value++
    if (historyIndex.value >= history.value.length) {
      historyIndex.value = -1
      input.value = ''
    } else {
      input.value = history.value[historyIndex.value]
    }
  }
}

function focusInput() {
  inputRef.value?.focus()
}
</script>

<template>
  <!-- macOS Terminal 风格容器 -->
  <div
    ref="terminalRef"
    class="
      h-full w-full overflow-y-auto p-3
      font-mono text-sm leading-relaxed
      bg-gray-950 text-green-400
      cursor-text
    "
    @click="focusInput"
  >
    <!-- 输出行 -->
    <div v-for="(line, i) in lines" :key="i" class="whitespace-pre-wrap break-all">
      <span v-if="line.startsWith('[')" class="text-green-400/70">{{ line }}</span>
      <span v-else-if="line.startsWith('  ID') || line.startsWith('  ──')" class="text-gray-500">
        {{ line }}
      </span>
      <span v-else-if="line.startsWith('  ') || line.startsWith('未知') || line.startsWith('错误') || line.startsWith('请先')" class="text-gray-400">
        {{ line }}
      </span>
      <span v-else-if="line.startsWith('✓')" class="text-green-400">
        {{ line }}
      </span>
      <span v-else-if="line.startsWith('认证失败')" class="text-red-400">
        {{ line }}
      </span>
      <span v-else>{{ line }}</span>
    </div>

    <!-- 输入行 -->
    <div class="flex items-center gap-1">
      <span class="text-green-400/70 shrink-0">
        {{ isLoggedIn ? '[admin@todo-blog ~]$' : '[guest@todo-blog ~]$' }}
      </span>
      <input
        ref="inputRef"
        v-model="input"
        class="
          flex-1 bg-transparent border-none outline-none
          text-green-400 font-mono text-sm
          caret-green-400
        "
        @keydown.enter="handleCommand"
        @keydown="onKeydown"
        autofocus
      />
    </div>
  </div>
</template>
