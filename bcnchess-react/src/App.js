import React from 'react';
import { useSelector } from 'react-redux';

import './App.css';
import { BrowserRouter as Router, Route, Switch, Redirect } from 'react-router-dom'
import Header from './components/Header';
import Login from './components/Login';
import ProtectedRoutes from './components/authenticated/ProtectedRoutes';
import { Container } from 'react-bootstrap';


function App() {
  const isUserLoggedIn = useSelector(state => state.user.isLoggedIn)

  return (
    <Router>
      <Header />
      <Container>
        <Switch>
          <Route exact path='/login'>{isUserLoggedIn ? <ProtectedRoutes /> : <Login />}</Route>
          <Route path='/'>{isUserLoggedIn ? <ProtectedRoutes /> : <Redirect to='/login' />}</Route>
        </Switch>
      </Container>
      
    </Router>
  );
}

export default App;
