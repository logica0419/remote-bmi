import { configureStore, createSlice, PayloadAction } from "@reduxjs/toolkit";

interface MeState {
  id: string;
  name: string;
}

const initialMeState: MeState = {
  id: "",
  name: "",
};

const meSlice = createSlice({
  name: "me",
  initialState: initialMeState,
  reducers: {
    setMe: (state, action: PayloadAction<MeState>) => {
      state.id = action.payload.id;
      state.name = action.payload.name;
    },
  },
});

export const { setMe } = meSlice.actions;

export const store = configureStore({
  reducer: {
    me: meSlice.reducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
