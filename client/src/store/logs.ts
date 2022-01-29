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
      state.unshift(action.payload);
    },
    deleteLogs: (state) => {
      state.length = 0;
    },
  },
});

export const { registerLogs, addLog, deleteLogs } = logsSlice.actions;
