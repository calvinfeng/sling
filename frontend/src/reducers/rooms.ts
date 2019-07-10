import {
    LOG_OUT,
    NEW_ROOM,
    LOAD_ROOMS,
    MARK_UNREAD,
    AppActionTypes
} from '../actions/types'

import { Room } from '../types'

export default function rooms(state: Room[] = [], action: AppActionTypes): Room[] {
    switch (action.type) {
        case LOG_OUT:
            return []
        case NEW_ROOM:
            return state.concat(action.room)
        case LOAD_ROOMS:
            return action.rooms
        case MARK_UNREAD:
            let newState = state.slice()
            newState.forEach((room) => {
                if (room.id === action.room.id) {
                    room.hasNotification = true
                }
            })
            return newState
        default:
            return state
    }
}