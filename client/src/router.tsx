import axios from "axios";
import { useEffect, useState, VFC } from "react";
import { useDispatch } from "react-redux";
import { Route, Routes } from "react-router-dom";
import App from "./components/App";
import Login from "./components/Login";
import OAuth from "./components/OAuth";
import { AppDispatch, setMe } from "./store";

interface GetMeResponse {
  id: string;
  name: string;
}

const Router: VFC = () => {
  const [authorized, setAuthorized] = useState(false);

  const CheckLogin: VFC = () => {
    const [checkComplete, setCheckComplete] = useState(false);

    const dispatch = useDispatch<AppDispatch>();

    useEffect(() => {
      axios
        .get<GetMeResponse>("/api/user/me")
        .then(({ data }) => {
          dispatch(setMe({ id: data.id, name: data.name }));

          setAuthorized(true);
        })
        .catch(() => {
          setCheckComplete(true);
        });
    }, []);

    return checkComplete ? <Login /> : <h2>Checking User Status...</h2>;
  };

  return (
    <Routes>
      <Route path="/" element={authorized ? <App /> : <CheckLogin />} />
      <Route path="/oauth" element={<OAuth setAuthorized={setAuthorized} />} />
      <Route path="*" element={<h2>Not Found</h2>} />
    </Routes>
  );
};

export default Router;
