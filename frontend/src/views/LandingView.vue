<script setup>
import { useToast } from 'vue-toastification'

const toast = useToast()

async function createQueue(category) {
  try {
    const res = await fetch(`http://localhost:8800/queue/create?category=${category}`, {
      method: 'POST', // ganti ke 'GET' kalau backend kamu pakai GET
    })

    if (!res.ok) {
      throw new Error('Gagal membuat antrean')
    }

    const data = await res.json() // kalau backend mengembalikan JSON
    toast.success('Tiket berhasil dicetak', { timeout: 2000 })
    // console.log(data) // lihat hasil kalau perlu
  } catch (err) {
    toast.error('Gagal mencetak tiket', { timeout: 2000 })
    console.error(err)
  }
}
</script>

<template>
  <header
    id="home"
    class="flex items-center flex-col min-h-screen overflow-hidden relative pt-[130px] bg-[url('@/assets/images/layout/main-bg.jpeg')] dark:bg-[url(@/assets/images/layout/main-bg.jpeg)] bg-cover"
  >
    <div class="absolute inset-0 bg-white/20 z-[1]"></div>
    <div class="container z-[2]">
      <div class="flex items-center justify-center flex-col gap-5">
        <img src="@/assets/images/bekasi-logo.jpg" class="h-auto w-[120px]" alt="" />
        <div class="text-center">
          <h1
            class="text-[22px] md:text-[36px] lg:text-[55px] leading-[1.2] mb-1 wow fadeInUp text-white shadow"
            style="text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.7)"
            data-wow-delay="0.2s"
          >
            LOKET ANTREAN PELAYANAN
          </h1>
          <h1 class="text-white" style="text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.7)">
            KELURAHAN JAKASAMPURNA
          </h1>
        </div>
        <div class="wow fadeInUp" data-wow-delay="0.4s">
          <button class="btn btn-primary max-w-md mt-2 mx-2" @click="createQueue('A')">
            <h2 class="text-white my-4">Layanan Kependudukan Dasar</h2>
          </button>
          <button class="btn btn-primary max-w-md mt-2 mx-2" @click="createQueue('B')">
            <h2 class="text-white my-4">Layanan Surat Keterangan dan Domisili</h2>
          </button>
        </div>
      </div>
    </div>
  </header>
</template>
