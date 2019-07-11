// Import types
import { User, Message, Room } from '../types'
import { 
    LOG_IN,
    LOG_OUT,
    NEW_ROOM,
    NEW_MESSAGE,
    NEW_USER,
    LOAD_ROOMS,
    LOAD_MESSAGES,
    LOAD_USERS,
    CHANGE_ROOM,
    MARK_UNREAD,
    JOIN_ROOM,
    AppActionTypes
} from './types'

// Define action creators
export const logIn = (user: User): AppActionTypes => ({
    type: LOG_IN,
    user
})

export const logOut = (): AppActionTypes => ({
    type: LOG_OUT
})

export const newRoom = (room: Room): AppActionTypes => ({
    type: NEW_ROOM,
    room
})

export const newMessage = (message: Message): AppActionTypes => ({
    type: NEW_MESSAGE,
    message
})


export const newUser = (user: User): AppActionTypes => ({
    type: NEW_USER,
    user
})

export const loadRooms = (rooms: Room[]): AppActionTypes => ({
    type: LOAD_ROOMS,
    rooms
})

export const loadMessages = (messages: Message[]): AppActionTypes => ({
    type: LOAD_MESSAGES,
    messages
})

export const loadUsers = (users: User[]): AppActionTypes => ({
    type: LOAD_USERS,
    users
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