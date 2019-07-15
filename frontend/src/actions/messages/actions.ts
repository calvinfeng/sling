import { Message } from '../../types'
import {
    NEW_MESSAGE,
    LOAD_MESSAGES,
    MessageAction
} from '.'

export const newMessage = (message: Message): MessageAction => ({
    type: NEW_MESSAGE,
    message
})

export const loadMessages = (messages: Message[]): MessageAction => ({
    type: LOAD_MESSAGES,
    messages
})
