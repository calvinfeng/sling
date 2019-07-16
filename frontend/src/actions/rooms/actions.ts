import { Room } from '../../types'
import {
    NEW_ROOM,
    LOAD_ROOMS,
    SET_ROOMS,
    CHANGE_ROOM,
    MARK_UNREAD,
    JOIN_ROOM,
    RoomAction,
    RoomStoreState,
    START_ROOM_LOADING,
    STOP_ROOM_LOADING,
    FAIL_ROOM_LOADING
} from '.'
import { ThunkDispatch } from 'redux-thunk'
import axios, { AxiosResponse } from 'axios'
import { AppState } from '../../store'

export const newRoom = (room: Room): RoomAction => ({
    type: NEW_ROOM,
    room
})

export const loadRooms = (token: string) => async (dispatch: ThunkDispatch<AppState, undefined, RoomAction>) => {
    dispatch({ type: START_ROOM_LOADING })

    try {
        const res: AxiosResponse = await axios.get('api/rooms', {
            headers: {
                'Token': token
            }
        })

        let rooms = res.data.map((room: any): Room => ({
            id: room.id,
            name: room.name,
            hasJoined: room.hasJoined,
            hasNotification: room.hasNotification,
            isDM: room.type === 1
        }))

        dispatch({ type: SET_ROOMS, rooms })

    } catch(err) {
        dispatch({ type: FAIL_ROOM_LOADING, message: 'Failed to fetch rooms: ' + err })
    }


    dispatch({ type: STOP_ROOM_LOADING })
}

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