import { configureStore } from "@reduxjs/toolkit";
import { logsSlice } from "./logs";
import { meSlice } from "./me";
import { serversSlice } from "./servers";

export const store = configureStore({
  reducer: {
    me: meSlice.reducer,
    servers: serversSlice.reducer,
    logs: logsSlice.reducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
