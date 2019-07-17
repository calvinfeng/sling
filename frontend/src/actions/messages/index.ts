import { Message } from '../../types'

export const NEW_MESSAGE = 'NEW_MESSAGE'
export const LOAD_MESSAGES = 'LOAD_MESSAGES'
export const CLEAR_MESSAGES = 'CLEAR_MESSAGES'
export const START_MESSAGE_LOADING = 'START_MESSAGE_LOADING'
export const STOP_MESSAGE_LOADING = 'STOP_MESSAGE_LOADING'
export const FAIL_MESSAGE_LOADING = 'FAIL_MESSAGE_LOADING'

type ClearMessagesAction = {
    type: typeof CLEAR_MESSAGES
}

type StartMessageLoadingAction = {
    type: typeof START_MESSAGE_LOADING
}

type StopMessageLoadingAction = {
    type: typeof STOP_MESSAGE_LOADING
}

type FailMessageLoadingAction = {
    type: typeof FAIL_MESSAGE_LOADING
    message: string
}

type NewMessageAction = {
    type: typeof NEW_MESSAGE
    message: Message
}

type LoadMessagesAction = {
    type: typeof LOAD_MESSAGES
    messages: Message[]
}

export type MessageAction = NewMessageAction |
    LoadMessagesAction | 
    ClearMessagesAction | 
    StartMessageLoadingAction | 
    StopMessageLoadingAction | 
    FailMessageLoadingAction

export type MessageStoreState = {
    loading: boolean
    data: Message[]
    error: string
}