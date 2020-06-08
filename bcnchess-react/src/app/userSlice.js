import { createSlice } from '@reduxjs/toolkit';

const userSlice = createSlice({
  name: 'user',
  initialState: {
    isLoggedIn: false
  },
  reducers: {
    updateIsLoggedIn: (state, action) => {
      state.isLoggedIn = action.payload
    }
  }
});

export const { updateIsLoggedIn } = userSlice.actions;

export default userSlice.reducer;
