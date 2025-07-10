<script setup>
const props = defineProps({
  size: {
    type: String,
    default: '14',
  },
  flexDirection: {
    type: String,
    default: 'col',
  },
  fontSize: {
    type: String,
    default: 'base',
  },
})
</script>

<template>
  <div
    :class="[
      'flex items-center justify-center gap-2',
      props.flexDirection === 'row' ? 'flex-row' : 'flex-col',
    ]"
  >
    <img src="@/assets/images/bekasi-logo.jpg" alt="img" :class="['mx-auto size-' + props.size]" />
    <div :class="[flexDirection === 'row' ? 'text-left' : 'text-center']">
      <h5 :class="['text-' + props.fontSize]">PELAYANAN KEPENDUDUKAN</h5>
      <h5 :class="['text-' + props.fontSize]">KELURAHAN JAKASAMPURNA</h5>
    </div>
  </div>
</template>
