import React, { useEffect, useRef } from 'react';
import { ListGroup, Form, Button } from 'react-bootstrap';
import { useSelector, useDispatch } from 'react-redux';
import { useHistory } from 'react-router';
import { updateIsBeingEdited, resetCreateTournamentForm } from '../../../app/createTournamentSlice';
import { tournamentAdded, selectTotalTournaments } from '../../../app/tournamentsSlice';

const CTReview = props => {
  const tournamentsCount = useSelector(selectTotalTournaments)

  const title = useSelector(state => state.createTournamentFormData.title);
  const location = useSelector(state => state.createTournamentFormData.location);
  const sDate = useSelector(state => state.createTournamentFormData.sDate);
  const sTime = useSelector(state => state.createTournamentFormData.sTime);
  const timeControl = useSelector(state => state.createTournamentFormData.timeControl);
  const type = useSelector(state => state.createTournamentFormData.type);
  const contactEmail = useSelector(state => state.createTournamentFormData.contactEmail);
  const isRated = useSelector(state => state.createTournamentFormData.isRated);
  const isOnline = useSelector(state => state.createTournamentFormData.isOnline);
  const additionalInformation = useSelector(state => state.createTournamentFormData.additionalInformation);

  const imgRef = useRef(null)

  useEffect(() => {
    if(props.poster) {
      const reader = new FileReader();
      const readerOnLoad = (event) => {
        imgRef.current.src = event.target.result;
      }
  
      reader.addEventListener('load', readerOnLoad);
      reader.readAsDataURL(props.poster);
  
      return () => {
        reader.removeEventListener('load', readerOnLoad);
      }
    }
    
  }, [props.poster])

  const history = useHistory()
  const dispatch = useDispatch()

  const editTournamentData = () => {
    dispatch(updateIsBeingEdited(true));
  }
  const createTournament = () => {
    // create new tournament
    dispatch(tournamentAdded({
      id: tournamentsCount,
      title, location, sDate, sTime, 
      timeControl, type, contactEmail, 
      isRated, isOnline, 
      additionalInformation
    }))
    // clear form
    dispatch(resetCreateTournamentForm())
    // go back
    history.goBack()
  }

  return (
    <div>
      <div>Review Tournament Information:</div>
      <ListGroup>
        <ListGroup.Item><strong>Title: </strong> {title}</ListGroup.Item>
        <ListGroup.Item><strong>Location: </strong> {location}</ListGroup.Item>
        <ListGroup.Item><strong>Starting at: </strong> {sDate} at {sTime}</ListGroup.Item>
        <ListGroup.Item><strong>Time Control: </strong> {timeControl}</ListGroup.Item>
        <ListGroup.Item><strong>Tournament Type: </strong> {type}</ListGroup.Item>
        <ListGroup.Item><strong>Contact Email: </strong> {contactEmail}</ListGroup.Item>
        <ListGroup.Item><strong>Rated: </strong> 
          <Form.Check 
            disabled
            type='checkbox'
            id={`is-rated`}
            value={isRated}
          />
      </ListGroup.Item>
        <ListGroup.Item>
          <strong>Online: </strong>
          <Form.Check 
            disabled
            type='checkbox'
            id={`is-online`}
            value={isOnline}
          />
        </ListGroup.Item>
        <ListGroup.Item><strong>Poster: </strong> <img ref={imgRef} src='' alt='uploaded poster preview' style={{ maxWidth: '7rem', maHeight: '7rem' }} /></ListGroup.Item>
        <ListGroup.Item><strong>Additional Information: </strong> {additionalInformation}</ListGroup.Item>
      </ListGroup>
      <Button variant='secondary' onClick={editTournamentData}>Edit</Button>
      <Button variant='primary' onClick={createTournament}>Create Tournament</Button>
    </div>
  );
};

export default CTReview;