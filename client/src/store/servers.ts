import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { EditServer, Server } from "../utils/types";

const initialServersState: Server[] = [];

export const serversSlice = createSlice({
  name: "servers",
  initialState: initialServersState,
  reducers: {
    registerServers: (state, action: PayloadAction<Server[]>) => {
      state.length = 0;
      state.push(...action.payload);
    },
    editServer: (state, action: PayloadAction<EditServer>) => {
      const { id, address } = action.payload;
      const server = state.find((s) => s.id === id);
      if (server) {
        server.address = address;
      }
    },
    deleteServers: (state) => {
      state.length = 0;
    },
  },
});

export const { registerServers, editServer, deleteServers } =
  serversSlice.actions;
