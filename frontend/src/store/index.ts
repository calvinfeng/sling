import { combineReducers } from 'redux'
import rooms from './reducers/rooms'
import auth from './reducers/auth'
import users from './reducers/users'
import messages from './reducers/messages'

export const rootReducer = combineReducers({
    auth,
    rooms,
    users,
    messages
})

export type AppState = ReturnType<typeof rootReducer>