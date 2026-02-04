import request from './request'
import type { Message, SendMessageRequest, PaginatedResponse } from '@/types'

export const getChatHistory = (
  id: string | number,
  params?: { page?: number }
): Promise<PaginatedResponse<Message>> => {
  return request.get(`/chat/${id}`, { params })
}

export const sendMessage = (data: SendMessageRequest): Promise<{ id: number }> => {
  return request.post('/chat', data)
}
