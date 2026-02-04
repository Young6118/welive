<template>
  <div class="create-question-page">
    <van-nav-bar
      title="发布问题"
      left-arrow
      @click-left="goBack"
      right-text="发布"
      @click-right="onSubmit"
    />
    
    <div class="content-wrapper">
      <van-field
        v-model="form.title"
        label="标题"
        placeholder="请输入问题标题"
        :rules="[{ required: true }]"
      />
      <van-field
        v-model="form.content"
        rows="6"
        autosize
        label="内容"
        type="textarea"
        maxlength="1000"
        placeholder="详细描述你的问题..."
        show-word-limit
      />
      <van-field
        v-model="tagInput"
        label="标签"
        placeholder="添加标签，按回车确认"
        @keyup.enter="addTag"
      />
      <div class="tags-preview">
        <van-tag
          v-for="(tag, index) in form.tags"
          :key="index"
          type="primary"
          closeable
          @close="removeTag(index)"
          class="tag-item"
        >
          {{ tag }}
        </van-tag>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import { createQuestion } from '@/api/question'

const router = useRouter()
const tagInput = ref('')

const form = reactive({
  title: '',
  content: '',
  tags: []
})

const addTag = () => {
  const tag = tagInput.value.trim()
  if (tag && !form.tags.includes(tag)) {
    form.tags.push(tag)
    tagInput.value = ''
  }
}

const removeTag = (index) => {
  form.tags.splice(index, 1)
}

const onSubmit = async () => {
  if (!form.title.trim()) {
    showToast('请输入标题')
    return
  }
  
  try {
    await createQuestion(form)
    showToast('发布成功')
    router.back()
  } catch (error) {
    showToast('发布失败')
  }
}

const goBack = () => {
  router.back()
}
</script>

<style lang="scss" scoped>
.create-question-page {
  min-height: 100vh;
  background-color: #f5f5f5;
  
  .tags-preview {
    padding: 12px 16px;
    background: #fff;
    
    .tag-item {
      margin-right: 8px;
      margin-bottom: 8px;
    }
  }
}
</style>
