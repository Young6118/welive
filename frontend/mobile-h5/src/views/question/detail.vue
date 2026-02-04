<template>
  <div class="question-detail-page">
    <van-nav-bar title="问题详情" left-arrow @click-left="goBack" />
    
    <div class="content-wrapper">
      <!-- 问题内容 -->
      <div class="question-section card">
        <h1 class="question-title">{{ question.title }}</h1>
        <div class="question-author">
          <van-icon name="user-circle-o" />
          <span>{{ question.author }}</span>
        </div>
        <p class="question-content">{{ question.content }}</p>
        <div class="question-tags">
          <van-tag
            v-for="tag in question.tags"
            :key="tag"
            type="primary"
            size="small"
            class="tag"
          >
            {{ tag }}
          </van-tag>
        </div>
        <div class="question-actions">
          <div class="action-item" @click="toggleLike">
            <van-icon :name="isLiked ? 'good-job' : 'good-job-o'" :color="isLiked ? '#ee0a24' : ''" />
            <span>{{ question.likes }}</span>
          </div>
          <div class="action-item">
            <van-icon name="comment-o" />
            <span>{{ question.comments }}</span>
          </div>
          <div class="action-item">
            <van-icon name="share-o" />
            <span>分享</span>
          </div>
        </div>
      </div>
      
      <!-- 回答列表 -->
      <div class="answers-section">
        <div class="section-title">{{ answers.length }} 个回答</div>
        <div class="answer-list">
          <div
            v-for="answer in answers"
            :key="answer.id"
            class="answer-card card"
          >
            <div class="answer-header">
              <van-icon name="user-circle-o" />
              <span class="answer-author">{{ answer.author }}</span>
            </div>
            <p class="answer-content">{{ answer.content }}</p>
            <div class="answer-footer">
              <span class="answer-time">{{ answer.time }}</span>
              <div class="answer-actions">
                <span class="action">
                  <van-icon name="good-job-o" />
                  {{ answer.likes }}
                </span>
                <span class="action">
                  <van-icon name="comment-o" />
                  回复
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 底部操作栏 -->
    <div class="bottom-bar">
      <van-field
        v-model="answerContent"
        placeholder="写回答..."
        class="answer-input"
      />
      <van-button type="primary" size="small" @click="submitAnswer">发布</van-button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { showToast } from 'vant'
import { getQuestionDetail, createAnswer, likeQuestion, unlikeQuestion } from '@/api/question'

const route = useRoute()
const router = useRouter()
const questionId = route.params.id

const question = ref({})
const answers = ref([])
const answerContent = ref('')
const isLiked = ref(false)

const loadQuestion = async () => {
  try {
    const res = await getQuestionDetail(questionId)
    question.value = res
  } catch (error) {
    console.error(error)
  }
}

const toggleLike = async () => {
  try {
    if (isLiked.value) {
      await unlikeQuestion(questionId)
      question.value.likes--
    } else {
      await likeQuestion(questionId)
      question.value.likes++
    }
    isLiked.value = !isLiked.value
  } catch (error) {
    showToast('操作失败')
  }
}

const submitAnswer = async () => {
  if (!answerContent.value.trim()) {
    showToast('请输入回答内容')
    return
  }
  
  try {
    await createAnswer({
      questionId: questionId,
      content: answerContent.value
    })
    showToast('回答成功')
    answerContent.value = ''
    loadQuestion()
  } catch (error) {
    showToast('回答失败')
  }
}

const goBack = () => {
  router.back()
}

onMounted(() => {
  loadQuestion()
  // 模拟回答数据
  answers.value = [
    { id: 1, author: '用户A', content: '这是一个很好的回答...', likes: 10, time: '2小时前' },
    { id: 2, author: '用户B', content: '我也这么认为...', likes: 5, time: '1小时前' },
  ]
})
</script>

<style lang="scss" scoped>
.question-detail-page {
  min-height: 100vh;
  background-color: #f5f5f5;
  padding-bottom: 60px;
  
  .question-section {
    .question-title {
      font-size: 18px;
      font-weight: 600;
      color: #333;
      margin-bottom: 12px;
    }
    
    .question-author {
      display: flex;
      align-items: center;
      gap: 8px;
      color: #666;
      font-size: 14px;
      margin-bottom: 12px;
    }
    
    .question-content {
      font-size: 15px;
      color: #333;
      line-height: 1.8;
      margin-bottom: 12px;
    }
    
    .question-tags {
      margin-bottom: 16px;
      
      .tag {
        margin-right: 8px;
      }
    }
    
    .question-actions {
      display: flex;
      gap: 24px;
      padding-top: 16px;
      border-top: 1px solid #f0f0f0;
      
      .action-item {
        display: flex;
        align-items: center;
        gap: 4px;
        font-size: 14px;
        color: #666;
      }
    }
  }
  
  .answers-section {
    margin-top: 16px;
    
    .answer-list {
      .answer-card {
        margin-bottom: 8px;
        
        .answer-header {
          display: flex;
          align-items: center;
          gap: 8px;
          color: #667eea;
          margin-bottom: 8px;
          
          .answer-author {
            font-size: 14px;
            font-weight: 500;
          }
        }
        
        .answer-content {
          font-size: 14px;
          color: #333;
          line-height: 1.6;
          margin-bottom: 12px;
        }
        
        .answer-footer {
          display: flex;
          justify-content: space-between;
          align-items: center;
          font-size: 12px;
          color: #999;
          
          .answer-actions {
            display: flex;
            gap: 16px;
            
            .action {
              display: flex;
              align-items: center;
              gap: 4px;
            }
          }
        }
      }
    }
  }
  
  .bottom-bar {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 8px 16px;
    background: #fff;
    border-top: 1px solid #f0f0f0;
    
    .answer-input {
      flex: 1;
      margin: 0;
      padding: 8px 12px;
      background: #f5f5f5;
      border-radius: 20px;
    }
  }
}
</style>
