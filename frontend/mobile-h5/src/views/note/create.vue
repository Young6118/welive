<template>
  <div class="create-note-page">
    <van-nav-bar
      title="创建笔记"
      left-arrow
      @click-left="goBack"
      right-text="保存"
      @click-right="onSubmit"
    />
    
    <div class="content-wrapper">
      <van-field
        v-model="form.title"
        label="标题"
        placeholder="请输入笔记标题"
      />
      <van-field
        v-model="form.category"
        label="分类"
        placeholder="选择分类"
        readonly
        @click="showCategoryPicker = true"
      />
      <van-field
        v-model="form.content"
        rows="10"
        autosize
        label="内容"
        type="textarea"
        maxlength="5000"
        placeholder="记录你的想法..."
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
          type="success"
          closeable
          @close="removeTag(index)"
          class="tag-item"
        >
          {{ tag }}
        </van-tag>
      </div>
    </div>
    
    <van-popup v-model:show="showCategoryPicker" position="bottom">
      <van-picker
        :columns="categories"
        @confirm="onCategoryConfirm"
        @cancel="showCategoryPicker = false"
      />
    </van-popup>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import { createNote } from '@/api/note'

const router = useRouter()
const tagInput = ref('')
const showCategoryPicker = ref(false)

const form = reactive({
  title: '',
  content: '',
  category: '',
  tags: []
})

const categories = ['学习', '工作', '生活', '灵感', '其他']

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

const onCategoryConfirm = (value) => {
  form.category = value
  showCategoryPicker.value = false
}

const onSubmit = async () => {
  if (!form.title.trim()) {
    showToast('请输入标题')
    return
  }
  
  try {
    await createNote(form)
    showToast('保存成功')
    router.back()
  } catch (error) {
    showToast('保存失败')
  }
}

const goBack = () => {
  router.back()
}
</script>

<style lang="scss" scoped>
.create-note-page {
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
