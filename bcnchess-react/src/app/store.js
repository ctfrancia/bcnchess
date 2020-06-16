import { configureStore, getDefaultMiddleware } from '@reduxjs/toolkit';
import { logger } from 'redux-logger';
import userReducer from './userSlice';
import tournamentsReducer from './tournamentsSlice';
import createTournamentFormDataReducer from './createTournamentSlice';

export default configureStore({
  reducer: {
    user: userReducer,
    tournaments: tournamentsReducer,
    createTournamentFormData: createTournamentFormDataReducer
  },
  middleware: [...getDefaultMiddleware(), logger]
});
