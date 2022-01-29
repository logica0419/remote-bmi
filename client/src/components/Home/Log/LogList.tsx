import { css } from "@emotion/react";
import { Dispatch, FormEvent, SetStateAction, VFC } from "react";
import { useSelector } from "react-redux";
import { RootState } from "../../../store";
import { TimeToStr } from "../../../utils/time";
import { Log } from "../../../utils/types";

const styles = {
  container: css`
    max-height: min(25vh, 30vw);
    overflow-y: auto;
  `,
  table: css`
    table-layout: fixed;
    border-collapse: collapse;
  `,
  th: css`
    padding: 0 10px;
    border: solid 2px;
  `,
  td: css`
    padding: 0 10px;
    border: solid 1px;
  `,
  button: (color: string) => {
    return css`
      font-size: min(20px, calc(2vw));
      border: 0;
      border-radius: 0.4em;
      padding: 0.2em 0.4em;
      background: ${color};
      :hover {
        opacity: 0.6;
        cursor: pointer;
      }
    `;
  },
};

interface Props {
  setSelectedLog: Dispatch<SetStateAction<Log | undefined>>;
}

const LogList: VFC<Props> = ({ setSelectedLog }) => {
  const logs = useSelector((state: RootState) => state.logs);

  const onSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const target = e.currentTarget;
    const id = target.name;

    const log = logs.find((log) => log.id === id);
    setSelectedLog(log);
  };

  return (
    <div css={styles.container}>
      <table css={styles.table}>
        <thead>
          <tr>
            <th css={styles.th}>Time</th>
            <th css={styles.th}>Server</th>
            <th css={styles.th}>Detail</th>
          </tr>
        </thead>
        <tbody>
          {logs.map((log) => {
            return (
              <tr key={log.id}>
                <td css={styles.td}>{TimeToStr(log.created_at)}</td>
                <td css={styles.td}>{log.server_number}</td>
                <td css={styles.td}>
                  <form name={log.id} onSubmit={onSubmit}>
                    <input
                      css={styles.button("#e0e0e0")}
                      type="submit"
                      value="See Log"
                    />
                  </form>
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
};

export default LogList;
