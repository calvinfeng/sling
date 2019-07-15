import { User } from '../../types'
import {
    NEW_USER,
    LOAD_USERS,
    UserAction
} from '.'

export const newUser = (user: User): UserAction => ({
    type: NEW_USER,
    user
})

export const loadUsers = (users: User[]): UserAction => ({
    type: LOAD_USERS,
    users
})