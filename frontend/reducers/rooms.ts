import {
    LOG_OUT,
    NEW_ROOM,
    LOAD_ROOMS,
    JOIN_ROOM,
    CHANGE_ROOM,
    MARK_UNREAD
} from '../actions'

export default function rooms(state = [], action) {
    switch (action.type) {
        case LOG_OUT:
        case NEW_ROOM:
        case LOAD_ROOMS:
        case JOIN_ROOM:
        case CHANGE_ROOM:
        case MARK_UNREAD:
        default:
            return state
    }
}