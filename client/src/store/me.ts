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
      state = action.payload;
    },
  },
});

export const { setMe } = meSlice.actions;
