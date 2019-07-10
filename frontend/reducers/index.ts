import { combineReducers } from 'redux'
import rooms from './rooms'

const slingApp = combineReducers({
    curUser,
    rooms,
    curRoom,
    users,
    messages
})

export default slingApp