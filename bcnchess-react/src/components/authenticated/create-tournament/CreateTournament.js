import React from 'react';
import { Form, Button } from 'react-bootstrap';

const CreateTournament = () => {
  return (
    <Form>
      <Form.Group controlId="createTournamentFormTitle">
        <Form.Label>Title</Form.Label>
        <Form.Control type="text" placeholder="title" />
        <Form.Text className="text-muted">
          Name of Tournament
        </Form.Text>
      </Form.Group>
      <Form.Group controlId="createTournamentFormTitle">
        <Form.Label>Location</Form.Label>
        <Form.Control type="text" placeholder="location" />
        <Form.Text className="text-muted">
          Address of Tournament/Website where tournament is held
        </Form.Text>
      </Form.Group>
      <Form.Group controlId="createTournamentFormTitle">
        <Form.Label>Date</Form.Label>
        <Form.Control type="date" placeholder="start date" />
        <Form.Control type="time" placeholder="start time" />
      </Form.Group>
      <Form.Group controlId="createTournamentFormTitle">
        <Form.Label>Time Control (3+2/5+0/etc):</Form.Label>
        <Form.Control type="text" placeholder="Time Control" />
      </Form.Group>
      <Form.Group controlId="createTournamentFormTitle">
        <Form.Label>Type(Round Robin/Swiss/etc):</Form.Label>
        <Form.Control type="text" placeholder="Type" />
      </Form.Group>
      <Form.Group controlId="createTournamentFormTitle">
        <Form.Label>Tournament Contact Email:</Form.Label>
        <Form.Control type="email" placeholder="email" />
        <Form.Text className="text-muted">
          Any questions should be emailed to this email address
        </Form.Text>
      </Form.Group>
      <Form.Group controlId="createTournamentFormTitle">
        <Form.Check 
          type='checkbox'
          id={`is-rated`}
          label={`Is rated`}
        />
        <Form.Check 
          type='checkbox'
          id={`is-online`}
          label={`Is online`}
        />
      </Form.Group>
      <Form.Group>
        <Form.File id="exampleFormControlFile1" label="Tournament Poster" />
      </Form.Group>
      <Form.Group controlId="createTournamentFormTitle">
        <Form.Label>Additional information:</Form.Label>
        <Form.Control as="textarea" rows="3" />
      </Form.Group>
      <Button variant="primary" type="submit">
        Submit
      </Button>
    </Form>
  );
};

export default CreateTournament;