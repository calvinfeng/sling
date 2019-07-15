import { User } from '../../types'
import { 
    LOG_IN,
    LOG_OUT,
    NEW_USER,
    LOAD_USERS,
    UserAction
} from '.'

export const logIn = (user: User): UserAction => ({
    type: LOG_IN,
    user
})

export const logOut = (): UserAction => ({
    type: LOG_OUT
})

export const newUser = (user: User): UserAction => ({
    type: NEW_USER,
    user
})

export const loadUsers = (users: User[]): UserAction => ({
    type: LOAD_USERS,
    users
})