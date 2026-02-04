import request from './request'
import type {
  Question,
  Answer,
  CreateQuestionRequest,
  CreateAnswerRequest,
  PaginatedResponse
} from '@/types'

export const getQuestions = (params?: {
  page?: number
  category?: string
}): Promise<PaginatedResponse<Question>> => {
  return request.get('/questions', { params })
}

export const getQuestionDetail = (id: string | number): Promise<Question> => {
  return request.get(`/question/${id}`)
}

export const createQuestion = (data: CreateQuestionRequest): Promise<{ id: number }> => {
  return request.post('/question', data)
}

export const createAnswer = (data: CreateAnswerRequest): Promise<{ id: number }> => {
  return request.post('/answer', data)
}

export const likeQuestion = (id: string | number): Promise<void> => {
  return request.post(`/question/${id}/like`)
}

export const unlikeQuestion = (id: string | number): Promise<void> => {
  return request.post(`/question/${id}/unlike`)
}
