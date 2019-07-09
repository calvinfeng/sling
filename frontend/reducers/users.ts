import {
    LOG_OUT,
    NEW_USER,
    LOAD_USERS,
    CHANGE_ROOM,
} from '../actions'

export default function users(state = [], action) {
    switch (action.type) {
        case LOG_OUT:
            return []
        case NEW_USER:
            return state.concat(action.message)
        case LOAD_USERS:
            return action.users
        case CHANGE_ROOM:
            return []
        default:
            return state
    }
}