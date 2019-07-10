import * as React from 'react';
import './MessagePage.css';
import SideBar from './components/SideBar';
import DisplayWindow from './components/DisplayWindow';
import InputBox from './components/InputBox';

export interface MessagePageProps {
    
}
 
export interface MessagePageState {
    
}
 
class MessagePage extends React.Component<MessagePageProps, MessagePageState> {
    // private msgWebsocket: WebSocket
    // private actWebsocket: WebSocket

    // componentDidMount() {
    //     this.msgWebsocket = new WebSocket("ws://localhost:8000/streams/messages")
    //     this.actWebsocket = new WebSocket("ws://localhost:8000/streams/actions")
    //     this.msgWebsocket.onopen = this.handleMsgWebsocketOpen
    //     this.msgWebsocket.onclose = this.handleMsgWebsocketClose
    //     this.msgWebsocket.onmessage = this.handleMsgWebsocketMessage
    //     this.msgWebsocket.onerror = this.handleMsgWebsocketError
    //     this.actWebsocket.onopen = this.handleActWebsocketOpen
    //     this.actWebsocket.onclose = this.handleActWebsocketClose
    //     this.actWebsocket.onerror = this.handleActWebsocketError
    //     this.actWebsocket.onmessage = this.handleActWebsocketMessage
    // }

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

    
    render() { 
        return (  
            <div className="App">
                <div className="left-div">
                <SideBar />
                </div>
                <div className="right-div">
                <div className="messages">
                    <DisplayWindow />
                </div>
                <div className="inputs">
                    <InputBox />
                </div>
                </div>
            </div>
        );
    }
}
 
export default MessagePage;
