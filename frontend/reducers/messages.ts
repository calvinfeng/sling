import {
    LOG_OUT,
    NEW_MESSAGE,
    LOAD_MESSAGES,
    CHANGE_ROOM,
} from '../actions'

export default function messages(state = [], action) {
    switch (action.type) {
        case LOG_OUT:
            return []
        case NEW_MESSAGE:
            return state.concat(action.message)
        case LOAD_MESSAGES:
            return action.messages
        case CHANGE_ROOM:
            return []
        default:
            return state
    }
}