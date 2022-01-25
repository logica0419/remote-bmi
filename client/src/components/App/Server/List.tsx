import { css } from "@emotion/react";
import { VFC } from "react";
import { Server } from "../../../utils/types";

const styles = {
  table: css`
    table-layout: fixed;
    border-collapse: collapse;s
  `,
  th: css`
    padding: 0 10px;
    border: solid 2px;
  `,
  td: css`
    padding: 0 10px;
    border: solid 1px;
  `,
};

interface Props {
  servers: Server[];
}

const ServerList: VFC<Props> = ({ servers }) => {
  return (
    <table css={styles.table}>
      <tr>
        <th css={styles.th}>Number</th>
        <th css={styles.th}>Private Address</th>
        <th css={styles.th}>BenchMark</th>
      </tr>
      {servers.map((server: Server) => {
        return (
          <tr>
            <td css={styles.td}>{server.server_number}</td>
            <td css={styles.td}>{server.address}</td>
            <td css={styles.td}></td>
          </tr>
        );
      })}
    </table>
  );
};

export default ServerList;
