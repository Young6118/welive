import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useUserStore } from '@/stores/user'
import type { RouteMeta } from '@/types'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/index.vue'),
    meta: { public: true } as RouteMeta
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/register/index.vue'),
    meta: { public: true } as RouteMeta
  },
  {
    path: '/',
    component: () => import('@/layouts/MainLayout.vue'),
    children: [
      {
        path: 'home',
        name: 'Home',
        component: () => import('@/views/home/index.vue'),
        meta: { title: '首页', keepAlive: true } as RouteMeta
      },
      {
        path: 'notes',
        name: 'Notes',
        component: () => import('@/views/notes/index.vue'),
        meta: { title: '笔记', keepAlive: true } as RouteMeta
      },
      {
        path: 'chat',
        name: 'Chat',
        component: () => import('@/views/chat/index.vue'),
        meta: { title: '聊天', keepAlive: true } as RouteMeta
      },
      {
        path: 'discover',
        name: 'Discover',
        component: () => import('@/views/discover/index.vue'),
        meta: { title: '发现', keepAlive: true } as RouteMeta
      },
      {
        path: 'profile',
        name: 'Profile',
        component: () => import('@/views/profile/index.vue'),
        meta: { title: '我的', keepAlive: true } as RouteMeta
      }
    ]
  },
  {
    path: '/question/:id',
    name: 'QuestionDetail',
    component: () => import('@/views/question/detail.vue'),
    meta: { title: '问题详情' } as RouteMeta
  },
  {
    path: '/question/create',
    name: 'CreateQuestion',
    component: () => import('@/views/question/create.vue'),
    meta: { title: '发布问题' } as RouteMeta
  },
  {
    path: '/note/:id',
    name: 'NoteDetail',
    component: () => import('@/views/note/detail.vue'),
    meta: { title: '笔记详情' } as RouteMeta
  },
  {
    path: '/note/create',
    name: 'CreateNote',
    component: () => import('@/views/note/create.vue'),
    meta: { title: '创建笔记' } as RouteMeta
  },
  {
    path: '/chat/:id',
    name: 'ChatDetail',
    component: () => import('@/views/chat/detail.vue'),
    meta: { title: '聊天' } as RouteMeta
  },
  {
    path: '/village/:id',
    name: 'VillageDetail',
    component: () => import('@/views/village/detail.vue'),
    meta: { title: '村落详情' } as RouteMeta
  },
  {
    path: '/user/:id',
    name: 'UserProfile',
    component: () => import('@/views/user/profile.vue'),
    meta: { title: '用户主页' } as RouteMeta
  },
  {
    path: '/settings',
    name: 'Settings',
    component: () => import('@/views/settings/index.vue'),
    meta: { title: '设置' } as RouteMeta
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  const isPublic = (to.meta as RouteMeta).public

  if (!isPublic && !userStore.isLoggedIn) {
    next('/login')
  } else {
    next()
  }
})

export default router
