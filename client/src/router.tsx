import { useEffect, VFC } from "react";
import { Route, Routes, useNavigate } from "react-router-dom";
import Home from "./components/Home";
import Auth from "./components/Auth";
import { useLoginCheck } from "./utils/login";

const Router: VFC = () => {
  const { authorized, isFetching, fetchLoginStatus } = useLoginCheck();
  const navigate = useNavigate();

  useEffect(() => {
    fetchLoginStatus();
  }, []);

  return (
    <Routes>
      {authorized ? (
        <>
          <Route path="/" element={<Home />} />
          <Route path="/*" element={<h2>Not Found</h2>} />
        </>
      ) : isFetching ? (
        <>
          <Route path="/" element={<h2>Checking User Status...</h2>} />
          <Route
            path="/*"
            element={() => {
              navigate("/");
            }}
          />
        </>
      ) : (
        <>
          <Route path="/" element={<Auth />} />
          <Route
            path="/*"
            element={() => {
              navigate("/");
            }}
          />
        </>
      )}
    </Routes>
  );
};

export default Router;
