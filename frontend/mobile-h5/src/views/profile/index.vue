<template>
  <div class="profile-page">

    <div class="content-wrapper">
      <!-- 用户信息卡片 -->
      <div class="user-card card">
        <div class="user-info">
          <div class="user-avatar">
            <van-icon name="user-circle-o" size="64" />
          </div>
          <div class="user-detail">
            <div class="user-name">{{ userInfo.username || '未登录' }}</div>
            <div class="user-bio">{{ userInfo.bio || '点击编辑个人简介' }}</div>
          </div>
        </div>
        <div class="user-stats">
          <div class="stat-item">
            <div class="stat-value">128</div>
            <div class="stat-label">关注</div>
          </div>
          <div class="stat-item">
            <div class="stat-value">256</div>
            <div class="stat-label">粉丝</div>
          </div>
          <div class="stat-item">
            <div class="stat-value">32</div>
            <div class="stat-label">获赞</div>
          </div>
        </div>
      </div>

      <!-- 功能菜单 -->
      <div class="menu-list">
        <van-cell-group>
          <van-cell title="我的问题" is-link @click="goToMyQuestions">
            <template #icon>
              <van-icon name="question-o" class="menu-icon" />
            </template>
          </van-cell>
          <van-cell title="我的笔记" is-link @click="goToMyNotes">
            <template #icon>
              <van-icon name="notes-o" class="menu-icon" />
            </template>
          </van-cell>
          <van-cell title="我的收藏" is-link @click="goToMyFavorites">
            <template #icon>
              <van-icon name="star-o" class="menu-icon" />
            </template>
          </van-cell>
          <van-cell title="我的村落" is-link @click="goToMyVillages">
            <template #icon>
              <van-icon name="cluster-o" class="menu-icon" />
            </template>
          </van-cell>
        </van-cell-group>

        <van-cell-group class="menu-group">
          <van-cell title="设置" is-link @click="goToSettings">
            <template #icon>
              <van-icon name="setting-o" class="menu-icon" />
            </template>
          </van-cell>
          <van-cell title="关于" is-link>
            <template #icon>
              <van-icon name="info-o" class="menu-icon" />
            </template>
          </van-cell>
        </van-cell-group>

        <van-cell-group class="menu-group">
          <van-cell title="退出登录" @click="onLogout" class="logout-cell">
            <template #icon>
              <van-icon name="close" class="menu-icon" color="#ee0a24" />
            </template>
          </van-cell>
        </van-cell-group>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { showConfirmDialog, showToast } from 'vant'
import { useUserStore } from '@/stores/user'
import { getUserInfo } from '@/api/user'

const router = useRouter()
const userStore = useUserStore()
const userInfo = ref({})

const loadUserInfo = async () => {
  try {
    const res = await getUserInfo()
    userInfo.value = res
  } catch (error) {
    console.error(error)
  }
}

const goToMyQuestions = () => {
  // TODO: 我的问题
}

const goToMyNotes = () => {
  // TODO: 我的笔记
}

const goToMyFavorites = () => {
  // TODO: 我的收藏
}

const goToMyVillages = () => {
  // TODO: 我的村落
}

const goToSettings = () => {
  router.push('/settings')
}

const onLogout = () => {
  showConfirmDialog({
    title: '确认退出',
    message: '确定要退出登录吗？',
  })
    .then(() => {
      userStore.logout()
      showToast('已退出登录')
      router.replace('/login')
    })
    .catch(() => {})
}

onMounted(() => {
  loadUserInfo()
})
</script>

<style lang="scss" scoped>
.profile-page {
  min-height: 100vh;

  .user-card {
    .user-info {
      display: flex;
      align-items: center;
      gap: 16px;
      margin-bottom: 20px;

      .user-avatar {
        color: #667eea;
      }

      .user-detail {
        .user-name {
          font-size: 18px;
          font-weight: 600;
          color: #333;
          margin-bottom: 4px;
        }

        .user-bio {
          font-size: 13px;
          color: #999;
        }
      }
    }

    .user-stats {
      display: flex;
      justify-content: space-around;
      padding-top: 16px;
      border-top: 1px solid #f0f0f0;

      .stat-item {
        text-align: center;

        .stat-value {
          font-size: 18px;
          font-weight: 600;
          color: #333;
          margin-bottom: 4px;
        }

        .stat-label {
          font-size: 12px;
          color: #999;
        }
      }
    }
  }

  .menu-list {
    margin-top: 12px;

    .menu-group {
      margin-top: 12px;
    }

    .menu-icon {
      margin-right: 8px;
      font-size: 20px;
      color: #667eea;
    }

    .logout-cell {
      :deep(.van-cell__title) {
        color: #ee0a24;
      }
    }
  }
}
</style>
