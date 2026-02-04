import request from './request'
import type { Question, Note } from '@/types'

interface SearchParams {
  keyword: string
  page?: number
  pageSize?: number
}

interface SearchResult<T> {
  list: T[]
  total: number
}

export const searchQuestions = (params: SearchParams): Promise<SearchResult<Question>> => {
  return request.get('/search/questions', { params })
}

export const searchNotes = (params: SearchParams): Promise<SearchResult<Note>> => {
  return request.get('/search/notes', { params })
}
