import { defineStore } from 'pinia'
import { ref, watch } from 'vue'

export type Theme = 'light' | 'dark' | 'auto'

export const useThemeStore = defineStore('theme', () => {
  // 从localStorage读取主题设置，默认auto
  const savedTheme = localStorage.getItem('theme') as Theme
  const theme = ref<Theme>(savedTheme || 'auto')

  // 判断是否暗黑模式
  const isDark = ref(false)

  // 更新暗黑模式状态
  const updateDarkMode = () => {
    if (theme.value === 'auto') {
      // 根据系统偏好
      isDark.value = window.matchMedia('(prefers-color-scheme: dark)').matches
    } else {
      isDark.value = theme.value === 'dark'
    }

    // 应用主题到DOM - 通过添加/移除dark类来切换主题
    // CSS变量定义在theme.css中
    if (isDark.value) {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
  }

  // 设置主题
  const setTheme = (newTheme: Theme) => {
    theme.value = newTheme
    localStorage.setItem('theme', newTheme)
    updateDarkMode()
  }

  // 监听系统主题变化
  const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
  mediaQuery.addEventListener('change', () => {
    if (theme.value === 'auto') {
      updateDarkMode()
    }
  })

  // 初始化
  updateDarkMode()

  // 监听主题变化
  watch(theme, updateDarkMode)

  return {
    theme,
    isDark,
    setTheme,
  }
})
