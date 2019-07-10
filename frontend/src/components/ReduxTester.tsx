import React, { Component, Dispatch } from 'react';
import { connect } from 'react-redux'
import { AppState } from '../store'
import * as actions from '../actions'
import { AppActionTypes } from '../actions/types'

import { User, Room, Message } from '../types'

const mapStateToProps = (state: AppState) => state
const mapDispatchToProps = (dispatch: Dispatch<AppActionTypes>) => {
    return {
        onLogIn: (user: User) => {
            dispatch(actions.logIn(user))
        },
        loadRooms: (rooms: Room[]) => {
            dispatch(actions.loadRooms(rooms))
        },
        changeRooms: (room: Room) => {
            dispatch(actions.changeRoom(room))
        },
        loadMessages: (messages: Message[]) => {
            dispatch(actions.loadMessages(messages))
        }
    }
}
type Props = ReturnType<typeof mapStateToProps> & ReturnType<typeof mapDispatchToProps>


class ReduxTester extends Component<Props, AppState> {
    componentDidMount() {
        console.log(this.props)
        let curRoom: Room = { id: 7, name: "test7", hasJoined: true, hasNotification: false, isDM: false }
        this.props.onLogIn({ username: "test", id: 123, jwtToken: "hey" })
        this.props.loadRooms([
            { id: 1, name: "test1", hasJoined: true, hasNotification: true, isDM: false },
            { id: 2, name: "test2", hasJoined: false, hasNotification: true, isDM: false },
            { id: 3, name: "test3", hasJoined: true, hasNotification: false, isDM: false },
            { id: 4, name: "test4", hasJoined: false, hasNotification: false, isDM: false },
            { id: 5, name: "test5", hasJoined: true, hasNotification: true, isDM: true },
            { id: 6, name: "test6", hasJoined: true, hasNotification: true, isDM: true },
            curRoom
        ])
        this.props.changeRooms(curRoom)
        this.props.loadMessages([
            { username: "Alice", body: "hello1", time: new Date() },
            { username: "Bob", body: "hello2", time: new Date() },
            { username: "Calvin", body: "hello3", time: new Date() },
            { username: "Bob", body: "hello4", time: new Date() },
            { username: "Alice", body: "hello5", time: new Date() },
            { username: "Bob", body: "hello6", time: new Date() },
            { username: "Calvin", body: "hello7", time: new Date() },
        ])

    }

    render() {
        return (
            <div></div>
        )
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(ReduxTester)