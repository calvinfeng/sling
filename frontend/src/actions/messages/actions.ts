import { Message } from '../../types'
import {
    NEW_MESSAGE,
    LOAD_MESSAGES,
    AppActionTypes
} from '../types'

export const newMessage = (message: Message): AppActionTypes => ({
    type: NEW_MESSAGE,
    message
})

export const loadMessages = (messages: Message[]): AppActionTypes => ({
    type: LOAD_MESSAGES,
    messages
})