<script setup>
import Logo from '@/components/Logo.vue'
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'

const formData = reactive({
  username: '',
  password: '',
})

const loading = ref(false)
const error = ref('')
const router = useRouter()

const login = async () => {
  error.value = ''

  if (!formData.username || !formData.password) {
    error.value = 'Username dan password diperlukan'
    return
  }

  try {
    loading.value = true

    const response = await fetch('http://localhost:8800/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username: formData.username,
        password: formData.password,
      }),
    })

    const data = await response.json()
    if (!response.ok) throw new Error(data.error || 'Login gagal')

    const { user, token } = data

    // Simpan ke sessionStorage (bukan localStorage)
    sessionStorage.setItem('auth', JSON.stringify({ user, token }))
    console.log('Login berhasil:', data)

    router.push({ name: 'dashboard' })
  } catch (err) {
    error.value = err.message || 'Login gagal. Silakan coba lagi.'
    console.error('Login error:', err)
  } finally {
    loading.value = false
  }
}

const clearErrorOnInput = () => {
  if (error.value) error.value = ''
}
</script>

<template>
  <div class="auth-main relative">
    <div class="auth-wrapper v1 flex items-center w-full h-full min-h-screen">
      <div
        class="flex items-center justify-center grow flex-col min-h-screen bg-cover relative p-6 bg-gradient-to-b from-blue-500 to-blue-400"
      >
        <div class="card sm:my-12 w-full max-w-[480px] shadow-none">
          <div class="card-body !p-10">
            <Logo class="mb-7" size="10" />
            <!-- Error Message -->
            <div v-if="error" class="mb-4 p-3 bg-red-100 text-red-700 rounded">
              {{ error }}
            </div>
            <div class="mb-3">
              <input
                type="text"
                v-model="formData.username"
                class="form-control"
                id="usernameInput"
                placeholder="Username"
                @input="clearErrorOnInput"
              />
            </div>
            <div class="mb-4">
              <input
                type="password"
                v-model="formData.password"
                class="form-control"
                id="passwordInput"
                placeholder="Password"
                @input="clearErrorOnInput"
              />
            </div>
            <div class="mt-4">
              <button
                type="button"
                class="btn btn-primary w-full"
                @click="login"
                :disabled="loading"
              >
                {{ loading ? 'Loading...' : 'Login' }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
