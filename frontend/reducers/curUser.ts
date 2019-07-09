import {
    LOG_IN,
    LOG_OUT
} from '../actions'

export default function curUser(state = {}, action) {
    switch (action.type) {
        case LOG_IN:
            return Object.assign({}, action.user)
        case LOG_OUT:
            return {}
        default:
            return state
    }
}