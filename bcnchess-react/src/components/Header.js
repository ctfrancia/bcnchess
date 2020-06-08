import React from 'react';
import { Navbar, Nav, Button } from 'react-bootstrap'
import { Link } from 'react-router-dom'
import { updateIsLoggedIn } from '../app/userSlice';
import { useDispatch } from 'react-redux';

const Header = () => {
  const dispatch = useDispatch();

  return (
    <Navbar bg="light" expand="lg">
      <Navbar.Brand as={Link} to="/home">React-Bootstrap</Navbar.Brand>
      <Navbar.Toggle aria-controls="basic-navbar-nav" />
      <Navbar.Collapse id="basic-navbar-nav">
        <Nav className="mr-auto">
          <Nav.Link as={Link} to="/home">Home</Nav.Link>
          <Nav.Link as={Link} to="/about">About</Nav.Link>
        </Nav>
        <Nav>
          <Button variant='outline-danger' onClick={() => dispatch(updateIsLoggedIn(false))}>logout</Button>
        </Nav>
      </Navbar.Collapse>
    </Navbar>
  );
};

export default Header;