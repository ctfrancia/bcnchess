import React, { useState } from 'react';

const LatestTournaments = () => {
  const [tournaments, set] = useState([])
  return (
    <>
      {tournaments.length > 0 ? (
        <table>
          <tr>
              <th>Title</th>
              <td>Tournament Date</td>
              <th>Created</th>
              <th>ID</th>
          </tr>
          {tournaments.map((tournament, index) => (
            <tr key={index}>
              <td><a href='/tournament/{{.ID}}'>{tournament.Title}</a></td>
              <td>{tournament.TournamentDate} at {tournament.MatchTimeStart}</td>
              <td>{tournament.Created}</td>
              <td>{tournament.ID}</td>
            </tr>
          ))}
        </table>
      ) : (
        <p> There's nothing to show yet ... </p>
      )}
    </>
  );
};

export default LatestTournaments;