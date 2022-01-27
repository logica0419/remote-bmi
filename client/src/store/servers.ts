import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { EditServer, Server } from "../utils/types";

const initialServersState: Server[] = [];

export const serversSlice = createSlice({
  name: "server",
  initialState: initialServersState,
  reducers: {
    registerServers: (state, action: PayloadAction<Server[]>) => {
      state = action.payload;
    },
    editServer: (state, action: PayloadAction<EditServer>) => {
      const { id, address } = action.payload;
      const server = state.find((s) => s.id === id);
      if (server) {
        server.address = address;
      }
    },
    deleteServers: (state) => {
      state = [];
    },
  },
});

export const { registerServers, editServer, deleteServers } =
  serversSlice.actions;
