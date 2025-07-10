<script setup>
const props = defineProps({
  title: {
    type: String,
    default: '',
  },
  items: {
    type: Array,
    required: true,
    validator: (items) => {
      return items.every((item) => typeof item.text === 'string')
    },
  },
})
</script>

<template>
  <!-- [ breadcrumb ] start -->
  <div class="page-header">
    <div class="page-block">
      <ul class="breadcrumb">
        <li v-for="(item, index) in items" :key="index" class="breadcrumb-item">
          <a v-if="item.link && index !== items.length - 1" :href="item.link">{{ item.text }}</a>
          <span v-else :aria-current="index === items.length - 1 ? 'page' : null">{{
            item.text
          }}</span>
        </li>
      </ul>
      <div v-if="title" class="page-header-title">
        <h2 class="mb-0">{{ title }}</h2>
      </div>
    </div>
  </div>
  <!-- [ breadcrumb ] end -->
</template>
