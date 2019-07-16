import { Dispatch } from 'react'
import { AppActionTypes } from '../actions/types'
import * as msgActions from '../actions/messages/actions'
import * as roomActions from '../actions/rooms/actions'
import * as userActions from '../actions/users/actions'
import * as authActions from '../actions/auth/actions'
import { User, Message, Room } from '../types'
import { ThunkDispatch } from 'redux-thunk';
import { AppState } from '.';

export const dispatchActions = (dispatch: ThunkDispatch<AppState, undefined, AppActionTypes>) => {
    return {
        onLogIn: async (username: string, password: string) => {
            await dispatch(authActions.logIn(username, password))
        },
        onLogOut: () => {
            dispatch(authActions.logOut())
        },
        authenticate: async (token: string) => {
            await dispatch(authActions.getCurrentUser(token))
        },

        onLoadMessages: (messages: Message[]) => {
            dispatch(msgActions.loadMessages(messages))
        },
        onLoadRooms: async (token: string) => {
            await dispatch(roomActions.loadRooms(token))
        },
        onLoadUsers: async (token: string) => {
            await dispatch(userActions.loadUsers(token))
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