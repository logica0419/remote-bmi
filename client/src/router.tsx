import { useEffect, VFC } from "react";
import { Route, Routes, useNavigate } from "react-router-dom";
import Auth from "./components/Auth";
import Home from "./components/Home";
import { useLoginCheck } from "./utils/login";

const RootNavigator: VFC = () => {
  const navigate = useNavigate();

  useEffect(() => {
    navigate("/");
  }, []);

  return <></>;
};

const Router: VFC = () => {
  const { authorized, isFetching, fetchLoginStatus } = useLoginCheck();

  useEffect(() => {
    fetchLoginStatus();
  }, []);

  return (
    <Routes>
      {authorized ? (
        <>
          <Route path="/" element={<Home />} />
          <Route path="/*" element={<RootNavigator />} />
        </>
      ) : isFetching ? (
        <>
          <Route path="/" element={<h2>Checking User Status...</h2>} />
          <Route path="/*" element={<RootNavigator />} />
        </>
      ) : (
        <>
          <Route
            path="/"
            element={<Auth fetchLoginStatus={fetchLoginStatus} />}
          />
          <Route path="/*" element={<RootNavigator />} />
        </>
      )}
    </Routes>
  );
};

export default Router;
