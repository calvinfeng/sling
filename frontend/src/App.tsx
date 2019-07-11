import React from 'react';
import './App.css';
import Login from './components/Login'
import ReduxTester from './components/ReduxTester'
import MessagePage from './MessagePage'

import { Provider } from 'react-redux';
import { applyMiddleware, createStore } from 'redux';
import { rootReducer } from './store';
import logger from 'redux-logger';

const store = createStore(rootReducer, applyMiddleware(logger))

const initialState = {
  loggedIn: false
}

interface AppState {
  loggedIn: boolean
}

class App extends React.Component<{}, AppState> {
  readonly state: AppState = initialState

  componentDidMount() {
    if (localStorage.getItem('jwt_token')) {
      this.setState({ loggedIn: true })
    }
  }

  setLoggedIn(isLoggedIn: boolean) {
    if (!isLoggedIn) {
      localStorage.removeItem('jwt_token')
    }
    this.setState({ loggedIn: isLoggedIn })
  }

  render() {
    return (
      <Provider store={store}>
        <div className="App">
          {this.state.loggedIn ?
            <MessagePage setLoggedOut={() => this.setLoggedIn(false)} /> :
            <Login setLoggedIn={() => this.setLoggedIn(true)} />
          }
        </div>
      </Provider>
    );
  }
}

export default App;
