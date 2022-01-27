import { css } from "@emotion/react";
import { StrictMode } from "react";
import { render } from "react-dom";
import { Provider } from "react-redux";
import { BrowserRouter } from "react-router-dom";
import Router from "./router";
import { store } from "./store";
import GlobalStyle from "./style";

const style = {
  app: css`
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
    color: #000000;
    min-height: 95vh;
    font-size: min(25px, calc(5px + 2vw));
  `,
  title: css`
    font-size: min(30px, calc(10px + 2vw));
    line-height: min(35px, calc(15px + 2vw));
  `,
  line: css`
    margin: 1em 0;
    width: max(400px, 90vw);
    height: 1px;
    background: #000000;
  `,
};

render(
  <StrictMode>
    <GlobalStyle />
    <div css={style.app}>
      <div css={style.title}>
        <h1>Remote-BMI</h1>- Bench Marker Web UI for ISUCON Practice -
      </div>
      <div css={style.line} />
      <Provider store={store}>
        <BrowserRouter>
          <Router />
        </BrowserRouter>
      </Provider>
    </div>
  </StrictMode>,
  document.getElementById("root")
);
