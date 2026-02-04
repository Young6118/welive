// 通用响应类型
export interface ApiResponse<T> {
  code: number
  message: string
  data: T
}

// 分页响应类型
export interface PaginatedResponse<T> {
  list: T[]
  total: number
  page?: number
  pageSize?: number
}

// 用户类型
export interface User {
  id: number
  username: string
  email?: string
  avatar?: string
  bio?: string
  createdAt?: string
}

// 登录请求
export interface LoginRequest {
  username: string
  password: string
}

// 注册请求
export interface RegisterRequest {
  username: string
  password: string
  email?: string
}

// 问题类型
export interface Question {
  id: number
  title: string
  content: string
  authorId: number
  author?: User
  tags: string[]
  likes: number
  views?: number
  comments?: number
  createdAt?: string
  status?: number
  isLiked?: boolean
}

// 创建问题请求
export interface CreateQuestionRequest {
  title: string
  content: string
  tags?: string[]
}

// 回答类型
export interface Answer {
  id: number
  questionId: number
  content: string
  authorId: number
  author?: User
  likes: number
  isAI?: boolean
  createdAt?: string
}

// 创建回答请求
export interface CreateAnswerRequest {
  questionId: number
  content: string
}

// 笔记类型
export interface Note {
  id: number
  title: string
  content: string
  authorId: number
  author?: User
  category: string
  tags: string[]
  createdAt?: string
}

// 创建笔记请求
export interface CreateNoteRequest {
  title: string
  content: string
  category?: string
  tags?: string[]
}

// 笔记分类
export interface NoteCategory {
  id: number
  name: string
  userId?: number
  sort?: number
}

// 评论类型
export interface Comment {
  id: number
  targetId: number
  targetType: 'question' | 'answer' | 'note' | 'post'
  content: string
  authorId: number
  author?: User
  parentId?: number
  likes: number
  createdAt?: string
}

// 创建评论请求
export interface CreateCommentRequest {
  targetId: number
  targetType: string
  content: string
  parentId?: number
}

// 聊天消息类型
export interface Message {
  id: number
  chatId?: number
  senderId: number
  content: string
  type: 'text' | 'image' | 'file'
  status?: number
  createdAt?: string
  isSelf?: boolean
  time?: string
}

// 发送消息请求
export interface SendMessageRequest {
  receiverId: number
  content: string
  type?: string
}

// 智能体/员工类型
export interface Agent {
  id: number
  name: string
  description: string
  type: 'agent' | 'employee'
  role?: string
}

// 村落类型
export interface Village {
  id: number
  name: string
  description: string
  icon?: string
  category?: string
  memberCount: number
  postCount?: number
  isJoined?: boolean
}

// 帖子类型
export interface Post {
  id: number
  villageId: number
  authorId: number
  author?: User
  content: string
  images?: string[]
  likes: number
  comments: number
  isLiked?: boolean
  createdAt?: string
}

// 创建帖子请求
export interface CreatePostRequest {
  content: string
  images?: string[]
}

// 路由元信息
export interface RouteMeta {
  title?: string
  public?: boolean
  keepAlive?: boolean
}
