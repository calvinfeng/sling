import {
    LOG_OUT,
    NEW_MESSAGE,
    LOAD_MESSAGES,
    CHANGE_ROOM,
    JOIN_ROOM,
    AppActionTypes
} from '../../actions/types'

import { Message } from '../../types'

export default function messages(state: Message[] = [], action: AppActionTypes) {
    switch (action.type) {
        case LOG_OUT:
        case CHANGE_ROOM:
        case JOIN_ROOM:
            return []
        case NEW_MESSAGE:
            return state.concat(action.message)
        case LOAD_MESSAGES:
            return action.messages
        default:
            return state
    }
}