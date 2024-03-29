import { css } from "@emotion/react";
import { FC } from "react";
import { Server } from "../../../utils/types";

const styles = {
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
};

interface Props {
  servers: Server[];
}

const ServerList: FC<Props> = ({ servers }) => {
  return (
    <table css={styles.table}>
      <thead>
        <tr>
          <th css={styles.th}>Number</th>
          <th css={styles.th}>Private Address</th>
        </tr>
      </thead>
      <tbody>
        {servers.map((server: Server) => {
          return (
            <tr key={server.server_number}>
              <td css={styles.td}>{server.server_number}</td>
              <td css={styles.td}>{server.address}</td>
            </tr>
          );
        })}
      </tbody>
    </table>
  );
};

export default ServerList;
