export type User = {
    username: string
    id: number | null
    jwtToken: string | null
}

export type Message = {
    msgID: number | null
    userID: number | null
    username: string
    time: Date
    body: string
}

export type Room = {
    id: number
    name: string
    hasJoined: boolean
    hasNotification: boolean
    isDM: boolean
}