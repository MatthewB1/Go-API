import React, { Component } from 'react';
import MainContentComponent from './components/mainContent'
import LoginFormComponent from './components/loginForm'
import DashboardComponent from './components/dashboard'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import { Redirect } from 'react-router-dom'
import decode from 'jwt-decode';
import './App.css';


const checkAuth = () => {
  const token = localStorage.getItem('token');
  //if no token, return false
  if (!token) {
    console.log("no token set")
    return false;
  }

  try {
    const { exp } = decode(token, {header:true});

    //check exp against current time, if expired return false
    if (exp < new Date().getTime() / 1000) {
      console.log("token invalid")
      return false;
    }

  } catch (e) {
    console.log("caught error : " + e)
    return false;
  }

  return true;
}

const AuthRoute = ({ component: Component, ...rest }) => (
  <Route {...rest} render={props => (
    //check token auth, if valid continue to content
    checkAuth() ? (
      <Component {...props} />
    ) : (
      //else, redirect to login
        <Redirect to={{ pathname: '/login' }} />
      )
  )} />
)

class App extends Component {
  render() {
    return ((
      <Router>
        <div className='App'>
        <Route path='/'component={MainContentComponent}></Route>
        <Route exact path='/login' component={LoginFormComponent}></Route>
        <AuthRoute exact path='/dashboard' component={DashboardComponent}></AuthRoute>
        </div>
      </Router>
    ));
  }
}

export default App;
