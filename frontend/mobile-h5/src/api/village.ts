import request from './request'
import type {
  Village,
  Post,
  CreatePostRequest,
  PaginatedResponse
} from '@/types'

export const getVillages = (params?: {
  page?: number
}): Promise<PaginatedResponse<Village>> => {
  return request.get('/earth-villages', { params })
}

export const getVillageDetail = (id: string | number): Promise<Village> => {
  return request.get(`/earth-village/${id}`)
}

export const joinVillage = (data: { villageId: string | number }): Promise<void> => {
  return request.post('/earth-village/join', data)
}

export const leaveVillage = (data: { villageId: string | number }): Promise<void> => {
  return request.post('/earth-village/leave', data)
}

export const getPosts = (
  id: string | number,
  params?: { page?: number }
): Promise<PaginatedResponse<Post>> => {
  return request.get(`/earth-village/${id}/posts`, { params })
}

export const createPost = (
  id: string | number,
  data: CreatePostRequest
): Promise<{ id: number }> => {
  return request.post(`/earth-village/${id}/post`, data)
}

export const likePost = (villageId: string | number, postId: string | number): Promise<void> => {
  return request.post(`/earth-village/${villageId}/post/${postId}/like`)
}

export const unlikePost = (villageId: string | number, postId: string | number): Promise<void> => {
  return request.post(`/earth-village/${villageId}/post/${postId}/unlike`)
}

export const deletePost = (villageId: string | number, postId: string | number): Promise<void> => {
  return request.delete(`/earth-village/${villageId}/post/${postId}`)
}

export const replyPost = (
  villageId: string | number,
  postId: string | number,
  data: { content: string }
): Promise<{ id: number }> => {
  return request.post(`/earth-village/${villageId}/post/${postId}/reply`, data)
}

export const getReplies = (
  villageId: string | number,
  postId: string | number,
  params?: { page?: number }
): Promise<PaginatedResponse<Post>> => {
  return request.get(`/earth-village/${villageId}/post/${postId}/replies`, { params })
}
