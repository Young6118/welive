import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { login as loginApi, getUserInfo } from '@/api/user'
import type { User, LoginRequest } from '@/types'

export const useUserStore = defineStore('user', () => {
  // State
  const token = ref<string>(localStorage.getItem('token') || '')
  const userInfo = ref<User | null>(null)
  const isLoading = ref(false)

  // Getters
  const isLoggedIn = computed(() => !!token.value)
  const userId = computed(() => userInfo.value?.id)

  // Actions
  const setToken = (newToken: string): void => {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  const clearToken = (): void => {
    token.value = ''
    userInfo.value = null
    localStorage.removeItem('token')
  }

  const login = async (credentials: LoginRequest): Promise<boolean> => {
    isLoading.value = true
    try {
      const res = await loginApi(credentials)
      setToken(res.token)
      await fetchUserInfo()
      return true
    } catch (error) {
      return false
    } finally {
      isLoading.value = false
    }
  }

  const fetchUserInfo = async (): Promise<User | null> => {
    try {
      const res = await getUserInfo()
      userInfo.value = res
      return res
    } catch (error) {
      return null
    }
  }

  const logout = (): void => {
    clearToken()
  }

  return {
    token,
    userInfo,
    isLoading,
    isLoggedIn,
    userId,
    login,
    logout,
    fetchUserInfo,
    setToken,
    clearToken
  }
})
