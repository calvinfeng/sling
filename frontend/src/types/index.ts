export interface User {
    username: string
    id?: number
    jwtToken: string
}

export interface Message {
    username: string
    time: Date
    body: string
}

export interface Room {
    id: number
    name: string
    hasJoined: boolean
    hasNotification: boolean
    isDM: boolean
}