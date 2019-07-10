import { combineReducers } from 'redux'
import rooms from './rooms'
import curUser from './curUser'
import curRoom from './curRoom'
import users from './users'
import messages from './messages'

const rootReducer = combineReducers({
    curUser,
    rooms,
    curRoom,
    users,
    messages
})

export type AppState = ReturnType<typeof rootReducer>