import axios from "axios";
import { useState } from "react";
import { useDispatch } from "react-redux";
import { AppDispatch, setMe } from "./store";
import { GetMeResponse } from "./types";

export const useLoginCheck = () => {
  const [authorized, setAuthorized] = useState(false);
  const [checkCompleted, setCheckCompleted] = useState(false);

  const dispatch = useDispatch<AppDispatch>();

  const fetchLoginStatus = async () => {
    setCheckCompleted(false);

    await axios
      .get<GetMeResponse>("/api/users/me")
      .then(({ data }) => {
        dispatch(setMe({ id: data.id, name: data.name }));
        setAuthorized(true);
      })
      .catch(() => {});

    setCheckCompleted(true);
  };

  return { authorized, checkCompleted, fetchLoginStatus };
};
