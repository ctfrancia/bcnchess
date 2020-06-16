import React from 'react';
import { useSelector } from 'react-redux';
import { Link } from 'react-router-dom';
import { selectAllTournaments } from '../../../app/tournamentsSlice';

const LatestTournaments = () => {
  const tournaments = useSelector(selectAllTournaments)

  return (
    <>
      {tournaments.length > 0 ? (
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
            {tournaments.map((tournament, index) => (
              <tr key={index}>
                <td><Link to={`/tournament/${tournament.id}`}>{tournament.title}</Link></td>
                <td>{new Date(tournament.sDate).toString()} at {new Date(tournament.sTime).toString()}</td>
                <td>{new Date(tournament.created).toString()}</td>
                <td>{tournament.id}</td>
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