import {
    NEW_USER,
    LOAD_USERS,
    UserAction,
    UserStoreState,
    CLEAR_USERS,
    SET_USERS,
    START_USER_LOADING,
    STOP_USER_LOADING,
    FAIL_USER_LOADING
} from '../../actions/users'

import { User } from '../../types'

const initial: UserStoreState = {
    loading: false,
    data: [],
    error: ''
}

export default function users(state = initial, action: UserAction): UserStoreState {
    switch (action.type) {
        case CLEAR_USERS:
            return initial
        case NEW_USER:
            return Object.assign({}, state, {
                data: [...state.data, action.user],
                error: ''
            })
        case SET_USERS:
            return Object.assign({}, state, {
                data: action.users,
                error: ''
            })
        case START_USER_LOADING:
            return Object.assign({}, state, {
                loading: true
            })
        case STOP_USER_LOADING:
            return Object.assign({}, state, {
                loading: false
            })
        case FAIL_USER_LOADING:
            return Object.assign({}, state, {
                error: action.message
            })
        default:
            return state
    }
}