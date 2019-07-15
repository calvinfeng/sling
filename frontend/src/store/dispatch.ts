import { Dispatch } from 'react'
import { AppActionTypes } from '../actions/types'
import * as msgActions from '../actions/messages/actions'
import * as roomActions from '../actions/rooms/actions'
import * as userActions from '../actions/users/actions'
import * as authActions from '../actions/auth/actions'
import { User, Message, Room } from '../types'

export const dispatchActions = (dispatch: Dispatch<AppActionTypes>) => {
    return {
        onLogIn: (user: User) => {
            dispatch(authActions.logIn(user))
        },
        onLogOut: () => {
            dispatch(authActions.logOut())
        },

        onLoadMessages: (messages: Message[]) => {
            dispatch(msgActions.loadMessages(messages))
        },
        onLoadRooms: (rooms: Room[]) => {
            dispatch(roomActions.loadRooms(rooms))
        },
        onLoadUsers: (users: User[]) => {
            dispatch(userActions.loadUsers(users))
        },

        onNewMessage: (message: Message) => {
            dispatch(msgActions.newMessage(message))
        },
        onNewRoom: (room: Room) => {
            dispatch(roomActions.newRoom(room))
        },
        onNewUser: (user: User) => {
            dispatch(userActions.newUser(user))
        },

        onMarkUnread: (roomID: number) => {
            dispatch(roomActions.markUnread(roomID))
        },
        onJoinRoom: (room: Room) => {
            dispatch(roomActions.joinRoom(room))
        },
        onChangeRoom: (room: Room) => {
            dispatch(roomActions.changeRoom(room))
        },
    }
}