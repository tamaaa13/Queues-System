<script setup>
const props = defineProps({
  name: {
    type: String,
    default: 'Selamat Datang',
  },
})
</script>

<template>
  <div class="w-full bg-transparent border-none py-4 mb-0">
    <div class="text-center wow fadeInUp" data-wow-delay="0.2s">
      <div class="md:w-8/12 xl:w-6/12 mx-auto">
        <h2 class="font-semibold">{{ props.name }}</h2>
      </div>
    </div>
  </div>
</template>
