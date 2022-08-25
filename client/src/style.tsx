import { css, Global } from "@emotion/react";
import { FC } from "react";

const GlobalStyle: FC = () => {
  return (
    <Global
      styles={css`
        body {
          background-color: #ffffff;
          margin: 0;
          min-width: 420px;
          min-height: 300px;
          font-family: -apple-system, BlinkMacSystemFont, "Roboto", "Oxygen",
            "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue";
          -webkit-font-smoothing: antialiased;
          -moz-osx-font-smoothing: grayscale;
        }
      `}
    />
  );
};

export default GlobalStyle;
