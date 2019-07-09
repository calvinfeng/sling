import React, { Component, ChangeEvent, FormEvent } from 'react';

import {
    Container,
    TextField,
    Button,
    makeStyles
} from '@material-ui/core'

const initialState = {
    login: true,
    username: '',
    password: '',
    email: '',
    error: '',
    loading: false
}
type LoginState = {
    login: boolean,
    username: string,
    password: string,
    email: string,
    error: string,
    loading: boolean
}

export default class Login extends Component<{}, LoginState> {
    readonly state: LoginState = initialState

    handleChange(value: string, field: string) {
        this.setState(prevState => ({
            ...prevState,
            [field]: value
        }))
    }

    handleSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        if (!this.validateInfo()) {
            return
        }

        this.setState({ error: '', loading: true })
        if (this.state.login) {
            this.tryLogin()
        } else {
            this.tryRegister()
        }
    }

    validateInfo(): boolean {
        if (this.state.username.length <= 0) {
            this.setState({ error: 'No username provided.' })
            return false
        }

        if (!this.state.login && this.state.email.length <= 0) {
            this.setState({ error: 'No email provided.' })
            return false
        }

        if (this.state.password.length <= 0) {
            this.setState({ error: 'No password provided.' })
            return false
        }

        return true
    }

    toggleLogin() {
        this.setState({ login: !this.state.login })
    }

    tryLogin() {
        // make server request here

        this.setState({ loading: false })
    }

    tryRegister() {
        // make server request here
        
        this.setState({ loading: false })
    }

    render() {
        return (
            <Container maxWidth="sm">
                <h1>{this.state.login ? 'Login' : 'Register'}</h1>

                <div>
                    <Button
                        disabled={this.state.loading}
                        variant='contained'
                        onClick={() => this.toggleLogin()}
                    >
                        {this.state.login ? 'Register' : 'Login'}
                    </Button>
                </div>

                <br />

                <form onSubmit={(e) => this.handleSubmit(e)}>
                    <div>
                        <TextField
                            disabled={this.state.loading}
                            id='username'
                            label='Username'
                            value={this.state.username}
                            onChange={e => this.handleChange(e.currentTarget.value, 'username')}
                        />
                    </div>

                    {!this.state.login &&
                        <div>
                            <TextField
                                disabled={this.state.loading}
                                id='email'
                                label='Email'
                                value={this.state.email}
                                onChange={e => this.handleChange(e.currentTarget.value, 'email')}
                            />
                        </div>
                    }

                    <div>
                        <TextField
                            disabled={this.state.loading}
                            id='password'
                            label='Password'
                            type='password'
                            value={this.state.password}
                            onChange={e => this.handleChange(e.currentTarget.value, 'password')}
                        />
                    </div>

                    <br />

                    <div style={{ color: 'red' }}>{this.state.error}</div>

                    <br />

                    <Button
                        disabled={this.state.loading}
                        variant='contained'
                        color='primary'
                        type='submit'
                        value='Submit'
                    >
                        Submit
                    </Button>
                </form>
            </Container>
        )
    }
}