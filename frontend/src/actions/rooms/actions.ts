import { Room } from '../../types'
import { 
    NEW_ROOM,
    LOAD_ROOMS,
    CHANGE_ROOM,
    MARK_UNREAD,
    JOIN_ROOM,
    AppActionTypes
} from '../types'

export const newRoom = (room: Room): AppActionTypes => ({
    type: NEW_ROOM,
    room
})

export const loadRooms = (rooms: Room[]): AppActionTypes => ({
    type: LOAD_ROOMS,
    rooms
})

export const changeRoom = (room: Room): AppActionTypes => ({
    type: CHANGE_ROOM,
    room
})

export const markUnread = (roomID: number): AppActionTypes => ({
    type: MARK_UNREAD,
    roomID
})

export const joinRoom = (room: Room): AppActionTypes => ({
    type: JOIN_ROOM,
    room
})