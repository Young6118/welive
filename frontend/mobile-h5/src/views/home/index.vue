<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import { getQuestions, likeQuestion, unlikeQuestion } from '@/api/question'
import { searchQuestions } from '@/api/search'
import type { Question } from '@/types'

const router = useRouter()
const searchValue = ref('')
const activeCategory = ref('recommend')
const questionList = ref<Question[]>([])
const loading = ref(false)
const finished = ref(false)
const refreshing = ref(false)
const page = ref(1)

const onLoad = async (): Promise<void> => {
  if (refreshing.value) {
    questionList.value = []
    page.value = 1
    refreshing.value = false
  }

  try {
    const res = await getQuestions({
      page: page.value,
      category: activeCategory.value
    })

    questionList.value.push(...res.list)
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

const onRefresh = (): void => {
  finished.value = false
  loading.value = true
  onLoad()
}

const onSearch = async (): Promise<void> => {
  if (!searchValue.value.trim()) {
    // 如果搜索为空，重新加载默认列表
    page.value = 1
    questionList.value = []
    finished.value = false
    onLoad()
    return
  }

  try {
    loading.value = true
    const res = await searchQuestions({
      keyword: searchValue.value,
      page: 1,
      pageSize: 10
    })
    questionList.value = res.list
    finished.value = true
  } catch (error) {
    console.error(error)
    showToast('搜索失败')
  } finally {
    loading.value = false
  }
}

const goToDetail = (id: number): void => {
  router.push(`/question/${id}`)
}

const goToCreate = (): void => {
  router.push('/question/create')
}

const onLike = async (item: Question, event: Event): Promise<void> => {
  event.stopPropagation()
  try {
    if (item.isLiked) {
      await unlikeQuestion(item.id)
      item.likes--
      item.isLiked = false
      showToast('已取消点赞')
    } else {
      await likeQuestion(item.id)
      item.likes++
      item.isLiked = true
      showToast('点赞成功')
    }
  } catch (error) {
    console.error(error)
    showToast('操作失败')
  }
}

onMounted(() => {
  onLoad()
})
</script>

<template>
  <div class="home-page">
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
            placeholder="搜索问题、笔记..."
            shape="round"
            @search="onSearch"
          />

          <!-- 分类标签 -->
          <div class="category-tabs">
            <van-tabs v-model:active="activeCategory" sticky>
              <van-tab title="推荐" name="recommend"></van-tab>
              <van-tab title="科技" name="tech"></van-tab>
              <van-tab title="生活" name="life"></van-tab>
              <van-tab title="情感" name="emotion"></van-tab>
              <van-tab title="财经" name="finance"></van-tab>
            </van-tabs>
          </div>

          <!-- 问题列表 -->
          <div class="question-list">
            <div
              v-for="item in questionList"
              :key="item.id"
              class="question-card card"
              @click="goToDetail(item.id)"
            >
              <h3 class="question-title">{{ item.title }}</h3>
              <p class="question-content text-ellipsis-2">{{ item.content }}</p>
              <div class="question-footer">
                <div class="tags">
                  <van-tag
                    v-for="tag in item.tags"
                    :key="tag"
                    type="primary"
                    size="small"
                    class="tag"
                  >
                    {{ tag }}
                  </van-tag>
                </div>
                <div class="stats">
                  <span
                    class="stat-item"
                    :class="{ liked: item.isLiked }"
                    @click="onLike(item, $event)"
                  >
                    <van-icon :name="item.isLiked ? 'good-job' : 'good-job-o'" />
                    {{ item.likes }}
                  </span>
                  <span class="stat-item">
                    <van-icon name="comment-o" />
                    {{ item.comments }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </van-list>
    </van-pull-refresh>

    <!-- 发布按钮 -->
    <div class="fab-button" @click="goToCreate">
      <van-icon name="plus" size="24" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.home-page {
  min-height: 100vh;

  .category-tabs {
    margin: 0 -12px;
  }

  .question-list {
    margin-top: 12px;
  }

  .question-card {
    .question-title {
      margin-bottom: 8px;
      font-size: 16px;
      font-weight: 600;
      color: #333;
    }

    .question-content {
      margin-bottom: 12px;
      font-size: 14px;
      color: #666;
    }

    .question-footer {
      display: flex;
      align-items: center;
      justify-content: space-between;

      .tags {
        .tag {
          margin-right: 8px;
        }
      }

      .stats {
        display: flex;
        gap: 16px;

        .stat-item {
          display: flex;
          align-items: center;
          gap: 4px;
          font-size: 12px;
          color: #999;
          cursor: pointer;

          &.liked {
            color: #667eea;
          }

          &:active {
            opacity: 0.7;
          }
        }
      }
    }
  }

  .fab-button {
    position: fixed;
    right: 20px;
    bottom: 70px;
    z-index: 99;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 50px;
    height: 50px;
    color: #fff;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    border-radius: 50%;
    box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
  }
}
</style>
