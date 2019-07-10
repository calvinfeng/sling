import {
    LOG_IN,
    LOG_OUT,
    AppActionTypes
} from '../actions/types'

import { User } from '../types'

export default function curUser(state: User | undefined | null, action: AppActionTypes): User | undefined | null {
    switch (action.type) {
        case LOG_IN:
            return Object.assign({}, action.user)
        case LOG_OUT:
            return null
        default:
            return state
    }
}