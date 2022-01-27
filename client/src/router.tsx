import { useEffect, VFC } from "react";
import { Route, Routes, useNavigate } from "react-router-dom";
import Home from "./components/Home";
import Login from "./components/Login";
import OAuth from "./components/OAuth";
import { useLoginCheck } from "./utils/login";

const Router: VFC = () => {
  const { authorized, isFetching, fetchLoginStatus } = useLoginCheck();
  const navigate = useNavigate();

  useEffect(() => {
    fetchLoginStatus();
  }, []);

  return (
    <Routes>
      <Route
        path="/oauth"
        element={<OAuth fetchLoginStatus={fetchLoginStatus} />}
      />
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
          <Route path="/" element={<Login />} />
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
