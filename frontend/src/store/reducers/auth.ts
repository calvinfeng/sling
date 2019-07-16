import {
    AuthStoreState, AuthAction, START_AUTH_LOADING, STOP_AUTH_LOADING, FAIL_AUTH_LOADING, SET_USER, CLEAR_USER
} from '../../actions/auth'

const initial: AuthStoreState = {
    loading: false,
    user: null,
    error: ''
}

export default function auth(state = initial, action: AuthAction): AuthStoreState {
    switch (action.type) {
        case SET_USER:
            return Object.assign({}, state, {
                user: action.user,
                error: ''
            })
        case CLEAR_USER:
            return initial
        case START_AUTH_LOADING:
            return Object.assign({}, state, {
                loading: true
            })
        case STOP_AUTH_LOADING:
            return Object.assign({}, state, {
                loading: false
            })
        case FAIL_AUTH_LOADING:
            return Object.assign({}, state, {
                error: action.message
            })
        default:
            return state
    }
}