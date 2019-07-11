import * as React from 'react';
import { Component } from 'react';
import { Room, User } from '../types'
import { Button } from '@material-ui/core'
import './component.css'
import curUser from '../reducers/curUser';

type SideBarProps = {
    curUser: User
    curRoom: Room | null
    users: User[]
    rooms: Room[]

    logOut: Function
    changeRoom: Function
    startDM: Function
    joinRoom: Function
}

type SideBarState = {
    displayMoreChannel: boolean
    displayMoreUser: boolean
}


class SideBar extends Component<SideBarProps, SideBarState> {
    constructor(props: SideBarProps) {
        super(props);
        this.state = {
            displayMoreChannel: false,
            displayMoreUser: false
        };
    }


    shouldComponentUpdate(nextProps: SideBarProps): boolean {
        return true
    }

    handleDisplayMoreChannel = () => {
        this.setState({
            displayMoreChannel: !this.state.displayMoreChannel,
        })
    }

    handleDisplayMoreUser = () => {
        this.setState({
            displayMoreUser: !this.state.displayMoreUser,
        })
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

    renderUserList = () => {
        let dms = this.props.rooms
            .filter(room => room.isDM)
            .map(room => this.findDirectMsgName(room.name))
        
        // Filter out current user and users with active DMs
        return this.props.users.filter((user) =>
            user.username !== this.props.curUser.username &&
            !dms.includes(user.username)
        ).map((user) =>
            <li
                className="SBhoverable SBdmuser"
                key={user.username}
                onClick={(e) => this.props.startDM(user)}
            >
                {user.username}
            </li>
        )
    }

    renderChannels = (hasJoined: boolean, isDM: boolean) => {
        return this.props.rooms.filter((room): boolean =>
            room.hasJoined === hasJoined &&
            room.isDM === isDM
        ).map((room) =>
            <li
                className={`SBhoverable ${this.getClassName(room)}`}
                key={room.id}
                onClick={hasJoined ?
                    (e) => this.props.changeRoom(room) :
                    (e) => this.props.joinRoom(room)
                }
            >
                {room.name}
            </li>
        );
    }

    render() {
        let listItems = this.renderChannels(true, false)
        let userItems = this.renderChannels(true, true)
        let unjoinedChannels = this.renderChannels(false, false)
        let noDMUsers = this.renderUserList()

        let moreChannel = <label onClick={this.handleDisplayMoreChannel} className="SBmore SBhoverable">
            {this.state.displayMoreChannel ? '-' : '+'} More Channels
        </label>
        let moreUser = <label onClick={this.handleDisplayMoreUser} className="SBmore SBhoverable">
            {this.state.displayMoreUser ? '-' : '+'} More People
        </label>

        return (
            <div>
                <div className="SBcurUser">
                    <label>{this.props.curUser && this.props.curUser!.username}'s sling</label>
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
                        {listItems.length > 0 ?
                            <ul className="SBlist">{listItems}</ul> :
                            <div className="SBnone">None</div>
                        }

                        <div className="SBMoreUser">
                            {moreChannel}
                            {this.state.displayMoreChannel && (
                                unjoinedChannels.length > 0 ?
                                    <ul className="SBlist">{unjoinedChannels}</ul> :
                                    <div className="SBnone">None</div>
                            )}
                        </div>

                        <label className="SBlabel">Direct Messages</label>
                        {userItems.length > 0 ?
                            <ul className="SBlist">{userItems}</ul> :
                            <div className="SBnone">None</div>
                        }
                    </div>
                    <div className="SBMoreUser">
                        {moreUser}
                        {this.state.displayMoreUser && (
                            noDMUsers.length > 0 ?
                                <ul className="SBlist">{noDMUsers}</ul> :
                                <div className="SBnone">None</div>
                        )}
                    </div>
                </div>
            </div>
        );
    }
}

export default SideBar;