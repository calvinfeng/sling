import * as React from 'react';
import { Component } from 'react';
import { Room, User } from '../../types'
import { IconButton } from '@material-ui/core'
import SideBarHeader from './SideBarHeader'
import AddBox from '@material-ui/icons/AddBox'
import ChannelList from './ChannelList'
import './styles.scss';

type SideBarProps = {
    curUser: User
    curRoom: Room | null
    users: User[]
    rooms: Room[]

    logOut: Function
    changeRoom: Function
    startDM: Function
    createRoom: Function
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
        return true // TODO: proper update logic
    }

    createChannel() {
        let name = prompt('Enter a new channel name.')
        if (name === null) {
            return
        }
        name = name.trim()

        while (!this.validateChannelName(name!)) {
            name = prompt('Invalid name, try again.')
            if (name === null) {
                return
            }
            name = name.trim()
        }

        this.props.createRoom(name)
    }

    validateChannelName(name: string) {
        return name.length > 0 &&
            !name.includes(' ') &&
            !name.includes('~')
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
        if (names.length < 2) return name;
        if (names[0] !== this.props.curUser.username) return names[0];
        return names[1];
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

    filterChannels = (hasJoined: boolean, isDM: boolean) => {
        return this.props.rooms.filter((room): boolean =>
            room.hasJoined === hasJoined &&
            room.isDM === isDM
        ).map((room) => {
            if (isDM) {
                room.name = this.findDirectMsgName(room.name)
            }

            return room
        })
    }

    render() {
        let joinedChannels = this.filterChannels(true, false)
        let joinedDMs = this.filterChannels(true, true)
        let unjoinedChannels = this.filterChannels(false, false)
        let unjoinedDMs = this.renderUserList()

        let moreChannel = <label onClick={this.handleDisplayMoreChannel} className="SBmore SBhoverable">
            {this.state.displayMoreChannel ? '-' : '+'} More Channels
        </label>
        let moreUser = <label onClick={this.handleDisplayMoreUser} className="SBmore SBhoverable">
            {this.state.displayMoreUser ? '-' : '+'} More People
        </label>

        return (
            <div>
                <SideBarHeader
                    curUser={this.props.curUser}
                    logOut={this.props.logOut}
                />
                <div >
                    <div className="SBRooms">
                        <label className="SBlabel">Channels</label>
                        <IconButton
                            onClick={() => this.createChannel()}
                            className="SBaddbtn"
                            color="primary"
                            size="small"
                        >
                            <AddBox />
                        </IconButton>
                        <ChannelList
                            data={joinedChannels}
                            curRoom={this.props.curRoom}
                            changeRoom={this.props.changeRoom}
                            joinRoom={this.props.joinRoom}
                        />

                        <div className="SBMoreUser">
                            {moreChannel}
                            {this.state.displayMoreChannel &&
                                <ChannelList
                                    data={unjoinedChannels}
                                    curRoom={this.props.curRoom}
                                    changeRoom={this.props.changeRoom}
                                    joinRoom={this.props.joinRoom}
                                />
                            }
                        </div>

                        <label className="SBlabel">Direct Messages</label>
                        <ChannelList
                            data={joinedDMs}
                            curRoom={this.props.curRoom}
                            changeRoom={this.props.changeRoom}
                            joinRoom={this.props.joinRoom}
                        />
                    </div>
                    <div className="SBMoreUser">
                        {moreUser}
                        {this.state.displayMoreUser && (
                            unjoinedDMs.length > 0 ?
                                <ul className="SBlist">{unjoinedDMs}</ul> :
                                <div className="SBnone">None</div>
                        )}
                    </div>
                </div>
            </div>
        );
    }
}

export default SideBar;