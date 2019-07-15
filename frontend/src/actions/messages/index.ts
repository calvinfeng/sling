import { Message } from '../../types'

export const NEW_MESSAGE = 'NEW_MESSAGE'
export const LOAD_MESSAGES = 'LOAD_MESSAGES'

type NewMessageAction = {
    type: typeof NEW_MESSAGE
    message: Message
}

type LoadMessagesAction = {
    type: typeof LOAD_MESSAGES
    messages: Message[]
}

export type MessageAction = NewMessageAction | LoadMessagesAction