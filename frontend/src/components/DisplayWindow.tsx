import * as React from 'react';
import { Component } from 'react';
import { MessageType } from '../store/sling/index';
import './component.css';

export interface DisplayWindowProps {
    
}
 
export interface DisplayWindowState {
    roomName: string;
    messages: MessageType[];
}
 
class DisplayWindow extends Component<DisplayWindowProps, DisplayWindowState> {
    constructor(props: DisplayWindowProps) {
        super(props);
        this.state = {
            roomName: "Ah ha!",
            messages: [
            {
                msgID: 1,
                userID:1,
                userName: "Bob",
                time: "12:03 am",
                body:"Hi",

            },
            {
                msgID: 2,
                userID: 2,
                userName: "Alice",
                time: "12:04 am",
                body:"Hi",

            },
        ],
        };
    }
   
    render() { 
        const displayMessages = this.state.messages.map((msg) =>
            <div key={msg.msgID} className="DWmessage">
                <div>{msg.userName} {msg.time} </div>
                <div>{msg.body}</div>
            </div>
        );
        return ( 
            <div>
                <div className="DWlabel"> 
                    <label>Channel: {this.state.roomName}</label>
                </div>
                <div>
                    {displayMessages}
                </div>
            </div>
         );
    }
}
 
export default DisplayWindow;