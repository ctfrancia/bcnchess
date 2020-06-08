import React from 'react';
import './App.css';
import { BrowserRouter as Router, Route, Redirect, Switch } from 'react-router-dom'
import Header from './components/Header';
import Home from './components/home/Home';
import About from './components/about/About';

function App() {
  return (
    <Router>
      <Header />
      <Switch>
        <Route exact path='/'><Redirect to='/home' /></Route>
        <Route path='/home'><Home /></Route>
        <Route path='/about'><About /></Route>
      </Switch>
    </Router>
  );
}

export default App;
