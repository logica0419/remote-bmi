import { css } from "@emotion/react";
import { Dispatch, SetStateAction, VFC } from "react";
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
    margin: 0.1em 0;
  `,
  button: (color: string) => {
    return css`
      margin: 0.5em 0.2em;
      font-size: min(25px, calc(5px + 2vw));
      border: 0;
      border-radius: 0.4em;
      padding: 0.3em 0.5em;
      background: ${color};
      :hover {
        opacity: 0.6;
        cursor: pointer;
      }
    `;
  },
};

interface Props {
  editingServers: Server[];
  setEditingServers: Dispatch<SetStateAction<Server[]>>;
}

const RegisterForm: VFC<Props> = ({ editingServers, setEditingServers }) => {
  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
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
        <tr>
          <th css={styles.th}>Number</th>
          <th css={styles.th}>Private Address</th>
        </tr>
        {editingServers.map((editingServer: Server) => {
          return (
            <tr>
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
      </table>
    </>
  );
};

export default RegisterForm;
