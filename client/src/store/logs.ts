import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Log } from "../utils/types";

const initialLogsState: Log[] = [];

export const logsSlice = createSlice({
  name: "logs",
  initialState: initialLogsState,
  reducers: {
    registerLogs: (state, action: PayloadAction<Log[]>) => {
      state.length = 0;
      state.push(...action.payload);
    },
    addLog: (state, action: PayloadAction<Log>) => {
      state.push(action.payload);
    },
  },
});

export const { registerLogs, addLog } = logsSlice.actions;
