<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import { useUserStore } from '@/stores/user'
import type { LoginRequest } from '@/types'

const router = useRouter()
const userStore = useUserStore()
const loading = ref(false)

const form = reactive<LoginRequest>({
  username: '',
  password: ''
})

const onSubmit = async (): Promise<void> => {
  loading.value = true
  try {
    const success = await userStore.login(form)
    if (success) {
      showToast('登录成功')
      router.replace('/home')
    } else {
      showToast('登录失败')
    }
  } catch (error) {
    showToast('登录失败')
  } finally {
    loading.value = false
  }
}

const goToRegister = (): void => {
  router.push('/register')
}
</script>

<template>
  <div class="login-page">
    <div class="login-header">
      <h1 class="app-name">蛋蛋</h1>
      <p class="app-slogan">AI时代的社交平台</p>
    </div>

    <div class="login-form card">
      <van-form @submit="onSubmit">
        <van-field
          v-model="form.username"
          name="username"
          label="用户名"
          placeholder="请输入用户名"
          :rules="[{ required: true, message: '请填写用户名' }]"
        />
        <van-field
          v-model="form.password"
          type="password"
          name="password"
          label="密码"
          placeholder="请输入密码"
          :rules="[{ required: true, message: '请填写密码' }]"
        />
        <div class="submit-btn">
          <van-button round block type="primary" native-type="submit" :loading="loading">
            登录
          </van-button>
        </div>
      </van-form>

      <div class="other-options">
        <span class="link" @click="goToRegister">注册账号</span>
        <span class="divider">|</span>
        <span class="link">忘记密码</span>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.login-page {
  min-height: 100vh;
  padding: 60px 20px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);

  .login-header {
    margin-bottom: 40px;
    color: #fff;
    text-align: center;

    .app-name {
      margin-bottom: 8px;
      font-size: 36px;
      font-weight: bold;
    }

    .app-slogan {
      font-size: 14px;
      opacity: 0.9;
    }
  }

  .login-form {
    padding: 30px 20px;
    background: #fff;
    border-radius: 16px;

    .submit-btn {
      margin-top: 24px;
    }

    .other-options {
      margin-top: 20px;
      font-size: 14px;
      color: #666;
      text-align: center;

      .link {
        color: #667eea;
        cursor: pointer;
      }

      .divider {
        margin: 0 16px;
        color: #ddd;
      }
    }
  }
}
</style>
