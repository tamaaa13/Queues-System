<script setup>
const props = defineProps({
  room: {
    type: String,
    default: 'Ruang',
  },
  loket: {
    type: String,
    default: 'LOKET',
  },
  bgLoket: {
    type: String,
    default: 'bg-blue-500',
  },
  queue: {
    type: String,
    default: '----',
  },
  description: {
    type: String,
    default: 'Menunggu',
  },
})
</script>

<template>
  <div
    class="w-full text-white mx-auto rounded-xl shadow-md bg-blue-500/10 dark:bg-themedark-cardbg/10"
  >
    <div
      class="bg-gradient-to-r from-blue-500 to-blue-300 rounded-t-xl text-center p-4 flex flex-col md:flex-row justify-between items-center"
    >
      <h3 class="text-white mb-2 md:mb-0">{{ room }}</h3>
      <span :class="[bgLoket, 'badge p-1 md:p-3']">
        <h6 class="text-white">{{ loket }}</h6>
      </span>
    </div>
    <div class="card-body p-4 rounded-b-xl text-center">
      <h1 class="mb-2">{{ queue }}</h1>
      <p class="text-muted">{{ description }}</p>
    </div>
  </div>
</template>
