import React, { Component } from 'react';
import MainContentComponent from './components/mainContent'
import LoginFormComponent from './components/loginForm'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import logo from './logo.svg';
import './App.css';

class App extends Component {
  render() {
    return ((
      <Router>
        <div className='App'>
        <Route path='/'component={MainContentComponent}></Route>
        <Route path='/login' component={LoginFormComponent}></Route>
        </div>
      </Router>
    ));
  }
}

export default App;
