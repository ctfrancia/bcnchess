import React from 'react';
import { useParams } from 'react-router';
import { useSelector } from 'react-redux';
import { selectTournamentById } from '../../../app/tournamentsSlice';

const TournamentDetails = () => {
  const {tournamentId} = useParams()
  const tournament = useSelector(state => selectTournamentById(state, tournamentId))
  //TODO: add tournament image
  return (
    <div className='snippet'>
      <div className='metadata'>
          <time><strong>{tournament.title}</strong></time>
          <time><strong> Share tournament: https://localhost:4000/tournament/{tournament.ID}</strong></time>
      </div>
      <div className='metadata'>
          <strong>Location</strong>: {tournament.location}<br/>
          <strong>Host Contact</strong>: {tournament.tournamentContact}
      </div>
      <pre><code>{tournament.additionalInformation}</code></pre>
      <div className='metadata'>
          <div className='metadata'>
              <time>Date: {new Date(tournament.sDate).toString()}</time>
              <time>Starts: {new Date(tournament.sTime).toString()}</time>
          </div>
          <div className='metadata'>
              <time>Rated: {tournament.isRated}</time>
              <time>Online: {tournament.isOnline}</time>
          </div>
          <div className='metadata'>
              <time>Tournament Type: {tournament.tournamentType}</time>
              <time>Time Control: {tournament.timeControl}</time>
          </div>
          <div className='metadata'>
              <time>Created: {new Date(tournament.created).toString()}</time>
              <time>Expires: {new Date(tournament.expires).toString()}</time>
          </div>
      </div>
    </div>
  );
};

export default TournamentDetails;