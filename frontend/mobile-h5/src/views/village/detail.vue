<template>
  <div class="village-detail-page">
    <van-nav-bar :title="village.name" left-arrow @click-left="goBack">
      <template #right>
        <van-button
          size="small"
          :type="isJoined ? 'default' : 'primary'"
          @click="toggleJoin"
        >
          {{ isJoined ? '已加入' : '加入' }}
        </van-button>
      </template>
    </van-nav-bar>
    
    <div class="content-wrapper">
      <!-- 村落信息 -->
      <div class="village-info card">
        <div class="village-header">
          <van-icon name="cluster-o" size="48" color="#667eea" />
          <div class="village-meta">
            <h2 class="village-name">{{ village.name }}</h2>
            <p class="village-members">{{ village.memberCount }} 成员</p>
          </div>
        </div>
        <p class="village-desc">{{ village.description }}</p>
      </div>
      
      <!-- 帖子列表 -->
      <div class="posts-section">
        <div class="section-title">最新帖子</div>
        <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
          <van-list
            v-model:loading="loading"
            :finished="finished"
            finished-text="没有更多了"
            @load="onLoad"
          >
            <div class="post-list">
              <div
                v-for="post in posts"
                :key="post.id"
                class="post-card card"
              >
                <div class="post-header">
                  <div class="post-author">
                    <van-icon name="user-circle-o" size="24" />
                    <span>{{ post.author }}</span>
                  </div>
                </div>
                <div class="post-content">{{ post.content }}</div>
                <div class="post-footer">
                  <span class="post-action" @click="likePost(post)">
                    <van-icon :name="post.isLiked ? 'good-job' : 'good-job-o'" />
                    {{ post.likes }}
                  </span>
                  <span class="post-action">
                    <van-icon name="comment-o" />
                    {{ post.comments }}
                  </span>
                </div>
              </div>
            </div>
          </van-list>
        </van-pull-refresh>
      </div>
    </div>
    
    <!-- 发帖按钮 -->
    <div class="fab-button" @click="showPostDialog = true">
      <van-icon name="edit" size="24" />
    </div>
    
    <!-- 发帖弹窗 -->
    <van-dialog
      v-model:show="showPostDialog"
      title="发布帖子"
      show-cancel-button
      @confirm="submitPost"
    >
      <van-field
        v-model="postContent"
        rows="4"
        autosize
        type="textarea"
        maxlength="500"
        placeholder="分享你的想法..."
        show-word-limit
      />
    </van-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { showToast } from 'vant'
import { getVillageDetail, getPosts, joinVillage, leaveVillage, createPost } from '@/api/village'

const route = useRoute()
const router = useRouter()
const villageId = route.params.id

const village = ref({})
const posts = ref([])
const isJoined = ref(false)
const loading = ref(false)
const finished = ref(false)
const refreshing = ref(false)
const page = ref(1)
const showPostDialog = ref(false)
const postContent = ref('')

const loadVillage = async () => {
  try {
    const res = await getVillageDetail(villageId)
    village.value = res
  } catch (error) {
    village.value = {
      name: '科技村',
      description: '科技交流社区，分享最新的科技资讯和见解',
      memberCount: 1280
    }
  }
}

const onLoad = async () => {
  if (refreshing.value) {
    posts.value = []
    page.value = 1
    refreshing.value = false
  }
  
  try {
    const res = await getPosts(villageId, { page: page.value })
    posts.value.push(...res.list)
    page.value++
    
    if (res.list.length < 10) {
      finished.value = true
    }
  } catch (error) {
    // 模拟数据
    posts.value = [
      { id: 1, author: '用户A', content: '今天学到了很多...', likes: 10, comments: 3, isLiked: false },
      { id: 2, author: '用户B', content: '分享一个有趣的发现...', likes: 5, comments: 1, isLiked: true },
    ]
    finished.value = true
  } finally {
    loading.value = false
  }
}

const onRefresh = () => {
  finished.value = false
  loading.value = true
  onLoad()
}

const toggleJoin = async () => {
  try {
    if (isJoined.value) {
      await leaveVillage({ villageId })
      isJoined.value = false
      village.value.memberCount--
      showToast('已退出村落')
    } else {
      await joinVillage({ villageId })
      isJoined.value = true
      village.value.memberCount++
      showToast('加入成功')
    }
  } catch (error) {
    showToast('操作失败')
  }
}

const likePost = (post) => {
  post.isLiked = !post.isLiked
  post.likes += post.isLiked ? 1 : -1
}

const submitPost = async () => {
  if (!postContent.value.trim()) {
    showToast('请输入内容')
    return
  }
  
  try {
    await createPost(villageId, { content: postContent.value })
    showToast('发布成功')
    postContent.value = ''
    onRefresh()
  } catch (error) {
    showToast('发布失败')
  }
}

const goBack = () => {
  router.back()
}

onMounted(() => {
  loadVillage()
  onLoad()
})
</script>

<style lang="scss" scoped>
.village-detail-page {
  min-height: 100vh;
  background-color: #f5f5f5;
  
  .village-info {
    .village-header {
      display: flex;
      align-items: center;
      gap: 16px;
      margin-bottom: 12px;
      
      .village-meta {
        .village-name {
          font-size: 18px;
          font-weight: 600;
          color: #333;
          margin-bottom: 4px;
        }
        
        .village-members {
          font-size: 13px;
          color: #999;
        }
      }
    }
    
    .village-desc {
      font-size: 14px;
      color: #666;
      line-height: 1.6;
    }
  }
  
  .posts-section {
    margin-top: 16px;
    
    .post-list {
      .post-card {
        margin-bottom: 8px;
        
        .post-header {
          display: flex;
          align-items: center;
          gap: 8px;
          color: #666;
          font-size: 14px;
          margin-bottom: 8px;
        }
        
        .post-content {
          font-size: 14px;
          color: #333;
          line-height: 1.6;
          margin-bottom: 12px;
        }
        
        .post-footer {
          display: flex;
          gap: 20px;
          
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
  
  .fab-button {
    position: fixed;
    right: 20px;
    bottom: 20px;
    width: 50px;
    height: 50px;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #fff;
    box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
    z-index: 99;
  }
}
</style>
