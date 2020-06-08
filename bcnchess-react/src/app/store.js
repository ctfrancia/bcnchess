import { configureStore } from '@reduxjs/toolkit';
import { logger } from 'redux-logger';
import userReducer from './userSlice';
import tournamentsReducer from './tournamentsSlice';

export default configureStore({
  reducer: {
    user: userReducer,
    tournaments: tournamentsReducer
  },
  middleware: [logger]
});
