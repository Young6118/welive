import request from './request'
import type {
  Note,
  NoteCategory,
  CreateNoteRequest,
  PaginatedResponse
} from '@/types'

export const getNotes = (params?: {
  page?: number
  category?: string | null
}): Promise<PaginatedResponse<Note>> => {
  return request.get('/notes', { params })
}

export const getNoteDetail = (id: string | number): Promise<Note> => {
  return request.get(`/note/${id}`)
}

export const createNote = (data: CreateNoteRequest): Promise<{ id: number }> => {
  return request.post('/note', data)
}

export const getNoteCategories = (): Promise<NoteCategory[]> => {
  return request.get('/note/categories')
}

export const getNotesByCategory = (
  id: string | number,
  params?: { page?: number }
): Promise<PaginatedResponse<Note>> => {
  return request.get(`/note/category/${id}`, { params })
}
