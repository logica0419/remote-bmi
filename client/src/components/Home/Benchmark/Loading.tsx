import { css, keyframes } from "@emotion/react";
import { VFC } from "react";

const animations = {
  spin: keyframes`
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }`,
};

const styles = {
  background: css`
    z-index: 1;
    background-color: #000000;
    opacity: 0.5;
    width: 100vw;
    height: 100vh;
    position: fixed;
  `,
  modal: css`
    z-index: 2;
    background-color: #ffffff;
    width: 400px;
    height: 200px;
    border-radius: 20px;
    opacity: 1;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    margin: auto;
    font-size: 20px;
    align-items: center;
  `,
  text: css`
    margin-top: 50px;
  `,
  circle: css`
    display: inline-block;
    width: 30px;
    height: 30px;
    border: 5px solid #000000;
    border-radius: 100%;
    animation: ${animations.spin} 1s linear infinite;
    ::before {
      content: "";
      position: relative;
      display: block;
      width: 15px;
      height: 10px;
      bottom: 2px;
      right: 2px;
      transform: rotate(45deg);
      background-color: #ffffff;
    }
  `,
};

const Loading: VFC = () => {
  return (
    <>
      <div css={styles.background} />
      <div css={styles.modal}>
        <h2 css={styles.text}>ベンチマーク中...</h2>
        <div css={styles.circle} />
      </div>
    </>
  );
};

export default Loading;
