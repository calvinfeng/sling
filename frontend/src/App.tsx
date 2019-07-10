import React from 'react';
import './App.css';
import Login from './components/Login'
import ReduxTester from './components/ReduxTester'

import { Provider } from 'react-redux';
import { applyMiddleware, createStore } from 'redux';
import { rootReducer } from './store';
import logger from 'redux-logger';

const store = createStore(rootReducer, applyMiddleware(logger))

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
