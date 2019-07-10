import * as React from 'react';
import { Component } from 'react';
import { Room, User } from '../types'
import { Button } from '@material-ui/core'
import './component.css'

export interface SideBarProps {
    curUser: User
    curRoom: Room | null
    users: User[]
    rooms: Room[]

    logOut: Function
    changeRoom: Function
}

export interface SideBarState {
    displayMore: boolean
}


class SideBar extends Component<SideBarProps, SideBarState> {
    constructor(props: SideBarProps) {
        super(props);
        this.state = {
            displayMore: false,
        };
    }

    handleDisplayMoreUser = () => {
        this.setState({
            displayMore: !this.state.displayMore,
        })
    }

    hasJoined = (room: Room) => {
        return room.hasJoined;
    }

    isDirectMsg = (room: Room) => {
        return room.isDM;
    }

    isNotDirectMsg = (room: Room) => {
        return !room.isDM;
    }

    findDirectMsgName = (name: string) => {
        const names: string[] = name.split("~")
        if (names[0] !== this.props.curUser.username) return names[0];
        return names[1];
    }

    getClassName = (room: Room) => {
        if (this.props.curRoom && room.id === this.props.curRoom!.id) return "curroom";
        if (room.hasNotification) return "notification";
        return "normal";
    }

    render() {
        const { hasJoined, isDirectMsg, findDirectMsgName, getClassName, isNotDirectMsg } = this;

        const listItems = this.props.rooms.filter(hasJoined).filter(isNotDirectMsg).map((room) =>
            <li
                className={`SBroom ${getClassName(room)}`}
                key={room.id}
                onClick={(e) => this.props.changeRoom(room)}
            >
                {room.name}
            </li>
        );

        const userItems = this.props.rooms.filter(isDirectMsg).map((room) =>
            <li
                className={`SBroom ${getClassName(room)}`}
                key={room.id}
                onClick={(e) => this.props.changeRoom(room)}
            >
                {findDirectMsgName(room.name)}
            </li>
        );

        let moreUser = (<label onClick={this.handleDisplayMoreUser} className="SBlabel">+ More People</label>);
        if (this.state.displayMore) {
            moreUser = (
                <div>
                    <label onClick={this.handleDisplayMoreUser} className="SBlabel">+ More People</label>
                    <ul className="SBlist">
                        {this.props.users.map((user) =>
                            <li key={user.id}> {user.username} </li>
                        )}
                    </ul>
                </div>
            );
        }
        return (
            <div>
                <div className="SBcurUser">
                    <label>{this.props.curUser!.username}'s sling</label>
                    <Button
                        onClick={() => this.props.logOut()}
                        className="SBlogout"
                        color="secondary"
                    >
                        Log out
                    </Button>
                </div>
                <div >
                    <div className="SBRooms">
                        <label className="SBlabel">Channels</label>
                        <ul className="SBlist">{listItems}</ul>
                        <label className="SBlabel">Direct Messages</label>
                        <ul className="SBlist">{userItems}</ul>
                    </div>
                    <div className="SBMoreUser">
                        {moreUser}
                    </div>
                </div>
            </div>
        );
    }
}

export default SideBar;