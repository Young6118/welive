import { defineStore } from 'pinia'
import { ref, watch } from 'vue'
import { useI18n } from 'vue-i18n'

export type Locale = 'zh-CN' | 'en-US'

export const useLocaleStore = defineStore('locale', () => {
  // 从localStorage读取语言设置，默认中文
  const savedLocale = localStorage.getItem('locale') as Locale
  const locale = ref<Locale>(savedLocale || 'zh-CN')

  // 设置语言
  const setLocale = (newLocale: Locale) => {
    locale.value = newLocale
    localStorage.setItem('locale', newLocale)
    document.documentElement.lang = newLocale

    // 获取i18n实例并切换语言
    const i18n = (window as any).__VUE_I18N__
    if (i18n) {
      i18n.global.locale.value = newLocale
    }
  }

  // 初始化
  document.documentElement.lang = locale.value

  // 监听语言变化
  watch(locale, (newVal) => {
    document.documentElement.lang = newVal
  })

  return {
    locale,
    setLocale,
  }
})
