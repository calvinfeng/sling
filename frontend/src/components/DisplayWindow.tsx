import * as React from 'react';
import { Component } from 'react';
import { Message, Room } from '../types';
import './component.css';

import moment from 'moment';

type DisplayWindowProps = {
    messages: Message[]
    curRoom: Room
}

type DisplayWindowState = {

}

class DisplayWindow extends Component<DisplayWindowProps, DisplayWindowState> {
    render() {
        const displayMessages = this.props.messages.map((msg) =>
            <div key={msg.username + msg.time} className="DWmessage">
                <div>
                    <span className="DWusername">{msg.username} </span>
                    <span className="DWtime">{moment(msg.time).fromNow()}</span>
                </div>
                <div className="DWmessagebody">{msg.body}</div>
            </div>
        );
        return (
            <div>
                {this.props.curRoom ?
                    <div>
                        <div className="DWlabel">
                            <label>#{this.props.curRoom.name}</label>
                        </div>
                        <div className="DWmessages">
                            {displayMessages}
                        </div>
                    </div> :
                    "No room selected."}
            </div>
        );
    }
}

export default DisplayWindow;