import {
    LOG_OUT,
    NEW_ROOM,
    LOAD_ROOMS,
    MARK_UNREAD,
    CHANGE_ROOM,
    RoomAction,
    RoomStoreState
} from '../../actions/rooms'
import { Room } from '../../types'

const initial: RoomStoreState = {
    loading: false,
    data: [],
    current: null,
    error: ""
}

export default function rooms(state = initial, action: RoomAction): RoomStoreState {
    switch (action.type) {
        case LOG_OUT:
            return initial
        case NEW_ROOM:
            return Object.assign({}, state, {
                data: [...state.data, action.room],
                error: ''
            })
        case LOAD_ROOMS:
            return Object.assign({}, state, {
                data: action.rooms,
                error: ''
            })
        case MARK_UNREAD:
            let newData = state.data.slice()
            newData.forEach((room: Room) => {
                if (room.id === action.roomID) {
                    room.hasNotification = true
                }
            })
            return Object.assign({}, state, {
                data: newData
            })
        case CHANGE_ROOM:
            return Object.assign({}, state, {
                current: action.room
            })
        default:
            return state
    }
}