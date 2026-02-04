<template>
  <div class="discover-page">

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list
        v-model:loading="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load="onLoad"
      >
        <div class="content-wrapper">
          <!-- 搜索栏 -->
          <van-search
            v-model="searchValue"
            placeholder="搜索村落、话题..."
            shape="round"
          />

          <!-- 地球村列表 -->
          <div class="section">
            <div class="section-header">
              <div class="section-title">热门村落</div>
              <span class="more" @click="viewAllVillages">
                查看全部 <van-icon name="arrow" size="12" />
              </span>
            </div>
            <div class="village-grid">
              <div
                v-for="village in villages"
                :key="village.id"
                class="village-card"
                @click="goToVillage(village.id)"
              >
                <div class="village-icon">
                  <van-icon name="cluster-o" size="32" />
                </div>
                <div class="village-name">{{ village.name }}</div>
                <div class="village-members">{{ village.memberCount }}人</div>
              </div>
            </div>
          </div>

          <!-- 热门帖子 -->
          <div class="section">
            <div class="section-title">热门帖子</div>
            <div class="post-list">
              <div
                v-for="post in hotPosts"
                :key="post.id"
                class="post-card card"
              >
                <div class="post-header">
                  <div class="post-author">
                    <van-icon name="user-circle-o" size="24" />
                    <span>{{ post.author }}</span>
                  </div>
                  <span class="post-village">{{ post.village }}</span>
                </div>
                <div class="post-content">{{ post.content }}</div>
                <div class="post-footer">
                  <span class="post-action">
                    <van-icon name="good-job-o" />
                    {{ post.likes }}
                  </span>
                  <span class="post-action">
                    <van-icon name="comment-o" />
                    {{ post.comments }}
                  </span>
                  <span class="post-action">
                    <van-icon name="share-o" />
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </van-list>
    </van-pull-refresh>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getVillages } from '@/api/village'

const router = useRouter()
const searchValue = ref('')
const villages = ref([])
const hotPosts = ref([])
const loading = ref(false)
const finished = ref(false)
const refreshing = ref(false)
const page = ref(1)

const onLoad = async () => {
  if (refreshing.value) {
    villages.value = []
    page.value = 1
    refreshing.value = false
  }

  try {
    const res = await getVillages({ page: page.value })
    villages.value.push(...res.list)
    page.value++

    if (res.list.length < 10) {
      finished.value = true
    }
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const onRefresh = () => {
  finished.value = false
  loading.value = true
  onLoad()
}

const loadHotPosts = () => {
  // 模拟热门帖子数据
  hotPosts.value = [
    { id: 1, author: '用户A', village: '科技村', content: '分享一个AI工具...', likes: 128, comments: 32 },
    { id: 2, author: '用户B', village: '生活村', content: '今天学到了...', likes: 96, comments: 18 },
  ]
}

const goToVillage = (id) => {
  router.push(`/village/${id}`)
}

const viewAllVillages = () => {
  // TODO: 查看全部村落
}

onMounted(() => {
  onLoad()
  loadHotPosts()
})
</script>

<style lang="scss" scoped>
.discover-page {
  min-height: 100vh;

  .section {
    margin-bottom: 20px;

    .section-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 12px;

      .more {
        display: flex;
        align-items: center;
        gap: 2px;
        font-size: 13px;
        color: #667eea;
      }
    }

    .village-grid {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
      gap: 12px;

      .village-card {
        text-align: center;
        padding: 16px 8px;
        background: #fff;
        border-radius: 12px;

        .village-icon {
          color: #667eea;
          margin-bottom: 8px;
        }

        .village-name {
          font-size: 14px;
          font-weight: 500;
          color: #333;
          margin-bottom: 4px;
        }

        .village-members {
          font-size: 12px;
          color: #999;
        }
      }
    }

    .post-list {
      .post-card {
        margin-bottom: 12px;

        .post-header {
          display: flex;
          justify-content: space-between;
          align-items: center;
          margin-bottom: 8px;

          .post-author {
            display: flex;
            align-items: center;
            gap: 8px;
            font-size: 14px;
            color: #333;

            .van-icon {
              color: #667eea;
            }
          }

          .post-village {
            font-size: 12px;
            color: #667eea;
            background: rgba(102, 126, 234, 0.1);
            padding: 2px 8px;
            border-radius: 10px;
          }
        }

        .post-content {
          font-size: 14px;
          color: #666;
          margin-bottom: 12px;
          line-height: 1.6;
        }

        .post-footer {
          display: flex;
          gap: 16px;

          .post-action {
            display: flex;
            align-items: center;
            gap: 4px;
            font-size: 13px;
            color: #999;
          }
        }
      }
    }
  }
}
</style>
