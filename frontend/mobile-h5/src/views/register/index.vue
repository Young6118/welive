<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import { register } from '@/api/user'
import type { RegisterRequest } from '@/types'

const router = useRouter()
const loading = ref(false)
const showCategoryPicker = ref(false)

const form = reactive<RegisterRequest & { confirmPassword: string }>({
  username: '',
  password: '',
  confirmPassword: '',
  email: ''
})

const validateConfirmPassword = (value: string): boolean => {
  return value === form.password
}

const onSubmit = async (): Promise<void> => {
  loading.value = true
  try {
    await register({
      username: form.username,
      password: form.password,
      email: form.email
    })
    showToast('注册成功')
    router.replace('/login')
  } catch (error) {
    showToast('注册失败')
  } finally {
    loading.value = false
  }
}

const goBack = (): void => {
  router.back()
}

const goToLogin = (): void => {
  router.replace('/login')
}
</script>

<template>
  <div class="register-page">
    <van-nav-bar title="注册" left-arrow @click-left="goBack" />

    <div class="content-wrapper">
      <div class="register-form card">
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
          <van-field
            v-model="form.confirmPassword"
            type="password"
            name="confirmPassword"
            label="确认密码"
            placeholder="请再次输入密码"
            :rules="[
              { required: true, message: '请确认密码' },
              { validator: validateConfirmPassword, message: '两次密码不一致' }
            ]"
          />
          <van-field
            v-model="form.email"
            name="email"
            label="邮箱"
            placeholder="请输入邮箱（选填）"
          />
          <div class="submit-btn">
            <van-button round block type="primary" native-type="submit" :loading="loading">
              注册
            </van-button>
          </div>
        </van-form>

        <div class="login-link">
          已有账号？<span class="link" @click="goToLogin">立即登录</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.register-page {
  min-height: 100vh;
  background-color: #f5f5f5;

  .register-form {
    margin-top: 20px;

    .submit-btn {
      margin-top: 24px;
    }

    .login-link {
      margin-top: 20px;
      font-size: 14px;
      color: #666;
      text-align: center;

      .link {
        color: #667eea;
        cursor: pointer;
      }
    }
  }
}
</style>
