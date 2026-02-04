import request from './request'
import type {
  User,
  LoginRequest,
  RegisterRequest,
  ApiResponse,
  UpdateUserRequest
} from '@/types'

export const login = (data: LoginRequest): Promise<{ token: string; userId: number }> => {
  return request.post('/login', data)
}

export const register = (data: RegisterRequest): Promise<{ userId: number }> => {
  return request.post('/register', data)
}

export const logout = (): Promise<void> => {
  return request.post('/logout')
}

export const checkLogin = (): Promise<{ isLoggedIn: boolean; userId: number }> => {
  return request.get('/check-login')
}

export const getUserInfo = (): Promise<User> => {
  return request.get('/user')
}

export const updateUserInfo = (data: UpdateUserRequest): Promise<void> => {
  return request.put('/user', data)
}
