import { User } from '../../types'
import { 
    LOG_IN,
    LOG_OUT,
    NEW_USER,
    LOAD_USERS,
    AppActionTypes
} from '../types'

export const logIn = (user: User): AppActionTypes => ({
    type: LOG_IN,
    user
})

export const logOut = (): AppActionTypes => ({
    type: LOG_OUT
})

export const newUser = (user: User): AppActionTypes => ({
    type: NEW_USER,
    user
})

export const loadUsers = (users: User[]): AppActionTypes => ({
    type: LOAD_USERS,
    users
})