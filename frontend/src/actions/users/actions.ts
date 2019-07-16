import { User } from '../../types'
import {
    NEW_USER,
    LOAD_USERS,
    SET_USERS,
    UserAction,
    UserStoreState,
    START_USER_LOADING,
    FAIL_USER_LOADING,
    STOP_USER_LOADING
} from '.'
import { ThunkDispatch } from 'redux-thunk';
import axios, { AxiosResponse } from 'axios';
import { AppState } from '../../store'

export const newUser = (user: User): UserAction => ({
    type: NEW_USER,
    user
})
export const loadUsers = (token: string) => async (dispatch: ThunkDispatch<AppState, undefined, UserAction>) => {
    dispatch({ type: START_USER_LOADING })

    try {
        const res: AxiosResponse = await axios.get('api/users/', {
            headers: {
                'Token': token
            }
        })

        let users = res.data.map((user: any): User => ({
            username: user.name,
            id: user.id,
            jwtToken: null // can't get other peoples' tokens
        }))

        dispatch({ type: SET_USERS, users })

    } catch(err) {
        dispatch({ type: FAIL_USER_LOADING, message: 'Failed to fetch users: ' + err })
    }


    dispatch({ type: STOP_USER_LOADING })
}