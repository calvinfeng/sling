import * as React from 'react';
import { Component } from 'react';
import { Message } from '../../types';
import './styles.scss';

import moment from 'moment';

type ChannelListProps = {
    messages: Message[]
}

export default class ChatMessage extends Component<ChannelListProps, {}> {
    render() {
        let messages = this.props.messages.map((msg) =>
            <div key={msg.username + msg.time} className="DWmessage">
                <div>
                    <span className="DWusername">{msg.username} </span>
                    <span className="DWtime">{moment(msg.time).fromNow()}</span>
                </div>
                <div className="DWmessagebody">{msg.body}</div>
            </div>
        );

        return messages
    }
}