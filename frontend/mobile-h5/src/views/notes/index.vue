<template>
  <div class="notes-page">

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list
        v-model:loading="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load="onLoad"
      >
        <div class="content-wrapper">
          <!-- 分类选择 -->
          <div class="category-section">
            <div class="section-title">分类</div>
            <div class="category-list">
              <div
                v-for="cat in categories"
                :key="cat.id"
                class="category-item"
                :class="{ active: activeCategory === cat.id }"
                @click="selectCategory(cat.id)"
              >
                {{ cat.name }}
              </div>
            </div>
          </div>

          <!-- 笔记列表 -->
          <div class="note-list">
            <div
              v-for="note in noteList"
              :key="note.id"
              class="note-card card"
              @click="goToDetail(note.id)"
            >
              <h3 class="note-title">{{ note.title }}</h3>
              <p class="note-content text-ellipsis-2">{{ note.content }}</p>
              <div class="note-footer">
                <van-tag type="success" size="small">{{ note.category }}</van-tag>
                <span class="note-time">2小时前</span>
              </div>
            </div>
          </div>
        </div>
      </van-list>
    </van-pull-refresh>

    <!-- 创建按钮 -->
    <div class="fab-button" @click="goToCreate">
      <van-icon name="plus" size="24" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getNotes, getNoteCategories } from '@/api/note'

const router = useRouter()
const noteList = ref([])
const categories = ref([])
const activeCategory = ref(null)
const loading = ref(false)
const finished = ref(false)
const refreshing = ref(false)
const page = ref(1)

const onLoad = async () => {
  if (refreshing.value) {
    noteList.value = []
    page.value = 1
    refreshing.value = false
  }

  try {
    const res = await getNotes({
      page: page.value,
      category: activeCategory.value
    })

    noteList.value.push(...res.list)
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

const loadCategories = async () => {
  try {
    const res = await getNoteCategories()
    categories.value = res
  } catch (error) {
    console.error(error)
  }
}

const selectCategory = (id) => {
  activeCategory.value = activeCategory.value === id ? null : id
  onRefresh()
}

const goToDetail = (id) => {
  router.push(`/note/${id}`)
}

const goToCreate = () => {
  router.push('/note/create')
}

onMounted(() => {
  loadCategories()
  onLoad()
})
</script>

<style lang="scss" scoped>
.notes-page {
  min-height: 100vh;

  .category-section {
    margin-bottom: 16px;

    .category-list {
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
      margin-top: 8px;

      .category-item {
        padding: 6px 16px;
        background: #fff;
        border-radius: 16px;
        font-size: 13px;
        color: #666;
        cursor: pointer;

        &.active {
          background: #667eea;
          color: #fff;
        }
      }
    }
  }

  .note-list {
    .note-card {
      .note-title {
        font-size: 16px;
        font-weight: 600;
        color: #333;
        margin-bottom: 8px;
      }

      .note-content {
        font-size: 14px;
        color: #666;
        margin-bottom: 12px;
      }

      .note-footer {
        display: flex;
        justify-content: space-between;
        align-items: center;

        .note-time {
          font-size: 12px;
          color: #999;
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
