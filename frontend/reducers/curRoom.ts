import {
    LOG_OUT,
    CHANGE_ROOM,
    JOIN_ROOM
} from '../actions'

export default function curRoom(state = {}, action) {
    switch (action.type) {
        case LOG_OUT:
            return {}
        case CHANGE_ROOM:
            action.room.hasNotification = false
            return action.room
        case JOIN_ROOM:
            action.room.hasJoined = true
            return action.room
        default:
            return state
    }
}