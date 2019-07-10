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

// Define action creators
export const logIn = user => ({
    type: LOG_IN,
    user
})

export const logOut = () => ({
    type: LOG_OUT
})

export const newRoom = room => ({
    type: NEW_ROOM,
    room
})

export const newMessage = message => ({
    type: NEW_MESSAGE,
    message
})

export const newUser = user => ({
    type: NEW_USER,
    user
})

export const loadRooms = rooms => ({
    type: LOAD_ROOMS,
    rooms
})

export const loadMessages = messages => ({
    type: LOAD_MESSAGES,
    messages
})

export const loadUsers = users => ({
    type: LOAD_USERS,
    users
})

export const changeRoom = roomID => ({
    type: CHANGE_ROOM,
    roomID
})

export const markUnread = roomID => ({
    type: MARK_UNREAD,
    roomID
})

export const joinRoom = roomID => ({
    type: JOIN_ROOM,
    roomID
})