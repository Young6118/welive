<template>
  <div class="settings-page">
    <van-nav-bar
      :title="$t('settings.title')"
      left-arrow
      @click-left="onClickLeft"
      fixed
    />

    <div class="settings-content">
      <!-- 主题设置 -->
      <van-cell-group :title="$t('settings.appearance')">
        <van-cell :title="$t('settings.theme')" :value="themeText" is-link @click="showThemePicker = true">
          <template #icon>
            <van-icon :name="themeIcon" class="setting-icon" />
          </template>
        </van-cell>
      </van-cell-group>

      <!-- 语言设置 -->
      <van-cell-group :title="$t('settings.language')">
        <van-cell :title="$t('settings.language')" :value="localeText" is-link @click="showLanguagePicker = true">
          <template #icon>
            <van-icon name="guide-o" class="setting-icon" />
          </template>
        </van-cell>
      </van-cell-group>

      <!-- 其他设置 -->
      <van-cell-group :title="$t('settings.other')">
        <van-cell :title="$t('settings.clearCache')" is-link @click="clearCache">
          <template #icon>
            <van-icon name="delete-o" class="setting-icon" />
          </template>
        </van-cell>
        <van-cell :title="$t('settings.aboutUs')" is-link @click="showAbout">
          <template #icon>
            <van-icon name="info-o" class="setting-icon" />
          </template>
        </van-cell>
      </van-cell-group>

      <!-- 版本信息 -->
      <div class="version-info">
        {{ $t('settings.version') }} 1.0.0
      </div>
    </div>

    <!-- 主题选择弹窗 -->
    <van-popup v-model:show="showThemePicker" round position="bottom">
      <van-picker
        :columns="themeColumns"
        @confirm="onThemeConfirm"
        @cancel="showThemePicker = false"
        :default-index="themeIndex"
      />
    </van-popup>

    <!-- 语言选择弹窗 -->
    <van-popup v-model:show="showLanguagePicker" round position="bottom">
      <van-picker
        :columns="languageColumns"
        @confirm="onLanguageConfirm"
        @cancel="showLanguagePicker = false"
        :default-index="languageIndex"
      />
    </van-popup>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { showToast, showDialog } from 'vant'
import { useThemeStore, type Theme } from '@/stores/theme'
import { useLocaleStore } from '@/stores/locale'

const router = useRouter()
const { t, locale } = useI18n()
const themeStore = useThemeStore()
const localeStore = useLocaleStore()

const showThemePicker = ref(false)
const showLanguagePicker = ref(false)

// 主题选项
const themeColumns = [
  { text: t('settings.themeLight'), value: 'light' },
  { text: t('settings.themeDark'), value: 'dark' },
  { text: t('settings.themeAuto'), value: 'auto' },
]

// 语言选项
const languageColumns = [
  { text: t('settings.languageZh'), value: 'zh-CN' },
  { text: t('settings.languageEn'), value: 'en-US' },
]

// 当前主题文本
const themeText = computed(() => {
  const map: Record<Theme, string> = {
    light: t('settings.themeLight'),
    dark: t('settings.themeDark'),
    auto: t('settings.themeAuto'),
  }
  return map[themeStore.theme]
})

// 当前主题图标
const themeIcon = computed(() => {
  const map: Record<Theme, string> = {
    light: 'sun-o',
    dark: 'moon-o',
    auto: 'clock-o',
  }
  return map[themeStore.theme]
})

// 当前主题索引
const themeIndex = computed(() => {
  return themeColumns.findIndex(item => item.value === themeStore.theme)
})

// 当前语言文本
const localeText = computed(() => {
  const item = languageColumns.find(item => item.value === localeStore.locale)
  return item?.text || t('settings.languageZh')
})

// 当前语言索引
const languageIndex = computed(() => {
  return languageColumns.findIndex(item => item.value === localeStore.locale)
})

// 返回上一页
const onClickLeft = () => {
  router.back()
}

// 确认主题选择
const onThemeConfirm = ({ selectedOptions }: { selectedOptions: Array<{ value: Theme }> }) => {
  themeStore.setTheme(selectedOptions[0].value)
  showThemePicker.value = false
  showToast(t('settings.themeUpdated'))
}

// 确认语言选择
const onLanguageConfirm = ({ selectedOptions }: { selectedOptions: Array<{ value: string }> }) => {
  const newLocale = selectedOptions[0].value
  localeStore.setLocale(newLocale as 'zh-CN' | 'en-US')
  locale.value = newLocale
  showLanguagePicker.value = false
  showToast(t('settings.languageUpdated'))
}

// 清除缓存
const clearCache = () => {
  showDialog({
    title: t('settings.clearCache'),
    message: '确定要清除所有缓存数据吗？',
    showCancelButton: true,
  }).then(() => {
    localStorage.removeItem('user')
    showToast(t('settings.cacheCleared'))
  }).catch(() => {})
}

// 关于我们
const showAbout = () => {
  showDialog({
    title: t('settings.aboutUs'),
    message: '蛋蛋 - 智能问答社区\n\n版本：1.0.0\n\n一个基于AI的智能问答社区，让知识分享更简单。',
    confirmButtonText: t('common.confirm'),
  })
}
</script>

<style lang="scss" scoped>
.settings-page {
  min-height: 100vh;
  background: #f5f5f5;

  :deep(.van-nav-bar) {
    background: rgba(255, 255, 255, 0.72);
    backdrop-filter: blur(20px) saturate(180%);
    -webkit-backdrop-filter: blur(20px) saturate(180%);
  }

  .settings-content {
    padding-top: 46px;

    :deep(.van-cell-group__title) {
      padding: 16px 16px 8px;
      color: #999;
      font-size: 13px;
    }

    .setting-icon {
      margin-right: 8px;
      font-size: 20px;
      color: #667eea;
    }
  }

  .version-info {
    text-align: center;
    padding: 24px;
    color: #999;
    font-size: 13px;
  }
}

// 暗黑模式适配
:global(.dark) {
  .settings-page {
    background: #1a1a1a;

    :deep(.van-nav-bar) {
      background: rgba(30, 30, 30, 0.72);
      border-bottom: 0.5px solid rgba(255, 255, 255, 0.08);

      .van-nav-bar__title,
      .van-nav-bar__arrow {
        color: #fff;
      }
    }

    :deep(.van-cell-group) {
      background: #2a2a2a;
    }

    :deep(.van-cell) {
      background: #2a2a2a;
      color: #fff;

      &::after {
        border-color: rgba(255, 255, 255, 0.08);
      }
    }

    :deep(.van-cell__value) {
      color: #999;
    }

    :deep(.van-cell-group__title) {
      color: #666;
    }
  }
}
</style>
