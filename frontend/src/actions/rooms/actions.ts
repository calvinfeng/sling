import { Room } from '../../types'
import { 
    NEW_ROOM,
    LOAD_ROOMS,
    CHANGE_ROOM,
    MARK_UNREAD,
    JOIN_ROOM,
    RoomAction
} from '.'

export const newRoom = (room: Room): RoomAction => ({
    type: NEW_ROOM,
    room
})

export const loadRooms = (rooms: Room[]): RoomAction => ({
    type: LOAD_ROOMS,
    rooms
})

export const changeRoom = (room: Room): RoomAction => ({
    type: CHANGE_ROOM,
    room
})

export const markUnread = (roomID: number): RoomAction => ({
    type: MARK_UNREAD,
    roomID
})

export const joinRoom = (room: Room): RoomAction => ({
    type: JOIN_ROOM,
    room
})