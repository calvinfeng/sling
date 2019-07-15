import { User } from '../../types'
import { 
    LOG_IN,
    LOG_OUT,
    START_AUTH_LOADING,
    STOP_AUTH_LOADING,
    FAIL_AUTH_LOADING,
    SET_USER,
    CLEAR_USER
    AuthAction,
    AuthStoreState
} from '.'
import { ThunkDispatch } from 'redux-thunk'
import axios, { AxiosResponse } from 'axios'

export const logIn = (username: string, password: string) => 
    async (dispatch: ThunkDispatch<AuthStoreState, undefined, AuthAction>) => {
    dispatch({ type: START_AUTH_LOADING })
    
    try {
        const res: AxiosResponse = await axios.post("api/login", {
            name: username,
            password: password
        })
        
        dispatch({
            type: "???",
            user: res.data
        })
    } catch {
        dispatch({
            type: FAIL_AUTH_LOADING,
            message: "invalid credentials"
        })
    }

    dispatch({ type: STOP_AUTH_LOADING })
}

export const logOut = (): AuthAction => ({
    type: LOG_OUT
})