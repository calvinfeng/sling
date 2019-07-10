import React from 'react';
import logo from './logo.svg';
import './App.css';
import Login from './components/Login'

import { Provider } from 'react-redux';
import { createStore } from 'redux';
import slingApp from './reducers';

const store = createStore(slingApp)

const App: React.FC = () => {
  return (
    <Provider store={store}>
      <div className="App">
        <Login />
      </div>
    </Provider>
  );
}

export default App;
