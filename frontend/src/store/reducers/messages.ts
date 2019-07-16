import {
    NEW_MESSAGE,
    LOAD_MESSAGES,
    MessageStoreState,
    MessageAction,
    CLEAR_MESSAGES
} from '../../actions/messages'

const initial: MessageStoreState = {
    loading: false,
    data: [],
    error: ''
}

export default function messages(state = initial, action: MessageAction): MessageStoreState {
    switch (action.type) {
        case CLEAR_MESSAGES:
            return initial
        case NEW_MESSAGE:
            return Object.assign({}, state, {
                data: [...state.data, action.message],
                error: ''
            })
        case LOAD_MESSAGES:
            return Object.assign({}, state, {
                data: action.messages,
                error: ''
            })
        default:
            return state
    }
}