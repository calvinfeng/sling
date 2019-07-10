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
interface LogInAction {
    type: typeof LOG_IN
    user: User
}

interface LogOutAction {
    type: typeof LOG_OUT
}

interface NewRoomAction {
    type: typeof NEW_ROOM
    room: Room
}

interface NewMessageAction {
    type: typeof NEW_MESSAGE
    message: Message
}

interface NewUserAction {
    type: typeof NEW_USER
    user: User
}

interface LoadRoomsAction {
    type: typeof LOAD_ROOMS
    rooms: Room[]
}

interface LoadMessagesAction {
    type: typeof LOAD_MESSAGES
    messages: Message[]
}

interface LoadUsersAction {
    type: typeof LOAD_USERS
    users: User[]
}

interface ChangeRoomAction {
    type: typeof CHANGE_ROOM
    room: Room
}

interface MarkUnreadAction {
    type: typeof MARK_UNREAD
    room: Room
}

interface JoinRoomAction {
    type: typeof JOIN_ROOM
    room: Room
}

export type AppActionTypes = (
    LogInAction | LogOutAction | 
    NewRoomAction | NewMessageAction | NewUserAction |
    LoadRoomsAction | LoadMessagesAction | LoadUsersAction |
    ChangeRoomAction | MarkUnreadAction | JoinRoomAction
)
