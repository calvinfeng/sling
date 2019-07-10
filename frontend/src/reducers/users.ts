import {
    LOG_OUT,
    NEW_USER,
    LOAD_USERS,
    CHANGE_ROOM,
    AppActionTypes
} from '../actions/types'

import { User } from '../types'

export default function users(state: User[] = [], action: AppActionTypes): User[] {
    switch (action.type) {
        case LOG_OUT:
            return []
        case NEW_USER:
            return state.concat(action.user)
        case LOAD_USERS:
            return action.users
        case CHANGE_ROOM:
            return []
        default:
            return state
    }
}