<script setup>
import Sidebar from '@/components/layouts/Sidebar.vue'
import Header from '@/components/layouts/Header.vue'
</script>

<template>
  <Sidebar />
  <Header />

  <!-- [ Main Content ] start -->
  <div class="pc-container">
    <div class="pc-content">
      <slot />
    </div>
  </div>
</template>
