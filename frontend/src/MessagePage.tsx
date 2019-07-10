import * as React from 'react';
import { connect } from 'react-redux'
import axios, { AxiosResponse } from 'axios'

import './MessagePage.css';
import SideBar from './components/SideBar';
import DisplayWindow from './components/DisplayWindow';
import InputBox from './components/InputBox';
import * as actions from './actions'
import { AppActionTypes } from './actions/types'

import { AppState } from './store'
import { User, Room, Message } from './types'

interface MessagePageState {
    inputEnabled: boolean
}

const initialState: MessagePageState = {
    inputEnabled: true
}

type OwnProps = {
    setLoggedOut: Function
}

const mapStateToProps = (state: AppState, ownProps: OwnProps) => ({ ...state, ...ownProps })
const mapDispatchToProps = (dispatch: React.Dispatch<AppActionTypes>) => {
    return {
        onLogOut: () => {
            dispatch(actions.logOut())
        },
        onChangeRoom: (room: Room) => {
            console.log(room)
            dispatch(actions.changeRoom(room))
        }
    }
}

type Props = ReturnType<typeof mapStateToProps> & ReturnType<typeof mapDispatchToProps>

class MessagePage extends React.Component<Props, MessagePageState> {
    // private msgWebsocket: WebSocket
    // private actWebsocket: WebSocket
    readonly state: MessagePageState = initialState

    componentDidMount() {
        // TODO: Fetch initial states

        // // Set up websocket handlers
        // this.msgWebsocket = new WebSocket("ws://localhost:8000/streams/messages")
        // this.actWebsocket = new WebSocket("ws://localhost:8000/streams/actions")
        // this.msgWebsocket.onopen = this.handleMsgWebsocketOpen
        // this.msgWebsocket.onclose = this.handleMsgWebsocketClose
        // this.msgWebsocket.onmessage = this.handleMsgWebsocketMessage
        // this.msgWebsocket.onerror = this.handleMsgWebsocketError
        // this.actWebsocket.onopen = this.handleActWebsocketOpen
        // this.actWebsocket.onclose = this.handleActWebsocketClose
        // this.actWebsocket.onerror = this.handleActWebsocketError
        // this.actWebsocket.onmessage = this.handleActWebsocketMessage
    }

    // handleMsgWebsocketOpen = (ev: Event) => {

    // }

    // handleMsgWebsocketClose = (ev:CloseEvent) => {

    // }

    // handleMsgWebsocketMessage = (mev:MessageEvent) => {

    // }

    // handleMsgWebsocketError = (ev:Event) => {

    // }

    // handleActWebsocketOpen = (ev: Event) => {

    // }

    // handleActWebsocketClose = (ev:CloseEvent) => {

    // }

    // handleActWebsocketError = (ev:Event) => {

    // }

    // handleActWebsocketMessage = (mev:MessageEvent) => {

    // }

    sendMessage(body: String) {
        console.log(`sending ${body}`)
        // TODO: send message
        this.setState({ inputEnabled: false })

    }

    changeRoom(nextRoom: Room) {
        this.props.onChangeRoom(nextRoom)
    }

    render() {
        console.log(this.props)
        return (
            <div className="App">
                <div className="left-div">
                    <SideBar
                        curUser={this.props.curUser!}
                        curRoom={this.props.curRoom}
                        rooms={this.props.rooms}
                        users={this.props.users}
                        logOut={this.props.setLoggedOut}

                        changeRoom={(room: Room) => this.changeRoom(room)}
                    />
                </div>
                <div className="right-div">
                    <div className="messages">
                        <DisplayWindow
                            curRoom={this.props.curRoom!}
                            messages={this.props.messages}
                        />
                    </div>
                    <div className="inputs">
                        <InputBox
                            sendMessage={(body: String) => this.sendMessage(body)}
                            enabled={this.state.inputEnabled}
                        />
                    </div>
                </div>
            </div>
        );
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(MessagePage);
