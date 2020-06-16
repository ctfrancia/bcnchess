import React from 'react';
import { Form, Button } from 'react-bootstrap';
import { useSelector, useDispatch } from 'react-redux';
import { 
  updateTitle, 
  updateLocation, 
  updateSDate, 
  updateSTime, 
  updateTimeControl, 
  updateType, 
  updateContactEmail, 
  updateIsRated, 
  updateIsOnline, 
  updateAdditionalInformation,
  updateIsBeingEdited
} from '../../../app/createTournamentSlice';

const CTForm = props => {
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

  const dispatch = useDispatch();

  const setTitle = value => dispatch(updateTitle(value));
  const setLocation = value => dispatch(updateLocation(value));
  const setSDate = value => dispatch(updateSDate(value));
  const setSTime = value => dispatch(updateSTime(value));
  const setTimeControl = value => dispatch(updateTimeControl(value));
  const setType = value => dispatch(updateType(value));
  const setContactEmail = value => dispatch(updateContactEmail(value));
  const setIsRated = value => dispatch(updateIsRated(value));
  const setIsOnline = value => dispatch(updateIsOnline(value));
  const setAdditionalInformation = value => dispatch(updateAdditionalInformation(value));

  const reviewNewTournamentData = () => {
    // navigate to review screen
    dispatch(updateIsBeingEdited(false));
  }

  return (
    <Form>
      <Form.Group controlId="createTournamentFormTitle">
        <Form.Label>Title</Form.Label>
        <Form.Control type="text" placeholder="title" value={title} onChange={ev => setTitle(ev.target.value)} />
        <Form.Text className="text-muted">
          Name of Tournament
        </Form.Text>
      </Form.Group>
      <Form.Group controlId="createTournamentFormLocation">
        <Form.Label>Location</Form.Label>
        <Form.Control type="text" placeholder="location" value={location} onChange={ev => setLocation(ev.target.value)} />
        <Form.Text className="text-muted">
          Address of Tournament/Website where tournament is held
        </Form.Text>
      </Form.Group>
      <Form.Group controlId="createTournamentStartDateAndTime">
        <Form.Label>Date</Form.Label>
        <Form.Control type="date" placeholder="start date" value={sDate} onChange={ev => setSDate(ev.target.value)} />
        <Form.Control type="time" placeholder="start time" value={sTime} onChange={ev => setSTime(ev.target.value)} />
      </Form.Group>
      <Form.Group controlId="createTournamentFormTimeControl">
        <Form.Label>Time Control (3+2/5+0/etc):</Form.Label>
        <Form.Control type="text" placeholder="Time Control" value={timeControl} onChange={ev => setTimeControl(ev.target.value)} />
      </Form.Group>
      <Form.Group controlId="createTournamentFormTitle">
        <Form.Label>Type(Round Robin/Swiss/etc):</Form.Label>
        <Form.Control type="text" placeholder="Type" value={type} onChange={ev => setType(ev.target.value)} />
      </Form.Group>
      <Form.Group controlId="createTournamentFormTitle">
        <Form.Label>Tournament Contact Email:</Form.Label>
        <Form.Control type="email" placeholder="email" value={contactEmail} onChange={ev => setContactEmail(ev.target.value)}  />
        <Form.Text className="text-muted">
          Any questions should be emailed to this email address
        </Form.Text>
      </Form.Group>
      <Form.Group controlId="createTournamentFormTitle">
        <Form.Check 
          type='checkbox'
          id={`is-rated`}
          label={`Is rated`}
          value={isRated}
          onChange={() => setIsRated(!isRated)}
        />
        <Form.Check 
          type='checkbox'
          id={`is-online`}
          label={`Is online`}
          value={isOnline}
          onChange={() => setIsOnline(!isOnline)}
        />
      </Form.Group>
      <Form.Group>
        <Form.File id="exampleFormControlFile1" label="Tournament Poster" value={props.poster} onChange={ev => props.setPoster(ev.target.files[0])} accept=".jpg, .jpeg, .png" />
      </Form.Group>
      <Form.Group controlId="createTournamentFormTitle">
        <Form.Label>Additional information:</Form.Label>
        <Form.Control as="textarea" rows="3" value={additionalInformation} onChange={ev => setAdditionalInformation(ev.target.value)} />
      </Form.Group>
      <Button variant="primary" onClick={reviewNewTournamentData} >
        Review
      </Button>
    </Form>
  );
};

export default CTForm;