import { css } from "@emotion/react";
import axios from "axios";
import { useState, VFC } from "react";
import { useSelector } from "react-redux";
import { RootState } from "../../../store";
import { PostBenchmarkResponse } from "../../../utils/types";
import Selector from "./Selector";

const styles = {
  container: css`
    margin-top: 1em;
    align-items: center;
  `,
  title: css`
    line-height: min(10px, calc(2vw - 10px));
  `,
  button: (color: string) => {
    return css`
      margin: 0.5em 0.2em 0;
      font-size: min(25px, calc(5px + 2vw));
      border: 0;
      border-radius: 0.4em;
      padding: 0.1em 0.3em;
      background: ${color};
      :hover {
        opacity: 0.6;
        cursor: pointer;
      }
    `;
  },
};

const Benchmark: VFC = () => {
  const servers = useSelector((state: RootState) => state.servers);

  const [serverNumber, setServerNumber] = useState(1);

  const onBenchmark = () => {
    axios
      .post<PostBenchmarkResponse>("/api/benchmark", {
        server_number: serverNumber,
      })
      .then(({ data }) => {
        console.log(data);
      })
      .catch((err) => {
        alert("benchmark failed");
      });
  };

  return (
    <div css={styles.container}>
      <h2 css={styles.title}>Benchmark</h2>
      {servers.length ? (
        <>
          Server:&nbsp;
          <Selector
            serverNumber={serverNumber}
            setServerNumber={setServerNumber}
          />
          &nbsp;
          <button css={styles.button("#bce4c9")} onClick={onBenchmark}>
            Benchmark
          </button>
        </>
      ) : (
        "Currently Unavailable"
      )}
    </div>
  );
};

export default Benchmark;
