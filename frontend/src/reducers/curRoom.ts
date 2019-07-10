import {
    LOG_OUT,
    CHANGE_ROOM,
    JOIN_ROOM,
    AppActionTypes
} from '../actions/types'

import { Room } from '../types'

export default function curRoom(state: Room | undefined | null, action: AppActionTypes): Room | undefined | null {
    switch (action.type) {
        case LOG_OUT:
            return null
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