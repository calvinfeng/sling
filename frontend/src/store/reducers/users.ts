import {
    LOG_OUT,
    NEW_USER,
    LOAD_USERS,
    UserAction
} from '../../actions/users'

import { User } from '../../types'

export default function users(state: User[] = [], action: UserAction): User[] {
    switch (action.type) {
        case LOG_OUT:
            return []
        case NEW_USER:
            return state.concat(action.user)
        case LOAD_USERS:
            return action.users
        default:
            return state
    }
}