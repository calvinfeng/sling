import { User } from '../../types'

export const SET_USER = 'SET_USER'
export const CLEAR_USER = 'CLEAR_USER'
export const START_AUTH_LOADING = 'START_AUTH_LOADING'
export const STOP_AUTH_LOADING = 'STOP_AUTH_LOADING'
export const FAIL_AUTH_LOADING = 'FAIL_AUTH_LOADING'

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

type SetUserAction = {
    type: typeof SET_USER
    user: User
}

type ClearUserAction = {
    type: typeof CLEAR_USER
}

export type AuthAction = StartAuthLoadingAction | 
    StopAuthLoadingAction | 
    FailAuthLoadingAction | 
    SetUserAction | 
    ClearUserAction

export type AuthStoreState = {
    loading: boolean
    user: User | null
    error: string
}