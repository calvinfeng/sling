import * as React from 'react';
import { Component } from 'react';
import { User } from '../../types';
import { Button } from '@material-ui/core'
import './styles.scss';

type SideBarHeaderProps = {
    curUser: User | null
    logOut: Function
}

export default class SideBarHeader extends Component<SideBarHeaderProps, {}> {
    render() {
        return (
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
        )
    }
}

