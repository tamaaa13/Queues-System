<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const user = ref(null)
const isDropdownOpen = ref(false)
const dropdownWrapperRef = ref(null)

const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value
}

const logout = () => {
  sessionStorage.removeItem('auth') // hanya hapus sesi saat ini
  router.push('/login')
}

const handleClickOutside = (event) => {
  if (dropdownWrapperRef.value && !dropdownWrapperRef.value.contains(event.target)) {
    isDropdownOpen.value = false
  }
}

onMounted(() => {
  const storedAuth = sessionStorage.getItem('auth')
  if (storedAuth) {
    const parsed = JSON.parse(storedAuth)
    user.value = parsed.user
  }
  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<template>
  <header class="pc-header">
    <div class="header-wrapper flex max-sm:px-[15px] px-[25px] grow">
      <div class="me-auto pc-mob-drp">
        <!-- Sidebar menu... -->
      </div>

      <div class="ms-auto">
        <ul class="inline-flex *:min-h-header-height *:inline-flex *:items-center">
          <li ref="dropdownWrapperRef" class="pc-h-item header-user-profile relative">
            <a
              @click.prevent="toggleDropdown"
              class="pc-head-link me-0"
              href="#"
              role="button"
              :aria-expanded="isDropdownOpen.toString()"
            >
              <img
                src="@/assets/images/user/avatar-2.jpg"
                alt="user-image"
                class="user-avtar w-10 h-10 rounded-full"
              />
            </a>

            <div
              class="dropdown-menu dropdown-user-profile dropdown-menu-end pc-h-dropdown p-2 absolute right-0 z-50 mt-2 bg-white border rounded shadow"
              v-show="isDropdownOpen"
            >
              <div class="dropdown-header flex items-center justify-between py-4 px-5">
                <h5 class="m-0">Profile</h5>
              </div>
              <div class="dropdown-body py-4 px-5">
                <div class="profile-notification-scroll" style="max-height: calc(100vh - 225px)">
                  <div class="flex mb-1 items-center">
                    <div class="shrink-0">
                      <img
                        src="@/assets/images/user/avatar-2.jpg"
                        alt="user-image"
                        class="w-10 rounded-full"
                      />
                    </div>
                    <div class="grow ms-3">
                      <h6 class="mb-1">{{ user?.username || 'User' }}</h6>
                    </div>
                  </div>
                  <hr class="border-secondary-500/10 my-4" />
                  <a href="#" class="dropdown-item" @click.prevent="logout">
                    <span>
                      <i class="ti ti-logout me-2"></i>
                      <span>Logout</span>
                    </span>
                  </a>
                </div>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </header>
</template>
