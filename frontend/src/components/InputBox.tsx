import * as React from 'react';
import './component.css';
import {
    TextField,
    Button
} from '@material-ui/core'

export interface InputBoxProps {
    sendMessage: Function
    enabled: boolean
}

export interface InputBoxState {
    input: string
}

class InputBox extends React.Component<InputBoxProps, InputBoxState> {
    constructor(props: InputBoxProps) {
        super(props);
        this.state = {
            input: "",
        }
    }
    handleInputChange = (value: string) => {
        this.setState({ input: value });
    }

    handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        this.props.sendMessage(this.state.input)
        this.setState({
            input: "",
        });
    }

    render() {
        return (
            <div className="InputBox">
                <form className="input-form" onSubmit={this.handleSubmit}>
                    <TextField
                        className="input-box"
                        id="filled-multiline-flexible"
                        value={this.state.input}
                        onChange={e => this.handleInputChange(e.currentTarget.value)}
                        margin="dense"
                        variant="filled"
                        disabled={!this.props.enabled}
                    />
                    <Button
                        className="input-submit"
                        color="primary"
                        variant="contained"
                        type="submit"
                        value="Submit"
                        disabled={!this.props.enabled}
                    >
                        Submit
                    </Button>
                </form>
            </div>
        );
    }
}

export default InputBox;