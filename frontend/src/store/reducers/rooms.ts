import {
    NEW_ROOM,
    CLEAR_ROOMS,
    MARK_UNREAD,
    SET_ROOMS,
    CHANGE_ROOM,
    RoomAction,
    RoomStoreState,
    START_ROOM_LOADING,
    STOP_ROOM_LOADING,
    FAIL_ROOM_LOADING
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
        case CLEAR_ROOMS:
            return initial
        case NEW_ROOM:
            return Object.assign({}, state, {
                data: [...state.data, action.room],
                error: ''
            })
        case SET_ROOMS:
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
        case START_ROOM_LOADING:
            return Object.assign({}, state, {
                loading: true
            })
        case STOP_ROOM_LOADING:
            return Object.assign({}, state, {
                loading: false
            })
        case FAIL_ROOM_LOADING:
            return Object.assign({}, state, {
                error: action.message
            })
        default:
            return state
    }
}