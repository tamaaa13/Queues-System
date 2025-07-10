import '@/assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import Toast from 'vue-toastification'

import 'vue-toastification/dist/index.css'

import '@/assets/js/index.js'

const toastOptions = {
  position: 'top-center',
}

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Toast, toastOptions)
app.mount('#app')
