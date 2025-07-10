<template>
  <Breadcrumb title="Pengguna" :items="breadcrumbItems" />

  <div class="card p-4 bg-white rounded shadow">
    <div class="flex justify-between items-center mb-4">
      <div>
        <h5 class="text-xl font-semibold">Daftar Pengguna</h5>
        <span class="text-sm text-gray-600">Kelola user dan hak akses pengguna sistem</span>
      </div>
      <div class="flex gap-2">
        <input
          type="text"
          placeholder="Cari pengguna..."
          v-model="searchQuery"
          class="border border-gray-300 rounded px-3 py-1 focus:outline-none focus:ring-2 focus:ring-indigo-500"
        />
        <button
          class="bg-indigo-600 text-white px-3 py-1 rounded hover:bg-indigo-700 flex items-center gap-1"
          @click="openModal"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 4v16m8-8H4"
            />
          </svg>
          Tambah User
        </button>
      </div>
    </div>

    <table class="w-full text-left border-collapse">
      <thead>
        <tr class="border-b border-gray-300">
          <th class="p-2">No</th>
          <th class="p-2">Username</th>
          <th class="p-2">Hak Akses</th>
          <th class="p-2">Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(user, index) in filteredUsers"
          :key="user.id"
          class="border-b border-gray-200 hover:bg-gray-50"
        >
          <td class="p-2">{{ index + 1 }}</td>
          <td class="p-2">{{ user.username }}</td>
          <td class="p-2">
            <span class="inline-block bg-green-500 text-white px-2 py-1 rounded text-xs">
              {{ user.category?.Valid ? user.category.String : user.role }}
            </span>
          </td>
          <td class="p-2 flex gap-2">
            <button
              class="bg-yellow-400 hover:bg-yellow-500 text-black px-2 rounded"
              @click="openEditModal(user)"
              title="Edit"
            >
              Edit
            </button>
            <button
              class="bg-red-500 hover:bg-red-600 text-white px-2 rounded"
              @click="deleteUser(user.id)"
              title="Hapus"
            >
              Delete
            </button>
          </td>
        </tr>
        <tr v-if="filteredUsers.length === 0">
          <td colspan="4" class="p-4 text-center text-gray-500">Tidak ada data</td>
        </tr>
      </tbody>
    </table>
  </div>

  <!-- Modal -->
  <div
    v-if="isModalOpen"
    class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
  >
    <div class="bg-white rounded-lg shadow-lg w-full max-w-md p-6 relative">
      <h3 class="text-lg font-semibold mb-4">
        {{ formMode === 'create' ? 'Tambah' : 'Edit' }} User
      </h3>
      <form @submit.prevent="saveUser" class="space-y-4">
        <div>
          <label for="username" class="block mb-1 font-medium">Username</label>
          <input
            id="username"
            type="text"
            v-model="formData.username"
            required
            class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
          />
        </div>
        <div>
          <label for="password" class="block mb-1 font-medium">Password</label>
          <input
            id="password"
            type="password"
            v-model="formData.password"
            :required="formMode === 'create'"
            placeholder="Isi password baru jika ingin mengganti"
            class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
          />
        </div>
        <div>
          <label for="role" class="block mb-1 font-medium">Loket</label>
          <select
            id="role"
            v-model="formData.role"
            class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
          >
            <option v-for="role in roles" :key="role" :value="role">{{ role }}</option>
          </select>
        </div>
        <div class="flex justify-end gap-2 pt-4">
          <button
            type="button"
            @click="closeModal"
            class="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400"
          >
            Batal
          </button>
          <button
            type="submit"
            class="px-4 py-2 bg-indigo-600 text-white rounded hover:bg-indigo-700"
          >
            {{ formMode === 'create' ? 'Simpan' : 'Update' }}
          </button>
        </div>
      </form>
      <button
        @click="closeModal"
        class="absolute top-2 right-2 text-gray-500 hover:text-gray-700"
        aria-label="Close modal"
      >
        &times;
      </button>
    </div>
  </div>
</template>

<script setup>
import Breadcrumb from '@/components/Breadcrumb.vue'
import { ref, reactive, computed, onMounted } from 'vue'

const getToken = () => {
  const auth = sessionStorage.getItem('auth')
  if (!auth) return null
  try {
    return JSON.parse(auth).token
  } catch {
    return null
  }
}

const users = ref([])

const formMode = ref('create')
const formData = reactive({
  id: null,
  username: '',
  oldUsername: '', // tambahan: simpan username lama untuk update
  password: '',
  role: 'A',
})

const roles = ['A', 'B', 'superadmin']

const searchQuery = ref('')

const isModalOpen = ref(false)

const filteredUsers = computed(() => {
  if (!searchQuery.value) return users.value
  return users.value.filter((user) => {
    const categoryString = user.category?.Valid ? user.category.String : ''
    return (
      user.username.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      user.role.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      categoryString.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  })
})

const resetForm = () => {
  formData.id = null
  formData.username = ''
  formData.oldUsername = ''
  formData.password = ''
  formData.role = 'A'
}

const openModal = () => {
  formMode.value = 'create'
  resetForm()
  isModalOpen.value = true
}

const openEditModal = (user) => {
  formMode.value = 'edit'
  formData.id = user.id
  formData.username = user.username
  formData.oldUsername = user.username // simpan username lama
  formData.password = ''
  if (user.category?.Valid && roles.includes(user.category.String)) {
    formData.role = user.category.String
  } else if (roles.includes(user.role)) {
    formData.role = user.role
  } else {
    formData.role = 'A'
  }
  isModalOpen.value = true
}

const closeModal = () => {
  isModalOpen.value = false
  resetForm()
}

const fetchUsers = async () => {
  const token = getToken()
  if (!token) {
    alert('User belum login')
    return
  }
  try {
    const res = await fetch('http://localhost:8800/admin/all-admin', {
      headers: { Authorization: `Bearer ${token}` },
    })
    if (!res.ok) throw new Error(`HTTP error ${res.status}`)
    users.value = await res.json()
  } catch (e) {
    alert('Gagal mengambil data admin: ' + e.message)
  }
}

const saveUser = async () => {
  const token = getToken()
  if (!token) {
    alert('User belum login')
    return
  }
  if (!formData.username) {
    alert('Username wajib diisi')
    return
  }
  if (formMode.value === 'create' && !formData.password) {
    alert('Password wajib diisi saat tambah user')
    return
  }

  try {
    let role = ''
    let category = null

    if (formData.role === 'A' || formData.role === 'B') {
      role = 'admin'
      category = formData.role
    } else if (formData.role === 'superadmin') {
      role = 'superadmin'
      category = ''
    } else {
      role = 'admin'
      category = 'A'
    }

    let payload = {
      username: formData.oldUsername,
      role,
      category,
    }

    if (formMode.value === 'create') {
      payload = {
        username: formData.username,
        password: formData.password,
        role,
        category,
      }
    } else {
      // Hanya sertakan jika ada perubahan
      if (formData.username !== formData.oldUsername) {
        payload.new_username = formData.username
      }
      if (formData.password.trim() !== '') {
        payload.password = formData.password
      }
    }

    const url =
      formMode.value === 'create'
        ? 'http://localhost:8800/admin/add-admin'
        : 'http://localhost:8800/admin/update-admin'

    const method = formMode.value === 'create' ? 'POST' : 'PUT'

    const response = await fetch(url, {
      method,
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
    })

    if (!response.ok) {
      const text = await response.text()
      alert(`Gagal menyimpan user: HTTP error ${response.status}: ${text}`)
      return
    }

    const savedUser = await response.json()

    if (formMode.value === 'create') {
      users.value.push(savedUser)
    } else {
      const idx = users.value.findIndex((u) => u.id === formData.id)
      if (idx !== -1) {
        users.value[idx] = {
          ...users.value[idx],
          ...savedUser,
        }
      }
      // atau kalau mau lebih aman:
      // await fetchUsers()
    }

    closeModal()
  } catch (e) {
    alert('Gagal menyimpan user: ' + e.message)
  }
}

const deleteUser = async (id) => {
  if (!confirm('Apakah anda yakin ingin menghapus user ini?')) return

  const token = getToken()
  if (!token) {
    alert('User belum login')
    return
  }

  try {
    const res = await fetch(`http://localhost:8800/admin/delete-admin/${id}`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })

    if (!res.ok) {
      const text = await res.text()
      throw new Error(`HTTP error ${res.status}: ${text}`)
    }

    alert('User berhasil dihapus')
    // Refresh user list dari server
    await fetchUsers()
  } catch (e) {
    alert('Gagal menghapus user: ' + e.message)
  }
}

const breadcrumbItems = [{ text: 'Home', link: '/dashboard/beranda' }, { text: 'Pengguna' }]

onMounted(() => {
  fetchUsers()
})
</script>
