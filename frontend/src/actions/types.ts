// Import types
import { User, Message, Room } from '../types'

// Define action constants
export const LOG_IN = 'LOG_IN'
export const LOG_OUT = 'LOG_OUT'
export const NEW_ROOM = 'NEW_ROOM'
export const NEW_MESSAGE = 'NEW_MESSAGE'
export const NEW_USER = 'NEW_USER'
export const LOAD_ROOMS = 'LOAD_ROOMS'
export const LOAD_MESSAGES = 'LOAD_MESSAGES'
export const LOAD_USERS = 'LOAD_USERS'
export const CHANGE_ROOM = 'CHANGE_ROOM'
export const MARK_UNREAD = 'MARK_UNREAD'
export const JOIN_ROOM = 'JOIN_ROOM'

// Define action types
type LogInAction = {
    type: typeof LOG_IN
    user: User
}

type LogOutAction = {
    type: typeof LOG_OUT
}

type NewRoomAction = {
    type: typeof NEW_ROOM
    room: Room
}

type NewMessageAction = {
    type: typeof NEW_MESSAGE
    message: Message
}

type NewUserAction = {
    type: typeof NEW_USER
    user: User
}

type LoadRoomsAction = {
    type: typeof LOAD_ROOMS
    rooms: Room[]
}

type LoadMessagesAction = {
    type: typeof LOAD_MESSAGES
    messages: Message[]
}

type LoadUsersAction = {
    type: typeof LOAD_USERS
    users: User[]
}

type ChangeRoomAction = {
    type: typeof CHANGE_ROOM
    room: Room
}

type MarkUnreadAction = {
    type: typeof MARK_UNREAD
    room: Room
}

type JoinRoomAction = {
    type: typeof JOIN_ROOM
    room: Room
}

export type AppActionTypes = (
    LogInAction | LogOutAction | 
    NewRoomAction | NewMessageAction | NewUserAction |
    LoadRoomsAction | LoadMessagesAction | LoadUsersAction |
    ChangeRoomAction | MarkUnreadAction | JoinRoomAction
)
