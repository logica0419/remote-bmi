import { useEffect, VFC } from "react";
import { Route, Routes, useNavigate } from "react-router-dom";
import App from "./components/App";
import Login from "./components/Login";
import OAuth from "./components/OAuth";
import { useLoginCheck } from "./utils/login";

const Router: VFC = () => {
  const { authorized, checkCompleted, fetchLoginStatus } = useLoginCheck();
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
          <Route path="/" element={<App />} />
          <Route path="/*" element={<h2>Not Found</h2>} />
        </>
      ) : checkCompleted ? (
        <>
          <Route path="/" element={<Login />} />
          <Route
            path="/*"
            element={() => {
              navigate("/");
            }}
          />
        </>
      ) : (
        <>
          <Route path="/" element={<h2>Checking User Status...</h2>} />
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
