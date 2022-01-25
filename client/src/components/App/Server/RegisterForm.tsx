import { css } from "@emotion/react";
import { ChangeEvent, Dispatch, SetStateAction, VFC } from "react";
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
  input: css`
    font-size: min(20px, calc(2vw));
  `,
};

interface Props {
  editingServers: Server[];
  setEditingServers: Dispatch<SetStateAction<Server[]>>;
}

const RegisterForm: VFC<Props> = ({ editingServers, setEditingServers }) => {
  const onChange = (e: ChangeEvent<HTMLInputElement>) => {
    const target = e.target;
    const value = target.value;
    const name = target.name;

    let newServers = [...editingServers];
    newServers[Number(name)].address = value;
    setEditingServers(newServers);
  };

  return (
    <>
      <table css={styles.table}>
        <thead>
          <tr>
            <th css={styles.th}>Number</th>
            <th css={styles.th}>Private Address</th>
          </tr>
        </thead>
        <tbody>
          {editingServers.map((editingServer: Server) => {
            return (
              <tr key={editingServer.server_number}>
                <td css={styles.td}>{editingServer.server_number}</td>
                <td css={styles.td}>
                  <input
                    css={styles.input}
                    name={(editingServer.server_number - 1).toString(10)}
                    type="text"
                    value={editingServer.address}
                    onChange={onChange}
                  />
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </>
  );
};

export default RegisterForm;
