import {
    LOG_OUT,
    NEW_ROOM,
    LOAD_ROOMS,
    MARK_UNREAD
} from '../actions'

export default function rooms(state = [], action) {
    switch (action.type) {
        case LOG_OUT:
            return []
        case NEW_ROOM:
            return state.concat(action.room)
        case LOAD_ROOMS:
            return action.rooms
        case MARK_UNREAD:
            return state.slice().forEach((room) => {
                if (room.id === action.room.id) {
                    room.hasNotification = true
                }
            })
        default:
            return state
    }
}