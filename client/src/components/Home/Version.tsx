import { css } from "@emotion/react";
import axios from "axios";
import { FC, useEffect, useState } from "react";

const styles = {
  container: css`
    display: flex;
  `,
  bold: css`
    font-weight: bold;
  `,
};

const VersionDisplay: FC = () => {
  const [version, setVersion] = useState("loading...");

  useEffect(() => {
    axios
      .get<string>("/api/version")
      .then(({ data }) => {
        setVersion(data);
      })
      .catch(() => {
        setVersion("unknown");
        alert("Failed to get version");
      });
  }, []);

  return (
    <div css={styles.container}>
      ISUCON version:&nbsp;
      <div css={styles.bold}>{version}</div>
    </div>
  );
};

export default VersionDisplay;
