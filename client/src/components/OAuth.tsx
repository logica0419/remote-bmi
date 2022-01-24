import { css } from "@emotion/react";
import axios from "axios";
import { Dispatch, VFC } from "react";
import { SetStateAction, useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";

interface Props {
  setAuthorized: Dispatch<SetStateAction<boolean>>;
}

const OAuth: VFC<Props> = ({ setAuthorized }) => {
  const [searchParams] = useSearchParams();
  const code = searchParams.get("code");
  const navigate = useNavigate();

  useEffect(() => {
    axios
      .post("/api/oauth/code", code)
      .then(({}) => {
        setAuthorized(true);
        navigate("/");
      })
      .catch(() => {
        alert("failed to get Access Token");
      });
  }, []);

  return <h2>Getting Access Token...</h2>;
};

export default OAuth;
