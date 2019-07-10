import * as React from 'react';
import './component.css';

export interface InputBoxProps {
    
}
 
export interface InputBoxState {
    input: string
}
 
class InputBox extends  React.Component<InputBoxProps, InputBoxState> {
    constructor(props: InputBoxProps) {
        super(props);
        this.state = {
            input:"",
        }
    }
    handleInputChange = (e:React.FormEvent<HTMLInputElement>) => {
        this.setState({input: e.currentTarget.value});
        //console.log(this.state.input);
    }

    handleSubmit = (e:React.FormEvent) => {
        e.preventDefault();
        //######
        this.setState({
            input: "",
        });
        //console.log(this.state.input);
    }

    render() { 
        return (  
            <div className="InputBox">
                <form className="input-form" onSubmit={this.handleSubmit}>
                    <input type="text" onChange={this.handleInputChange} value={this.state.input}/>
                    <input type="submit" value="Submit" />
                </form>
            </div>
        );
    }
}
 
export default InputBox;