import React from 'react';
import { useSelector } from 'react-redux';
import { Link } from 'react-router-dom';

const LatestTournaments = () => {
  const tournaments = useSelector(state => state.tournaments)

  let tournamentList = []
  Object.keys(tournaments).map(id => tournamentList.push(tournaments[id]))

  return (
    <>
      {tournamentList.length > 0 ? (
        <table>
          <thead>
            <tr>
              <th>Title</th>
              <td>Tournament Date</td>
              <th>Created</th>
              <th>ID</th>
            </tr>
          </thead>
          <tbody>
            {tournamentList.map((tournament, index) => (
              <tr key={index}>
                <td><Link to={`/tournament/${tournament.ID}`}>{tournament.Title}</Link></td>
                <td>{new Date(tournament.TournamentDate).toString()} at {new Date(tournament.MatchTimeStart).toString()}</td>
                <td>{new Date(tournament.Created).toString()}</td>
                <td>{tournament.ID}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p> There's nothing to show yet ... </p>
      )}
    </>
  );
};

export default LatestTournaments;