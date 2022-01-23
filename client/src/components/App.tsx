import { css } from "@emotion/react";

const style = {
  app: css`
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: center;
    color: #000000;
    min-height: 100vh;
  `,
  title: css`
    font-size: calc(10px + 2vmin);
    line-height: calc(15px + 2vmin);
  `,
  line: css`
    margin-top: 1em;
    width: 70vw;
    height: 1px;
    background: #000000;
  `,
};

function App() {
  return (
    <div css={style.app}>
      <div css={style.title}>
        <h1>Remote-BMI</h1>- Bench Marker Web UI for ISUCON Practice -
      </div>
      <div css={style.line} />
    </div>
  );
}

export default App;
