<script setup lang="ts">
import { ref } from 'vue'

const active = ref(0)
</script>

<template>
  <div class="main-layout">
    <div class="page-content">
      <router-view />
    </div>
    <van-tabbar
      v-model="active"
      route
      class="main-tabbar"
      :fixed="false"
    >
      <van-tabbar-item to="/home" icon="home-o">{{ $t('nav.home') }}</van-tabbar-item>
      <van-tabbar-item to="/notes" icon="notes-o">{{ $t('nav.notes') }}</van-tabbar-item>
      <van-tabbar-item to="/chat" icon="chat-o">{{ $t('nav.chat') }}</van-tabbar-item>
      <van-tabbar-item to="/discover" icon="compass-o">{{ $t('nav.discover') }}</van-tabbar-item>
      <van-tabbar-item to="/profile" icon="user-o">{{ $t('nav.profile') }}</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<style lang="scss" scoped>
.main-layout {
  min-height: 100vh;
  padding-bottom: calc(50px + env(safe-area-inset-bottom));

  .page-content {
    min-height: calc(100vh - 50px - env(safe-area-inset-bottom));
  }

  .main-tabbar {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: 100;
    padding-bottom: env(safe-area-inset-bottom);
    background: rgba(255, 255, 255, 0.72) !important;
    backdrop-filter: blur(20px) saturate(180%);
    -webkit-backdrop-filter: blur(20px) saturate(180%);
    border-top: 0.5px solid rgba(0, 0, 0, 0.08);

    :deep(.van-tabbar-item) {
      background: transparent !important;

      .van-tabbar-item__icon {
        font-size: 22px;
      }
    }

    :deep(.van-tabbar-item--active) {
      background: transparent !important;
    }
  }
}
</style>
