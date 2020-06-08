import React from 'react';
import { useParams } from 'react-router';
import { useSelector } from 'react-redux';

const TournamentDetails = () => {
  const {tournamentId} = useParams()
  const tournament = useSelector(state => state.tournaments[tournamentId])
  //TODO: add tournament image
  return (
    <div class='snippet'>
      <div class='metadata'>
          <time><strong>{tournament.Title}</strong></time>
          <time><strong> Share tournament: https://localhost:4000/tournament/{tournament.ID}</strong></time>
      </div>
      <div class='metadata'>
          <strong>Location</strong>: {tournament.Location}<br/>
          <strong>Host Contact</strong>: {tournament.TournamentContact}
      </div>
      <pre><code>{tournament.AdditionalInformation}</code></pre>
      <div class='metadata'>
          <div class='metadata'>
              <time>Date: {new Date(tournament.TournamentDate).toString()}</time>
              <time>Starts: {new Date(tournament.MatchTimeStart).toString()}</time>
          </div>
          <div class='metadata'>
              <time>Rated: {tournament.Rated}</time>
              <time>Online: {tournament.IsOnline}</time>
          </div>
          <div class='metadata'>
              <time>Tournament Type: {tournament.TournamentType}</time>
              <time>Time Control: {tournament.TimeControl}</time>
          </div>
          <div class='metadata'>
              <time>Created: {new Date(tournament.Created).toString()}</time>
              <time>Expires: {new Date(tournament.Expires).toString()}</time>
          </div>
      </div>
    </div>
  );
};

export default TournamentDetails;