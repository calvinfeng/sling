import { User } from '../../types'

export const LOG_IN = 'LOG_IN'
export const LOG_OUT = 'LOG_OUT'
export const NEW_USER = 'NEW_USER'
export const LOAD_USERS = 'LOAD_USERS'

type LogInAction = {
    type: typeof LOG_IN
    user: User
}

type LogOutAction = {
    type: typeof LOG_OUT
}

type NewUserAction = {
    type: typeof NEW_USER
    user: User
}

type LoadUsersAction = {
    type: typeof LOAD_USERS
    users: User[]
}

export type UserAction = LogInAction |
    LogOutAction |
    NewUserAction |
    LoadUsersAction