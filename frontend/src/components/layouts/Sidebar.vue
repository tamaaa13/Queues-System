<script setup>
import Logo from '@/components/Logo.vue'
import { useRoute } from 'vue-router'

const route = useRoute()
// Ambil user dari sessionStorage
const auth = JSON.parse(sessionStorage.getItem('auth') || '{}')
const user = auth.user || {}
</script>

<template>
  <nav class="pc-sidebar">
    <div class="navbar-wrapper">
      <div class="m-header flex items-center py-4 px-6 h-header-height">
        <Logo flexDirection="row" size="10" fontSize="sm" />
      </div>
      <div class="navbar-content h-[calc(100vh_-_74px)] py-2.5">
        <ul class="pc-navbar">
          <li :class="['pc-item', route.path === '/dashboard/beranda' ? 'active' : '']">
            <router-link to="/dashboard/beranda" class="pc-link">
              <span class="pc-micon">
                <i class="ti ti-home"></i>
              </span>
              <span class="pc-mtext">Dashboard</span>
            </router-link>
          </li>
          <li :class="['pc-item', route.path === '/dashboard/queue-list' ? 'active' : '']">
            <router-link to="/dashboard/queue-list" class="pc-link">
              <span class="pc-micon">
                <i class="ti ti-list"></i>
              </span>
              <span class="pc-mtext">List Antrian</span>
            </router-link>
          </li>
          <li
            v-if="user?.role === 'superadmin'"
            :class="['pc-item', route.path === '/dashboard/users' ? 'active' : '']"
          >
            <router-link to="/dashboard/users" class="pc-link">
              <span class="pc-micon">
                <i class="ti ti-user"></i>
              </span>
              <span class="pc-mtext">Users</span>
            </router-link>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>
