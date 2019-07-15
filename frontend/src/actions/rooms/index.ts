import { Room } from '../../types'

export const NEW_ROOM = 'NEW_ROOM'
export const LOAD_ROOMS = 'LOAD_ROOMS'
export const CHANGE_ROOM = 'CHANGE_ROOM'
export const MARK_UNREAD = 'MARK_UNREAD'
export const JOIN_ROOM = 'JOIN_ROOM'

type NewRoomAction = {
    type: typeof NEW_ROOM
    room: Room
}

type LoadRoomsAction = {
    type: typeof LOAD_ROOMS
    rooms: Room[]
}

type ChangeRoomAction = {
    type: typeof CHANGE_ROOM
    room: Room
}

type MarkUnreadAction = {
    type: typeof MARK_UNREAD
    roomID: number
}

type JoinRoomAction = {
    type: typeof JOIN_ROOM
    room: Room
}

export type RoomAction = NewRoomAction |
    LoadRoomsAction |
    ChangeRoomAction |
    MarkUnreadAction |
    JoinRoomAction