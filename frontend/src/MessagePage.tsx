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
import curRoom from './reducers/curRoom';

type MessagePageState = {
    inputEnabled: boolean
    error: string
    loading: boolean
    connectedToMsgSocket: boolean
    connectedToActSocket: boolean
}

const initialState: MessagePageState = {
    inputEnabled: false, // don't enable until messages successfully fetched
    error: '',
    loading: true,
    connectedToMsgSocket: false,
    connectedToActSocket: false,
}

type OwnProps = {
    setLoggedOut: Function
}

const mapStateToProps = (state: AppState, ownProps: OwnProps) => ({ ...state, ...ownProps })
const mapDispatchToProps = (dispatch: React.Dispatch<AppActionTypes>) => {
    return {
        onLogIn: (user: User) => {
            dispatch(actions.logIn(user))
        },
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
    private msgWebsocket!: WebSocket
    private actWebsocket!: WebSocket 
    readonly state: MessagePageState = initialState
    private messagesEnd = React.createRef<HTMLDivElement>()

    componentDidMount() {
        let token = localStorage.getItem('jwt_token')
        if (!token || token.length === 0) {
            this.props.setLoggedOut()
        }

        // Get current user from token
        let curUserPromise = axios.get('api/users/current', {
            headers: { 'Token': token }
        }).then((res: AxiosResponse) => {
            let user: User = {
                id: res.data.id,
                username: res.data.name,
                jwtToken: res.data.jwt_token
            }
            this.props.onLogIn(user)
        })

        // Fetch users
        let usersPromise = axios.get('api/users/', {
            headers: {
                'Token': token
            }
        }).then((res: AxiosResponse) => {
            this.props.onLoadUsers(res.data.map((user: any): User => ({
                username: user.name,
                id: user.id,
                jwtToken: null // can't get other peoples' tokens
            })))
        })

        // Fetch rooms
        let roomsPromise = axios.get('api/rooms', {
            headers: {
                'Token': token
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

        Promise.all([curUserPromise, usersPromise, roomsPromise]).catch((err) => {
            console.log(err)

            // Force user back to login page
            this.props.setLoggedOut()
        }).finally(() => {
            console.log(this.state)
            this.setState({ loading: false })
            console.log("before sockets")

            // // Set up websocket handlers
            this.msgWebsocket = new WebSocket("ws://localhost:8888/api/stream/messages")
            this.actWebsocket = new WebSocket("ws://localhost:8888/api/stream/actions")
            console.log("after sockets")
            console.log("after send")
 
            this.msgWebsocket.onopen = this.handleMsgWebsocketOpen
            this.msgWebsocket.onclose = this.handleMsgWebsocketClose
            // this.msgWebsocket.onmessage = this.handleMsgWebsocketMessage
            this.msgWebsocket.onerror = this.handleMsgWebsocketError
            this.actWebsocket.onopen = this.handleActWebsocketOpen
            this.actWebsocket.onclose = this.handleActWebsocketClose
            this.actWebsocket.onerror = this.handleActWebsocketError
            // this.actWebsocket.onmessage = this.handleActWebsocketMessage
        })

    }

    handleMsgWebsocketOpen = (ev: Event) => {
        if (this.props.curUser != null){
            var token = this.props.curUser.jwtToken
            this.msgWebsocket.send(JSON.stringify({jwt_token: token}));
        }
        else {
            console.log("curUser is null")
        }
        this.setState({connectedToMsgSocket: true,});
    }

    handleMsgWebsocketClose = (ev:CloseEvent) => {
        this.setState({connectedToMsgSocket: false,});
    }

    // handleMsgWebsocketMessage = (mev:MessageEvent) => {
    //     const msgResponsePayload = JSON.parse(mev.data);
    //     if (msgResponsePayload.messageType === "new_message") {  
    //         this.props.onNewMessage({
    //             username: msgResponsePayload.uerName,
    //             time: msgResponsePayload.time,
    //             body: msgResponsePayload.body,
    //         });
    //     } else if (msgResponsePayload.messageType === "notification"){
    //         this.props.onMarkUnread(msgResponsePayload.roomID);
    //     } else {
    //         console.log("undefined type")
    //     }
    // }

    handleMsgWebsocketError = (ev:Event) => {
        this.setState({ error: "encountered message websocket error" + ev })
    }

    handleActWebsocketOpen = (ev: Event) => {
        if (this.props.curUser != null){
            var token = this.props.curUser.jwtToken
            this.actWebsocket.send(JSON.stringify({jwt_token: token}));
        }
        else {
            console.log("curUser is null")
        }
        this.setState({connectedToActSocket: true,});
    }

    handleActWebsocketClose = (ev:CloseEvent) => {
        this.setState({connectedToActSocket: false,});
    }

    handleActWebsocketError = (ev:Event) => {
        this.setState({ error: "encountered action websocket error" + ev })
    }

    // handleActWebsocketMessage = (mev:MessageEvent) => {
    //     const actResponsePayload = JSON.parse(mev.data);
    //     if (actResponsePayload.actionType === "message_history") {
    //         const msgs = actResponsePayload.messageHistory;
    //         let messages = [];
    //         for (let i = 0; i < msgs.length ; i++) {
    //             messages.push({
    //                 username: msgs[i].senderName,
    //                 time: msgs[i].time,
    //                 body: msgs[i].body,
    //             })
    //         }       
    //         this.props.onLoadMessages(messages);
    //     } else if (actResponsePayload.actionType === "create_dm") {
    //         this.props.onNewRoom({
    //             id: actResponsePayload.roomID,
    //             name: actResponsePayload.roomName,
    //             hasJoined: true,
    //             hasNotification: false,
    //             isDM: true,
    //         })
    //     } else if (actResponsePayload.actionType === "new_user") {
    //         this.props.onNewUser({
    //             username: actResponsePayload.userName,
    //             id: actResponsePayload.userID,
    //         })
    //     } else if (actResponsePayload.actionType === "new_room") {
    //         this.props.onNewRoom({
    //             id: actResponsePayload.roomID,
    //             name: actResponsePayload.roomName,
    //             hasJoined: true,
    //             hasNotification: false,
    //             isDM: false,
    //         })
    //     } else {
    //         console.log("undefined type");
    //     }

    //}

    sendMessage(body: String) {
        console.log(`sending ${body}`)
        // TODO: send message
        this.setState({ inputEnabled: false })

    }

    changeRoom(nextRoom: Room) {
        if (this.props.curRoom && nextRoom.id === this.props.curRoom.id) {
            return
        }

        this.props.onChangeRoom(nextRoom)

        // TODO: load next room's messages
        // var actionPayload = {
        //     actionType: ,
        //     userID
        //     roomID
        //     newRoomID
        //     dmUserID
        //     newR
        //         ActionType  string `json:"actionType"`
        //         UserID      uint   `json:"userID"`
        //         RoomID      uint   `json:"roomID"`
        //         NewRoomID   uint   `json:"newRoomID"`
        //         DMUserID    uint   `json:"dmUserID"`
        //         NewRoomName string `json:"newRoomName"`
        //     }
        // }
        // this.actWebsocket.send(JSON.stringify)
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
