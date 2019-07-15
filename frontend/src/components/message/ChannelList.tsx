import * as React from 'react';
import { Component } from 'react';
import { Room } from '../../types';
import './styles.scss';

type ChannelListProps = {
    data: Room[]
    curRoom: Room | null
    changeRoom: Function
    joinRoom: Function
}

export default class DisplayWindow extends Component<ChannelListProps, {}> {
    getClassName = (room: Room) => {
        if (this.props.curRoom && room.id === this.props.curRoom!.id) return "curroom";
        if (room.hasNotification) return "notification";
        return "normal";
    }

    render() {
        let list = this.props.data.map((room) => {
            return <li
                className={`SBhoverable ${this.getClassName(room)}`}
                key={room.id}
                onClick={room.hasJoined ?
                    (e) => this.props.changeRoom(room) :
                    (e) => this.props.joinRoom(room)
                }
            >
                {room.name}
            </li>
        })

        return list.length > 0 ?
            <ul className="SBlist">{list}</ul> :
            <div className="SBnone">None</div>
    }
}