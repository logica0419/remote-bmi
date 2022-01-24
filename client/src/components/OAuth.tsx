import axios from "axios";
import { VFC } from "react";
import { useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";

interface Props {
  fetchLoginStatus: () => Promise<void>;
}

const OAuth: VFC<Props> = ({ fetchLoginStatus }) => {
  const [searchParams] = useSearchParams();
  const code = searchParams.get("code");
  const navigate = useNavigate();

  useEffect(() => {
    if (!code) {
      navigate("/");
      return;
    }

    axios
      .post("/api/oauth/code", code)
      .then(async () => {
        await fetchLoginStatus();
        navigate("/");
      })
      .catch(() => {
        alert("failed to get Access Token");
        navigate("/");
      });
  }, []);

  return <h2>Getting Access Token...</h2>;
};

export default OAuth;
