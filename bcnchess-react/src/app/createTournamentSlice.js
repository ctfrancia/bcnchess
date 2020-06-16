import { createSlice } from '@reduxjs/toolkit';

const createTournamentSlice = createSlice({
  name: 'createTournamentFormData',
  initialState: {
    title: '',
    location: '',
    sDate: '',
    sTime: '',
    timeControl: '',
    type: '',
    contactEmail: '',
    isRated: false,
    isOnline: false,
    additionalInformation: '',
    isBeingEdited: true
  },
  reducers: {
    updateTitle: (state, action) => {
      state.title = action.payload
    },
    updateLocation: (state, action) => {
      state.location = action.payload
    },
    updateSDate: (state, action) => {
      state.sDate = action.payload
    },
    updateSTime: (state, action) => {
      state.sTime = action.payload
    },
    updateTimeControl: (state, action) => {
      state.timeControl = action.payload
    },
    updateType: (state, action) => {
      state.type = action.payload
    },
    updateContactEmail: (state, action) => {
      state.contactEmail = action.payload
    },
    updateIsRated: (state, action) => {
      state.isRated = action.payload
    },
    updateIsOnline: (state, action) => {
      state.isOnline = action.payload
    },
    updateAdditionalInformation: (state, action) => {
      state.additionalInformation = action.payload
    },
    updateIsTournamentDataEntered: (state, action) => {
      state.isTournamentDataEntered = action.payload
    },
    updateIsBeingEdited: (state, action) => {
      state.isBeingEdited = action.payload
    },
    resetCreateTournamentForm: (state) => {
      state.title = ''
      state.location = ''
      state.sDate = ''
      state.sTime = ''
      state.timeControl = ''
      state.type = ''
      state.contactEmail = ''
      state.isRated = false
      state.isOnline = false
      state.additionalInformation = ''
      state.isBeingEdited = true
    }
  }
});

export const { 
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
  updateIsBeingEdited,
  resetCreateTournamentForm
} = createTournamentSlice.actions;

export default createTournamentSlice.reducer;
