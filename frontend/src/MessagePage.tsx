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
    error: string
    loading: boolean
}

const initialState: MessagePageState = {
    inputEnabled: false,
    error: '',
    loading: true
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

        onLoadMessages: (messages: Message[]) => {
            dispatch(actions.loadMessages(messages))
        },
        onLoadRooms: (rooms: Room[]) => {
            dispatch(actions.loadRooms(rooms))
        },
        onLoadUsers: (users: User[]) => {
            dispatch(actions.loadUsers(users))
        },

        onNewMessage: (message: Message) => {
            dispatch(actions.newMessage(message))
        },
        onNewRoom: (room: Room) => {
            dispatch(actions.newRoom(room))
        },
        onNewUser: (user: User) => {
            dispatch(actions.newUser(user))
        },

        onJoinRoom: (room: Room) => {
            dispatch(actions.joinRoom(room))
        },
        onChangeRoom: (room: Room) => {
            dispatch(actions.changeRoom(room))
        },
    }
}

type Props = ReturnType<typeof mapStateToProps> & ReturnType<typeof mapDispatchToProps>

class MessagePage extends React.Component<Props, MessagePageState> {
    // private msgWebsocket: WebSocket
    // private actWebsocket: WebSocket
    readonly state: MessagePageState = initialState
    private messagesEnd = React.createRef<HTMLDivElement>()

    componentDidMount() {
        // Fetch users
        let usersPromise = axios.get('api/users/', {
            headers: {
                'Token': localStorage.getItem('jwt_token')
            }
        }).then((res: AxiosResponse) => {
            this.props.onLoadUsers(res.data.map((user: any): User => ({
              username: user.name,
              id: user.id
            })))
        })

        // Fetch rooms
        let roomsPromise = axios.get('api/rooms', {
            headers: {
                'Token': localStorage.getItem('jwt_token')
            }
        }).then((res: AxiosResponse) => {
            this.props.onLoadRooms(res.data.map((room: any): Room => ({
               id: room.id,
               name: room.name,

               // TODO: determine whether the user has joined and has notif
               hasJoined: false,
               hasNotification: false,
               isDM: false
            })))
        })

        Promise.all([usersPromise, roomsPromise]).catch((err) => {
            console.log(err)
            this.setState({ error: 'Could not fetch rooms.' })
        }).finally(() => {
            console.log(this.state)
            this.setState({ loading: false })
        })

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
        this.setState({ inputEnabled: true })
        this.props.onChangeRoom(nextRoom)

        // TODO: load next room's messages
    }

    startDM(user: User) {
        console.log(user)

        //TODO: start DM
    }

    scrollToBottom = () => {
        this.messagesEnd.current!.scrollIntoView({ behavior: "smooth" });
    }

    render() {
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
                        joinRoom={(room: Room) => this.props.onJoinRoom(room)}
                        startDM={(user: User) => this.startDM(user)}
                    />
                </div>
                <div className="right-div">
                    <div className="messages">
                        <DisplayWindow
                            curRoom={this.props.curRoom!}
                            messages={this.props.messages}
                        />
                        <div style={{ float: "left", clear: "both" }} ref={this.messagesEnd}></div>
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
