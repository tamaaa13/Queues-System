<script setup>
import Breadcrumb from '@/components/Breadcrumb.vue'
import { ref, computed, onMounted } from 'vue'

// Ambil dari sessionStorage
const authData = JSON.parse(sessionStorage.getItem('auth') || '{}')
const user = authData.user || {}
const token = authData.token || ''

const username = user?.username || ''
const userCategory = user?.category?.String || ''
const userRole = user?.role || ''

// Mapping username ke ruang dan loket
const adminAssignments = {
  adminA1: { room: 'Ruang 1', loket: 'LOKET A' },
  adminA2: { room: 'Ruang 4', loket: 'LOKET A' },
  adminB1: { room: 'Ruang 2', loket: 'LOKET B' },
  adminB2: { room: 'Ruang 3', loket: 'LOKET B' },
}
const assignment = adminAssignments[username] || {}
const assignedRoom = assignment.room
const assignedLoket = assignment.loket

const queueData = ref([])
const selectedLoket = ref('Semua')
const loketOptions = ref(['Semua', 'A', 'B'])
const lastCalledNumber = ref('')

// Fungsi untuk memutar suara
const playCallSound = (queueNumber, loket, room) => {
  const msg = new SpeechSynthesisUtterance(
    `Nomor antrian ${queueNumber}, silakan ke ruang ${room} ${loket},`,
  )
  msg.lang = 'id-ID'
  msg.rate = 1
  msg.pitch = 1
  window.speechSynthesis.cancel()
  window.speechSynthesis.speak(msg)
}

// Ambil data antrian
const fetchQueueData = async () => {
  try {
    const response = await fetch('http://localhost:8800/admin/called-queues', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    })

    const result = await response.json()
    queueData.value = result

    // Cek dan mainkan suara jika perlu
    if (userRole !== 'superadmin') {
      const latest = [...result]
        .reverse()
        .find(
          (q) =>
            q.Status === 'called' &&
            q.Category === userCategory &&
            q.Room === assignedRoom &&
            q.Loket === assignedLoket,
        )

      if (latest && latest.Number !== lastCalledNumber.value) {
        playCallSound(latest.Number, assignedLoket)
        lastCalledNumber.value = latest.Number
      }
    }
  } catch (error) {
    console.error('Gagal mengambil data antrian:', error)
  }
}

const filteredByCategory = computed(() => {
  return userRole === 'superadmin'
    ? queueData.value
    : queueData.value.filter((q) => q.Category === userCategory)
})

const filteredQueue = computed(() => {
  return selectedLoket.value === 'Semua'
    ? filteredByCategory.value
    : filteredByCategory.value.filter((q) => q.Category === selectedLoket.value)
})

const currentQueue = computed(() => {
  return [...filteredByCategory.value].reverse().find((q) => q.Status === 'called') || null
})

const currentQueueA = computed(
  () =>
    [...queueData.value].reverse().find((q) => q.Status === 'called' && q.Category === 'A') || null,
)

const currentQueueB = computed(
  () =>
    [...queueData.value].reverse().find((q) => q.Status === 'called' && q.Category === 'B') || null,
)

const nextQueue = async () => {
  if (userRole === 'superadmin') return
  try {
    const response = await fetch(`http://localhost:8800/admin/call-next-queue/${userCategory}`, {
      method: 'GET',
      headers: { Authorization: `Bearer ${token}` },
    })

    if (!response.ok) throw new Error('Gagal memanggil antrian')

    const result = await response.json()

    // Langsung mainkan suara jika assignedLoket tersedia
    if (result.queue_number && assignedLoket && assignedRoom) {
      playCallSound(result.queue_number, assignedLoket, assignedRoom)
      lastCalledNumber.value = result.queue_number
    }

    await fetchQueueData()
  } catch (error) {
    console.error('Gagal memanggil antrian:', error)
  }
}

onMounted(fetchQueueData)
</script>

<template>
  <Breadcrumb title="List Antrian" :items="breadcrumbItems" />

  <!-- Tampilan Nomor Antrian -->
  <div class="row mb-4">
    <div class="col">
      <div class="card rounded-lg">
        <div class="card-header bg-primary rounded-t-lg">
          <h5 class="text-white">Nomor Antrian Saat Ini</h5>
        </div>
        <div class="card-body text-center">
          <!-- Untuk superadmin: tampilkan A dan B -->
          <template v-if="userRole === 'superadmin'">
            <div class="flex justify-center gap-8">
              <div class="text-center">
                <h4 class="font-semibold">Kategori A</h4>
                <p class="text-2xl">{{ currentQueueA ? currentQueueA.Number : '-' }}</p>
              </div>
              <div class="text-center">
                <h4 class="font-semibold">Kategori B</h4>
                <p class="text-2xl">{{ currentQueueB ? currentQueueB.Number : '-' }}</p>
              </div>
            </div>
          </template>

          <!-- Untuk admin biasa -->
          <template v-else>
            <h3 class="font-weight-bold mb-3">
              {{ currentQueue ? currentQueue.Number : '-' }}
            </h3>
            <button class="btn btn-primary btn-lg" @click="nextQueue">
              Next <i class="ti ti-arrow-right"></i>
            </button>
          </template>
        </div>
      </div>
    </div>
  </div>

  <!-- Tabel Antrian -->
  <div class="card">
    <div class="card-header d-flex justify-content-between align-items-center">
      <div>
        <h5>Tabel Antrian</h5>
        <span class="d-block m-t-5">Daftar antrian berdasarkan kategori</span>
      </div>
      <select v-model="selectedLoket" class="form-control w-auto">
        <option v-for="option in loketOptions" :key="option" :value="option">
          {{ option === 'Semua' ? 'Semua Loket' : `Loket ${option}` }}
        </option>
      </select>
    </div>

    <div class="card-body table-border-style">
      <div class="table-responsive">
        <table class="table">
          <thead>
            <tr>
              <th>Nomor Antrian</th>
              <th>Loket</th>
              <th>Waktu Daftar</th>
              <th>Status</th>
              <th>Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="queue in filteredQueue" :key="queue.ID">
              <td>{{ queue.Number }}</td>
              <td>{{ queue.Category }}</td>
              <td>{{ new Date(queue.CreatedAt).toLocaleString() }}</td>
              <td>
                <span
                  class="text-white badge"
                  :class="queue.Status === 'called' ? 'badge bg-success' : 'badge bg-warning'"
                >
                  {{ queue.Status === 'waiting' ? 'Menunggu' : 'Aktif' }}
                </span>
              </td>
              <td>
                <!-- Tombol hanya untuk admin, bukan superadmin -->
                <button
                  v-if="queue.Status === 'waiting' && userRole !== 'superadmin'"
                  @click="nextQueue"
                  class="btn btn-sm btn-primary"
                >
                  Panggil
                </button>
              </td>
            </tr>
            <tr v-if="filteredQueue.length === 0">
              <td colspan="5" class="text-center">Tidak ada data antrian</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
