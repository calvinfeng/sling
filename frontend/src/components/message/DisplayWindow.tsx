import * as React from 'react';
import { Component } from 'react';
import { Message, Room } from '../../types';
import ChatMessages from './ChatMessages'
import './styles.scss';

type DisplayWindowProps = {
    messages: Message[]
    curRoom: Room
}

class DisplayWindow extends Component<DisplayWindowProps, {}> {
    render() {
        return (
            <div>
                {this.props.curRoom ?
                    <div>
                        <div className="DWlabel">
                            <label>{
                                (this.props.curRoom.isDM ?
                                    'Direct message with ' :
                                    '#') + this.props.curRoom.name
                            }</label>
                        </div>
                        <div className="DWmessages">
                            <ChatMessages messages={this.props.messages} />
                        </div>
                    </div> :
                    "No room selected."}
            </div>
        );
    }
}

export default DisplayWindow;