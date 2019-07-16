import { User } from '../../types'

export const LOG_IN = 'LOG_IN'
export const LOG_OUT = 'LOG_OUT'
export const SET_USER = 'SET_USER'
export const CLEAR_USER = 'CLEAR_USER'
export const START_AUTH_LOADING = 'START_AUTH_LOADING'
export const STOP_AUTH_LOADING = 'STOP_AUTH_LOADING'
export const FAIL_AUTH_LOADING = 'FAIL_AUTH_LOADING'
export const GET_CURRENT_USER = 'GET_CURRENT_USER'

type StartAuthLoadingAction = {
    type: typeof START_AUTH_LOADING
}

type StopAuthLoadingAction = {
    type: typeof STOP_AUTH_LOADING
}

type FailAuthLoadingAction = {
    type: typeof FAIL_AUTH_LOADING
    message: string
}

type LogInAction = {
    type: typeof LOG_IN
    username: string
    password: string
}

type LogOutAction = {
    type: typeof LOG_OUT
}

type SetUserAction = {
    type: typeof SET_USER
    user: User
}

type ClearUserAction = {
    type: typeof CLEAR_USER
}

type GetCurrentUserAction = {
    type: typeof GET_CURRENT_USER
    token: string
}

export type AuthAction = LogInAction |
    LogOutAction | 
    StartAuthLoadingAction | 
    StopAuthLoadingAction | 
    FailAuthLoadingAction | 
    SetUserAction | 
    ClearUserAction | 
    GetCurrentUserAction

export type AuthStoreState = {
    loading: boolean
    user: User | null
    error: string
}