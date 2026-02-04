<template>
  <div class="chat-page">

    <div class="content-wrapper">
      <!-- 智能体列表 -->
      <div class="section">
        <div class="section-title">智能体</div>
        <div class="agent-list">
          <div
            v-for="agent in agents"
            :key="agent.id"
            class="agent-item"
            @click="goToChat(agent.id)"
          >
            <div class="agent-avatar">
              <van-icon name="user-circle-o" size="48" />
            </div>
            <div class="agent-name">{{ agent.name }}</div>
            <div class="agent-desc">{{ agent.description }}</div>
          </div>
        </div>
      </div>

      <!-- 员工列表 -->
      <div class="section">
        <div class="section-title">我的员工</div>
        <div class="employee-list">
          <div
            v-for="emp in employees"
            :key="emp.id"
            class="employee-item card"
            @click="goToChat(emp.id)"
          >
            <div class="employee-info">
              <van-icon name="manager-o" size="32" />
              <div class="employee-detail">
                <div class="employee-name">{{ emp.name }}</div>
                <div class="employee-role">{{ emp.role }}</div>
              </div>
            </div>
            <van-icon name="arrow" />
          </div>
        </div>
      </div>

      <!-- 最近聊天 -->
      <div class="section">
        <div class="section-title">最近聊天</div>
        <div class="chat-list">
          <div
            v-for="chat in recentChats"
            :key="chat.id"
            class="chat-item card"
            @click="goToChat(chat.id)"
          >
            <div class="chat-avatar">
              <van-icon name="user-circle-o" size="40" />
            </div>
            <div class="chat-info">
              <div class="chat-name">{{ chat.name }}</div>
              <div class="chat-preview">{{ chat.lastMessage }}</div>
            </div>
            <div class="chat-meta">
              <div class="chat-time">{{ chat.time }}</div>
              <van-badge v-if="chat.unread" :content="chat.unread" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const agents = ref([
  { id: 1, name: '财经专家', description: '金融投资顾问' },
  { id: 2, name: '科技达人', description: '科技数码专家' },
  { id: 3, name: '心理咨询师', description: '情感心理顾问' },
  { id: 4, name: '职业规划师', description: '职业发展指导' },
])

const employees = ref([
  { id: 101, name: '年度总结助手', role: '数据分析' },
  { id: 102, name: '健康顾问', role: '身心健康' },
])

const recentChats = ref([
  { id: 1, name: '财经专家', lastMessage: '关于投资组合的建议...', time: '10:30', unread: 2 },
  { id: 2, name: '年度总结助手', lastMessage: '您的年度总结已生成', time: '昨天', unread: 0 },
])

const goToChat = (id) => {
  router.push(`/chat/${id}`)
}
</script>

<style lang="scss" scoped>
.chat-page {
  min-height: 100vh;

  .section {
    margin-bottom: 20px;

    .agent-list {
      display: flex;
      gap: 12px;
      overflow-x: auto;
      padding: 4px;

      .agent-item {
        flex-shrink: 0;
        width: 100px;
        text-align: center;
        padding: 12px 8px;
        background: #fff;
        border-radius: 12px;

        .agent-avatar {
          color: #667eea;
          margin-bottom: 8px;
        }

        .agent-name {
          font-size: 13px;
          font-weight: 500;
          color: #333;
          margin-bottom: 4px;
        }

        .agent-desc {
          font-size: 11px;
          color: #999;
        }
      }
    }

    .employee-list {
      .employee-item {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 8px;

        .employee-info {
          display: flex;
          align-items: center;
          gap: 12px;
          color: #667eea;

          .employee-detail {
            .employee-name {
              font-size: 15px;
              font-weight: 500;
              color: #333;
            }

            .employee-role {
              font-size: 12px;
              color: #999;
            }
          }
        }
      }
    }

    .chat-list {
      .chat-item {
        display: flex;
        align-items: center;
        gap: 12px;
        margin-bottom: 8px;

        .chat-avatar {
          color: #667eea;
        }

        .chat-info {
          flex: 1;

          .chat-name {
            font-size: 15px;
            font-weight: 500;
            color: #333;
            margin-bottom: 4px;
          }

          .chat-preview {
            font-size: 13px;
            color: #999;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
          }
        }

        .chat-meta {
          text-align: right;

          .chat-time {
            font-size: 12px;
            color: #999;
            margin-bottom: 4px;
          }
        }
      }
    }
  }
}
</style>
