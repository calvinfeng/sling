import React, { Component, Dispatch } from 'react';
import { connect } from 'react-redux'
import { AppState } from '../store'
import { logIn, changeRoom } from '../actions'
import { AppActionTypes } from '../actions/types'

import { User, Room, Message } from '../types'

const mapStateToProps = (state: AppState) => state
const mapDispatchToProps = (dispatch: Dispatch<AppActionTypes>) => {
    return {
        onLogIn: (user: User) => {
            dispatch(logIn(user))
        },
        changeRoom: (room: Room) => {
            dispatch(changeRoom(room))
        }
    }
}
type Props = ReturnType<typeof mapStateToProps> & ReturnType<typeof mapDispatchToProps>


class ReduxTester extends Component<Props, AppState> {
    componentDidMount() {
        console.log(this.props)
        this.props.onLogIn({username: "test", id: 123, jwtToken: "hey"})
        this.props.changeRoom({id: 123, name: "test", hasJoined: true, hasNotification: true, isDM: false})
    }

    render() {
        return (
            <div></div>
        )
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(ReduxTester)