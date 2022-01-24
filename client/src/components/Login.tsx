import axios from "axios";
import { useEffect, VFC } from "react";

const Login: VFC = () => {
  useEffect(() => {
    axios
      .get<string>("/api/oauth/callback")
      .then(({ data }) => {
        window.location.replace(data);
      })
      .catch(() => {
        alert("failed to get OAuth callback URL");
      });
  }, []);

  return <h2>Logging in...</h2>;
};

export default Login;
