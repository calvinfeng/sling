import {
    LOG_OUT,
    CHANGE_ROOM,
    JOIN_ROOM,
    RoomAction
} from '../../actions/rooms'

import { Room } from '../../types'

export default function curRoom(state: Room | null = null, action: RoomAction): Room | null {
    switch (action.type) {
        case LOG_OUT:
            return null
        case CHANGE_ROOM:
            action.room.hasNotification = false
            return action.room
        case JOIN_ROOM:
            action.room.hasJoined = true
            action.room.hasNotification = false
            return action.room
        default:
            return state
    }
}