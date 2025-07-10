<script setup>
import Navbar from '@/components/layouts/Navbar.vue'
import Title from '@/components/Title.vue'
import Card from '@/components/Card.vue'
import { ref, onMounted } from 'vue'

// Mapping admin ke ruang, loket, warna, dan urutan tetap
const adminAssignments = {
  adminA1: { room: 'Ruang 1', loket: 'LOKET A', bgLoket: 'bg-red-500', order: 1 },
  adminA2: { room: 'Ruang 4', loket: 'LOKET A', bgLoket: 'bg-red-500', order: 2 },
  adminB1: { room: 'Ruang 2', loket: 'LOKET B', bgLoket: 'bg-green-500', order: 3 },
  adminB2: { room: 'Ruang 3', loket: 'LOKET B', bgLoket: 'bg-green-500', order: 4 },
}

const queueData = ref([])

const fetchQueueData = async () => {
  try {
    const res = await fetch('http://localhost:8800/public/called-queues')
    const results = await res.json() // array of { admin, queue_number, status }

    const updatedData = []

    results.forEach((item) => {
      const adminKey = item.admin
      const queueNumber = item.queue_number || '----'
      const status = item.status || 'waiting'

      const info = adminAssignments[adminKey]
      if (!info) return

      updatedData.push({
        admin: adminKey,
        room: info.room,
        queue: queueNumber,
        loket: info.loket,
        bgLoket: info.bgLoket,
        description:
          status === 'called'
            ? `Nomor antrian ${queueNumber} dipanggil, silakan ke ${info.loket} di ${info.room}`
            : 'Belum ada panggilan',
        order: info.order, // tambahkan order di sini
      })
    })

    // Urutkan berdasarkan order supaya posisi kartu stabil
    queueData.value = updatedData.sort((a, b) => a.order - b.order)
  } catch (err) {
    console.error('Gagal fetch data antrian:', err)
    queueData.value = []
  }
}

onMounted(() => {
  fetchQueueData()
  setInterval(fetchQueueData, 3000) // refresh tiap 3 detik
})
</script>

<template>
  <header
    id="home"
    class="flex items-center flex-col justify-center min-h-screen overflow-hidden relative pt-[50px] lg:pt-[130px] bg-[url('@/assets/images/landing/img-headerbg.jpg')] dark:bg-[url(@/assets/images/landing/img-headerbg-dark.jpg)] bg-cover"
  >
    <Navbar />
    <Title name="Informasi Antrian" />

    <div class="md:container mx-auto">
      <div class="grid grid-cols-2 gap-7 my-8">
        <Card
          v-for="item in queueData"
          :key="item.admin"
          :room="item.room"
          :queue="item.queue"
          :loket="item.loket"
          :bgLoket="item.bgLoket"
          :description="item.description"
        />
      </div>
    </div>
  </header>
</template>
