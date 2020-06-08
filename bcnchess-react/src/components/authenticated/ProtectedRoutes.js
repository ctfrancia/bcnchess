import React from 'react';
import { Switch, Route, Redirect } from 'react-router-dom';
import Home from './home/Home'
import About from './about/About'
import TournamentDetails from './home/TournamentDetails';

const ProtectedRoutes = () => {
  return (
    <Switch>
      <Route exact path='/'><Redirect to='/home' /></Route>
      <Route path='/home'><Home /></Route>
      <Route path='/about'><About /></Route>
      <Route path='/tournament/:tournamentId'><TournamentDetails /></Route>
    </Switch>
  );
};

export default ProtectedRoutes;