import { css } from "@emotion/react";
import VersionDisplay from "./Version";

const style = {
  app: css`
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: center;
    color: #000000;
    min-height: 100vh;
    font-size: calc(5px + 2vmin);
  `,
  title: css`
    font-size: calc(10px + 2vmin);
    line-height: calc(15px + 2vmin);
  `,
  line: css`
    margin: 1em 0;
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
      <VersionDisplay />
    </div>
  );
}

export default App;
