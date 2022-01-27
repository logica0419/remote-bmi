import { configureStore } from "@reduxjs/toolkit";
import { meSlice } from "./me";

export const store = configureStore({
  reducer: {
    me: meSlice.reducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
