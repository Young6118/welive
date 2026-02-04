<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useThemeStore } from '@/stores/theme'

const cachedViews = ref<string[]>(['Home', 'Notes', 'Chat', 'Discover', 'Profile'])
const themeStore = useThemeStore()

onMounted(() => {
  // 初始化主题
  themeStore.setTheme(themeStore.theme)
})
</script>

<template>
  <router-view v-slot="{ Component }">
    <keep-alive :include="cachedViews">
      <component :is="Component" />
    </keep-alive>
  </router-view>
</template>

<style lang="scss">
#app {
  min-height: 100vh;
  background-color: var(--body-background);
  transition: background-color 0.3s ease;
}
</style>
