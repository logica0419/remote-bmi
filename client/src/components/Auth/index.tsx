import { css } from "@emotion/react";
import axios from "axios";
import { FormEvent, useState, VFC } from "react";
import { useLoginCheck } from "../../utils/login";
import Form from "./Form";

const styles = {
  title: css`
    line-height: min(10px, calc(2vw - 10px));
  `,
};

const Auth: VFC = () => {
  const { fetchLoginStatus } = useLoginCheck();

  const [signUpName, setSignUpName] = useState("");

  const onSignUp = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    axios
      .post("/api/signup", signUpName)
      .then(() => {
        fetchLoginStatus();
      })
      .catch(() => {
        alert("サインアップに失敗しました");
      });
  };

  const [loginName, setLoginName] = useState("");

  const onLogin = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    axios
      .post("/api/login", loginName)
      .then(() => {
        fetchLoginStatus();
      })
      .catch((err) => {
        if (err.status === 404) {
          alert("存在しないユーザーです");
          return;
        }
        alert("ログインに失敗しました");
      });
  };

  return (
    <>
      <h2 css={styles.title}>Sign Up</h2>
      <h3 css={styles.title}>for new users</h3>
      <Form
        action={"Sign Up"}
        onSubmit={onSignUp}
        name={signUpName}
        setName={setSignUpName}
      />
      <br />
      <h2 css={styles.title}>Login</h2>
      <h3 css={styles.title}>for who have an account</h3>
      <Form
        action={"Login"}
        onSubmit={onLogin}
        name={loginName}
        setName={setLoginName}
      />
    </>
  );
};

export default Auth;
