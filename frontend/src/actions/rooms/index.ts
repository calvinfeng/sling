import { Room } from '../../types'

export const NEW_ROOM = 'NEW_ROOM'
export const CHANGE_ROOM = 'CHANGE_ROOM'
export const SET_ROOMS = 'SET_ROOMS'
export const MARK_UNREAD = 'MARK_UNREAD'
export const JOIN_ROOM = 'JOIN_ROOM'
export const CLEAR_ROOMS = 'CLEAR_ROOMS'
export const START_ROOM_LOADING = 'START_ROOM_LOADING'
export const STOP_ROOM_LOADING = 'STOP_ROOM_LOADING'
export const FAIL_ROOM_LOADING = 'FAIL_ROOM_LOADING'

type ClearRoomsAction = {
    type: typeof CLEAR_ROOMS
}

type StartRoomLoadingAction = {
    type: typeof START_ROOM_LOADING
}

type StopRoomLoadingAction = {
    type: typeof STOP_ROOM_LOADING
}

type FailRoomLoadingAction = {
    type: typeof FAIL_ROOM_LOADING
    message: string
}

type NewRoomAction = {
    type: typeof NEW_ROOM
    room: Room
}

type SetRoomsAction = {
    type: typeof SET_ROOMS
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
    ChangeRoomAction |
    MarkUnreadAction |
    JoinRoomAction | 
    ClearRoomsAction | 
    StartRoomLoadingAction | 
    StopRoomLoadingAction | 
    FailRoomLoadingAction | 
    SetRoomsAction

export type RoomStoreState = {
    loading: boolean
    data: Room[]
    current: Room | null
    error: string
}