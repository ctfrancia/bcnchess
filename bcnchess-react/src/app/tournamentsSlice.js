import { createSlice } from '@reduxjs/toolkit';

const tournamentsSlice = createSlice({
  name: 'user',
  initialState: {},
  reducers: {
    addTournaments: (state, action) => {
      action.payload.map(tournament => {
        state[tournament.ID] = tournament
      })
    },
    removeTournament: (state, action) => {
      delete state[action.payload]
    },
    updateTournament: (state, action) => {
      state[action.payload.ID] = { ...state[action.payload.ID], ...action.payload }
    }
  }
});

export const { addTournaments, removeTournament, updateTournament } = tournamentsSlice.actions;

export default tournamentsSlice.reducer;
