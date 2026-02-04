<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import type { User } from '@/types'

const route = useRoute()
const router = useRouter()
const userId = route.params.id

const user = ref<User>({
  id: Number(userId),
  username: '用户' + userId,
  bio: '这个人很懒，什么都没有写~'
})

const goBack = (): void => {
  router.back()
}

onMounted(() => {
  // TODO: 加载用户信息
})
</script>

<template>
  <div class="user-profile-page">
    <van-nav-bar title="用户主页" left-arrow @click-left="goBack" />

    <div class="content-wrapper">
      <div class="user-card card">
        <div class="user-info">
          <van-icon name="user-circle-o" size="64" />
          <div class="user-detail">
            <div class="user-name">{{ user.username }}</div>
            <div class="user-bio">{{ user.bio }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.user-profile-page {
  min-height: 100vh;
  background-color: #f5f5f5;

  .user-card {
    .user-info {
      display: flex;
      align-items: center;
      gap: 16px;

      .user-detail {
        .user-name {
          margin-bottom: 4px;
          font-size: 18px;
          font-weight: 600;
          color: #333;
        }

        .user-bio {
          font-size: 13px;
          color: #999;
        }
      }
    }
  }
}
</style>
