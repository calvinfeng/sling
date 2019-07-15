import { combineReducers } from 'redux'
import rooms from './reducers/rooms'
import curUser from './reducers/curUser'
import curRoom from './reducers/curRoom'
import users from './reducers/users'
import messages from './reducers/messages'

export const rootReducer = combineReducers({
    curUser,
    rooms,
    curRoom,
    users,
    messages
})

export type AppState = ReturnType<typeof rootReducer>