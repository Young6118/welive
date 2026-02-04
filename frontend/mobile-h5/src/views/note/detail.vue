<template>
  <div class="note-detail-page">
    <van-nav-bar title="笔记详情" left-arrow @click-left="goBack" />
    
    <div class="content-wrapper">
      <div class="note-card card">
        <h1 class="note-title">{{ note.title }}</h1>
        <div class="note-meta">
          <van-tag type="success" size="small">{{ note.category }}</van-tag>
          <span class="note-time">{{ note.time }}</span>
        </div>
        <div class="note-content">{{ note.content }}</div>
        <div class="note-tags">
          <van-tag
            v-for="tag in note.tags"
            :key="tag"
            type="primary"
            size="small"
            class="tag"
          >
            {{ tag }}
          </van-tag>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getNoteDetail } from '@/api/note'

const route = useRoute()
const router = useRouter()
const noteId = route.params.id

const note = ref({})

const loadNote = async () => {
  try {
    const res = await getNoteDetail(noteId)
    note.value = res
  } catch (error) {
    // 模拟数据
    note.value = {
      title: '示例笔记标题',
      content: '这是笔记的详细内容...',
      category: '学习',
      tags: ['AI', '学习'],
      time: '2024-01-01 12:00'
    }
  }
}

const goBack = () => {
  router.back()
}

onMounted(() => {
  loadNote()
})
</script>

<style lang="scss" scoped>
.note-detail-page {
  min-height: 100vh;
  background-color: #f5f5f5;
  
  .note-card {
    .note-title {
      font-size: 20px;
      font-weight: 600;
      color: #333;
      margin-bottom: 12px;
    }
    
    .note-meta {
      display: flex;
      align-items: center;
      gap: 12px;
      margin-bottom: 16px;
      
      .note-time {
        font-size: 13px;
        color: #999;
      }
    }
    
    .note-content {
      font-size: 15px;
      color: #333;
      line-height: 1.8;
      margin-bottom: 16px;
      white-space: pre-wrap;
    }
    
    .note-tags {
      .tag {
        margin-right: 8px;
      }
    }
  }
}
</style>
