import { css } from "@emotion/react";
import axios from "axios";
import { FC, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { AppDispatch, RootState } from "../../../store";
import { registerLogs } from "../../../store/logs";
import { TimeToStr } from "../../../utils/time";
import { Log } from "../../../utils/types";
import LogList from "./LogList";

const styles = {
  container: css`
    margin-top: 1em;
    align-items: center;
    display: flex;
    flex-direction: column;
  `,
  title: css`
    line-height: min(10px, calc(2vw - 10px));
  `,
  codeBlock: css`
    margin: 1em 0;
    padding: 1em;
    border-radius: 10px;
    background: #25292f;
    color: #ffffff;
    white-space: pre;
    font-family: source-code-pro, Menlo, Monaco, Consolas, "Courier New",
      monospace;
    font-size: 15px;
    text-align: left;
    max-width: min(500px, 90vw);
    max-height: min(500px, 60vh);
    overflow: auto;
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

const Logs: FC = () => {
  const [selectedLogIndex, setSelectedLogIndex] = useState(0);

  const logs = useSelector((state: RootState) => state.logs);
  const dispatch = useDispatch<AppDispatch>();

  useEffect(() => {
    axios
      .get<Log[]>("/api/logs")
      .then(({ data }) => {
        dispatch(registerLogs(data));
      })
      .catch(() => {
        alert("ログの取得に失敗しました");
      });
  }, []);

  const selectedLog = logs[selectedLogIndex];

  return (
    <div css={styles.container}>
      <h2 css={styles.title}>Logs</h2>
      {logs.length ? (
        <LogList setSelectedLogIndex={setSelectedLogIndex} />
      ) : (
        "No Logs"
      )}
      {selectedLog && (
        <>
          <h3 css={styles.title}>Selected Log</h3>
          Time: {TimeToStr(selectedLog.created_at)}
          <br />
          <pre css={styles.codeBlock}>{selectedLog.stdout}</pre>
        </>
      )}
    </div>
  );
};

export default Logs;
