import { StrictMode } from "react";
import { render } from "react-dom";
import App from "./components/App";
import GlobalStyle from "./style";

render(
  <StrictMode>
    <GlobalStyle />
    <App />
  </StrictMode>,
  document.getElementById("root")
);
