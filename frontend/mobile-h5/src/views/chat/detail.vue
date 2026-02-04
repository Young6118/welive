<template>
  <div class="chat-detail-page">
    <van-nav-bar :title="chatTitle" left-arrow @click-left="goBack" />
    
    <div class="chat-container" ref="chatContainer">
      <div class="message-list">
        <div
          v-for="msg in messages"
          :key="msg.id"
          class="message-item"
          :class="{ 'message-self': msg.isSelf }"
        >
          <div class="message-avatar">
            <van-icon name="user-circle-o" size="36" />
          </div>
          <div class="message-content">
            <div class="message-bubble">{{ msg.content }}</div>
            <div class="message-time">{{ msg.time }}</div>
          </div>
        </div>
      </div>
    </div>
    
    <div class="input-bar">
      <van-field
        v-model="messageText"
        placeholder="输入消息..."
        class="message-input"
        @keyup.enter="sendMessage"
      />
      <van-button type="primary" icon="guide-o" @click="sendMessage" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const chatId = route.params.id
const chatTitle = ref('聊天')
const chatContainer = ref(null)
const messageText = ref('')

const messages = ref([
  { id: 1, content: '你好！我是AI助手，有什么可以帮助你的吗？', isSelf: false, time: '10:00' },
  { id: 2, content: '你好，我想咨询一些问题', isSelf: true, time: '10:01' },
])

const scrollToBottom = () => {
  nextTick(() => {
    if (chatContainer.value) {
      chatContainer.value.scrollTop = chatContainer.value.scrollHeight
    }
  })
}

const sendMessage = () => {
  if (!messageText.value.trim()) return
  
  const newMessage = {
    id: Date.now(),
    content: messageText.value,
    isSelf: true,
    time: new Date().toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit' })
  }
  
  messages.value.push(newMessage)
  messageText.value = ''
  scrollToBottom()
  
  // 模拟AI回复
  setTimeout(() => {
    messages.value.push({
      id: Date.now() + 1,
      content: '收到您的消息，我正在思考如何回答...',
      isSelf: false,
      time: new Date().toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit' })
    })
    scrollToBottom()
  }, 1000)
}

const goBack = () => {
  router.back()
}

onMounted(() => {
  scrollToBottom()
})
</script>

<style lang="scss" scoped>
.chat-detail-page {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #f5f5f5;
  
  .chat-container {
    flex: 1;
    overflow-y: auto;
    padding: 12px;
    
    .message-list {
      .message-item {
        display: flex;
        gap: 8px;
        margin-bottom: 16px;
        
        .message-avatar {
          color: #999;
          flex-shrink: 0;
        }
        
        .message-content {
          max-width: 70%;
          
          .message-bubble {
            background: #fff;
            padding: 10px 14px;
            border-radius: 8px;
            font-size: 14px;
            color: #333;
            line-height: 1.5;
          }
          
          .message-time {
            font-size: 11px;
            color: #999;
            margin-top: 4px;
          }
        }
        
        &.message-self {
          flex-direction: row-reverse;
          
          .message-content {
            .message-bubble {
              background: #667eea;
              color: #fff;
            }
            
            .message-time {
              text-align: right;
            }
          }
        }
      }
    }
  }
  
  .input-bar {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 8px 12px;
    background: #fff;
    border-top: 1px solid #f0f0f0;
    
    .message-input {
      flex: 1;
      margin: 0;
      padding: 8px 12px;
      background: #f5f5f5;
      border-radius: 20px;
    }
  }
}
</style>
