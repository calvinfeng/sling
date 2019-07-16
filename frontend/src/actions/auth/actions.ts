import { User } from '../../types'
import {
    START_AUTH_LOADING,
    STOP_AUTH_LOADING,
    FAIL_AUTH_LOADING,
    SET_USER,
    CLEAR_USER,
    AuthAction
} from '.'
import { ThunkDispatch } from 'redux-thunk'
import axios, { AxiosResponse } from 'axios'
import { AppState } from '../../store'
import { CLEAR_MESSAGES } from '../messages';
import { CLEAR_ROOMS } from '../rooms';
import { CLEAR_USERS } from '../users';
import { AppActionTypes } from '../types';

export const logIn = (username: string, password: string) =>
    async (dispatch: ThunkDispatch<AppState, undefined, AuthAction>) => {
        dispatch({ type: START_AUTH_LOADING })

        try {
            const res: AxiosResponse = await axios.post("api/login", {
                name: username,
                password: password
            })

            localStorage.setItem('jwt_token', res.data.jwt_token)

            let user: User = {
                id: res.data.id,
                username: res.data.name,
                jwtToken: res.data.jwt_token
            }

            dispatch({
                type: SET_USER,
                user: user
            })
        } catch {
            dispatch({
                type: FAIL_AUTH_LOADING,
                message: "invalid credentials"
            })
        }

        dispatch({ type: STOP_AUTH_LOADING })
    }

export const register = (username: string, password: string, email: string) =>
    async (dispatch: ThunkDispatch<AppState, undefined, AuthAction>) => {
        dispatch({ type: START_AUTH_LOADING })

        try {
            const res: AxiosResponse = await axios.post("api/register", {
                name: username,
                password: password,
                email: email
            })

            localStorage.setItem('jwt_token', res.data.jwt_token)

            let user: User = {
                id: res.data.id,
                username: res.data.name,
                jwtToken: res.data.jwt_token
            }

            dispatch({
                type: SET_USER,
                user: user
            })
        } catch {
            dispatch({
                type: FAIL_AUTH_LOADING,
                message: "invalid credentials"
            })
        }

        dispatch({ type: STOP_AUTH_LOADING })
    }

export const getCurrentUser = (token: string) =>
    async (dispatch: ThunkDispatch<AppState, undefined, AuthAction>) => {
        dispatch({ type: START_AUTH_LOADING })

        try {
            const res: AxiosResponse = await axios.get('api/users/current', {
                headers: { 'Token': token }
            })

            let user: User = {
                id: res.data.id,
                username: res.data.name,
                jwtToken: res.data.jwt_token
            }

            dispatch({
                type: SET_USER,
                user
            })
        } catch {
            dispatch({
                type: FAIL_AUTH_LOADING,
                message: "invalid token"
            })
        }
    }

export const logOut = () => (dispatch: ThunkDispatch<AppState, undefined, AppActionTypes>) => {
    dispatch({ type: CLEAR_USER })
    dispatch({ type: CLEAR_MESSAGES })
    dispatch({ type: CLEAR_ROOMS })
    dispatch({ type: CLEAR_USERS })
}