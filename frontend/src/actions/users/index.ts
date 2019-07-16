import { User } from '../../types'

export const NEW_USER = 'NEW_USER'
export const LOAD_USERS = 'LOAD_USERS'
export const SET_USERS = 'SET_USERS'
export const CLEAR_USERS = 'CLEAR_USERS'
export const START_USER_LOADING = 'START_USER_LOADING'
export const STOP_USER_LOADING = 'STOP_USER_LOADING'
export const FAIL_USER_LOADING = 'FAIL_USER_LOADING'

type NewUserAction = {
    type: typeof NEW_USER
    user: User
}

type SetUsersAction = {
    type: typeof SET_USERS
    users: User[]
}

type LoadUsersAction = {
    type: typeof LOAD_USERS
    token: string
}

type ClearUsersAction = {
    type: typeof CLEAR_USERS
}

type StartUserLoadingAction = {
    type: typeof START_USER_LOADING
}

type StopUserLoadingAction = {
    type: typeof STOP_USER_LOADING
}

type FailUserLoadingAction = {
    type: typeof FAIL_USER_LOADING
    message: string
}

export type UserAction = NewUserAction |
    LoadUsersAction |
    ClearUsersAction | 
    StartUserLoadingAction | 
    StopUserLoadingAction |
    FailUserLoadingAction | 
    SetUsersAction

export type UserStoreState = {
    loading: boolean
    data: User[]
    error: string
}