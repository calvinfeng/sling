import * as React from 'react';
import { Component } from 'react';
import { RoomType, UserType } from '../store/sling/index'
import './component.css'

export interface SideBarProps {
    
}
 
export interface SideBarState {
    curUserName: string,
    curRoom: RoomType,
    rooms: RoomType[],
    users: UserType[],
    displayMore: boolean,
}

 
class  SideBar extends Component<SideBarProps, SideBarState> {
    constructor(props: SideBarProps) {
        super(props);
        this.state = {
            curUserName: "a",
            curRoom: {
                roomID: 1,
                roomName: "test",
                hasJoined: true,
                hasNotification: true,
                isDirectMessage: false,
            },
            rooms:[
                {
                    roomID: 3,
                    roomName: "Room1",
                    hasJoined:true,
                    hasNotification:true,
                    isDirectMessage:false,
                },
                {
                    roomID: 2,
                    roomName: "a~b",
                    hasJoined:true,
                    hasNotification:false,
                    isDirectMessage:true,
                },
                {
                    roomID: 1,
                    roomName: "test",
                    hasJoined: true,
                    hasNotification: true,
                    isDirectMessage: false,
                }
            ],
            users:[
                {
                    userID: 1,
                    userName: "Calvin"
                },
            ],
            displayMore: false,
        };
    }

    handleDisplayMoreUser = () => {
        this.setState ({
            displayMore: !this.state.displayMore,
        })
    }

    hasJoined = (room: RoomType) =>{
        return room.hasJoined;
    }

    isDirectMsg = (room: RoomType) => {
        return room.isDirectMessage;
    }

    isNotDirectMsg = (room: RoomType) => {
        return !room.isDirectMessage;
    }

    findDirectMsgName =(roomName: string) => {
        const names:string[] = roomName.split("~")
        if (names[0] !== this.state.curUserName) return names[0];
        return names[1];
    }

    getClassName = ( room: RoomType ) => {
        if (room.roomID === this.state.curRoom.roomID) return "curroom";
        if (room.hasNotification) return "notification";
        return "normal";
    }

    render() { 
        const { hasJoined, isDirectMsg, findDirectMsgName, getClassName, isNotDirectMsg} = this;

        const listItems = this.state.rooms.filter(hasJoined).filter(isNotDirectMsg).map ((room) =>       
            <li className={getClassName(room)} key={room.roomID}>{room.roomName}</li>        
        );

        const userItems = this.state.rooms.filter(isDirectMsg).map ((room) =>
            <li className={getClassName(room)} key={room.roomID}>{findDirectMsgName(room.roomName)}</li>
        );

        let moreUser = (<label onClick={this.handleDisplayMoreUser } className="SBlabel">+ More People</label>);
        if (this.state.displayMore) {
            moreUser = (
                <div>
                    <label onClick={this.handleDisplayMoreUser } className="SBlabel">+ More People</label>
                    <ul className="SBlist">
                        {this.state.users.map ((user) =>
                            <li key={user.userID}> {user.userName} </li>
                        )}
                    </ul>
                </div>    
            );
        }
        return (
            <div>
                <div className="SBcurUser">
                    <label>{this.state.curUserName}'s sling</label>  
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