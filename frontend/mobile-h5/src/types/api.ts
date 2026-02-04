// API 相关类型定义

// 更新用户信息请求
export interface UpdateUserRequest {
  username?: string
  email?: string
  avatar?: string
  bio?: string
}

// 回复评论请求
export interface ReplyCommentRequest {
  content: string
}

// 创建帖子请求
export interface CreatePostRequest {
  content: string
  images?: string[]
}
