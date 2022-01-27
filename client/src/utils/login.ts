import axios from "axios";
import { useState } from "react";
import { useDispatch } from "react-redux";
import { AppDispatch } from "../store";
import { setMe } from "../store/me";
import { GetMeResponse } from "./types";

export const useLoginCheck = () => {
  const [authorized, setAuthorized] = useState(false);
  const [isFetching, setIsFetching] = useState(true);

  const dispatch = useDispatch<AppDispatch>();

  const fetchLoginStatus = async () => {
    setIsFetching(true);

    await axios
      .get<GetMeResponse>("/api/users/me")
      .then(({ data }) => {
        dispatch(setMe({ id: data.id, name: data.name }));
        setAuthorized(true);
      })
      .catch(() => {
        setAuthorized(false);
      });

    setIsFetching(false);
  };

  return { authorized, isFetching, fetchLoginStatus };
};
