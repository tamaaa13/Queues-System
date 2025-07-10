import { createRouter, createWebHistory } from 'vue-router'
import LandingView from '@/views/LandingView.vue'
import LoginView from '@/views/LoginView.vue'
import QueueDisplayView from '@/views/QueueDisplayView.vue'

import DashboardLayout from '@/views/admin/DashboardLayout.vue'
import DashboardView from '@/views/admin/DashboardView.vue'
import QueueListView from '@/views/admin/QueueListView.vue'
import UserListView from '@/views/admin/UserListView.vue'

const routes = [
  {
    path: '/',
    name: 'landing',
    component: LandingView,
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView,
  },
  {
    path: '/queue',
    name: 'queue',
    component: QueueDisplayView,
  },
  {
    path: '/dashboard',
    component: DashboardLayout,
    meta: { requiresAuth: true }, // ✅ ini kuncinya
    children: [
      {
        path: 'beranda',
        name: 'dashboard',
        component: DashboardView,
      },
      {
        path: 'queue-list',
        name: 'queue-list',
        component: QueueListView,
      },
      {
        path: 'users',
        name: 'users',
        component: UserListView,
      },
      {
        path: '',
        redirect: 'beranda',
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

router.beforeEach((to, from, next) => {
  const auth = JSON.parse(sessionStorage.getItem('auth') || '{}')
  const isLoggedIn = !!auth.token
  const userRole = auth.user?.role || ''

  if (to.matched.some((record) => record.meta.requiresAuth)) {
    if (!isLoggedIn) {
      console.log('Not logged in → redirecting to login')
      return next({ name: 'login' })
    }

    // 🚫 Khusus route 'users', hanya superadmin boleh masuk
    if (to.name === 'users' && userRole !== 'superadmin') {
      console.log('Bukan superadmin → redirect ke dashboard')
      return next({ name: 'dashboard' })
    }

    // ✅ Semua aman, lanjut
    return next()
  }

  return next()
})

export default router
