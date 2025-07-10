<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import Logo from '@/components/Logo.vue'

const currentTime = ref('')
const currentDate = ref('')

const updateTime = () => {
  const now = new Date()
  currentTime.value = now.toLocaleTimeString('id-ID')
  currentDate.value = now.toLocaleDateString('id-ID', {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}

let clockInterval
onMounted(() => {
  updateTime()
  clockInterval = setInterval(updateTime, 1000)
})

onUnmounted(() => {
  clearInterval(clockInterval)
})
</script>

<template>
  <!-- [ Nav ] start -->
  <nav
    class="navbar bg-theme-cardbg dark:bg-themedark-cardbg bg-opacity-75 dark:bg-opacity-75 fixed top-0 z-50 w-full backdrop-blur shadow-[0_0_24px_rgba(27,46,94,.05)] dark:shadow-[0_0_24px_rgba(27,46,94,.05)]"
  >
    <div class="container">
      <div class="static flex py-4 items-center justify-between">
        <div class="flex flex-1 items-center justify-center sm:items-stretch sm:justify-between">
          <Logo flexDirection="row" size="10" />
        </div>
        <div class="flex items-center hidden md:flex">
          <div class="card mb-0 max-w-xs bg-blue-500 text-white">
            <div class="card-body text-center">
              <h4 class="text-white">{{ currentTime }}</h4>
              <p>{{ currentDate }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </nav>
  <!-- [ Nav ] end -->
</template>
