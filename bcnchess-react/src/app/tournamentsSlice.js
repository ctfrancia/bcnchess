import { createSlice, createEntityAdapter } from '@reduxjs/toolkit';

const tournamentAdapter = createEntityAdapter()

const tournamentsSlice = createSlice({
  name: 'tournaments',
  initialState: tournamentAdapter.getInitialState(),
  reducers: {
    tournamentAdded: tournamentAdapter.addOne,
    tournamentUpdated: tournamentAdapter.updateOne,
    tournamentRemoved: tournamentAdapter.removeOne,
    allTournamentsUpdated: tournamentAdapter.addMany
  }
});

export const { 
  tournamentAdded,
  tournamentUpdated,
  tournamentRemoved,
  allTournamentsUpdated
} = tournamentsSlice.actions;

export const {
  selectById: selectTournamentById,
  selectIds: selectTournamentIds,
  selectAll: selectAllTournaments,
  selectTotal: selectTotalTournaments
} = tournamentAdapter.getSelectors(state => state.tournaments)

export default tournamentsSlice.reducer;
