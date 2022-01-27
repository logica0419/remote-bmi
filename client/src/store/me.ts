import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Me } from "../utils/types";

const initialMeState: Me = {
  id: "",
  name: "",
};

export const meSlice = createSlice({
  name: "me",
  initialState: initialMeState,
  reducers: {
    setMe: (state, action: PayloadAction<Me>) => {
      state.id = action.payload.id;
      state.name = action.payload.name;
    },
  },
});

export const { setMe } = meSlice.actions;
