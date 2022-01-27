import { css } from "@emotion/react";
import axios from "axios";
import {
  Dispatch,
  SetStateAction,
  VFC,
  ChangeEvent,
  FormEvent,
  useState,
} from "react";
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
  editingServers: Server[];
  setEditingServers: Dispatch<SetStateAction<Server[]>>;
}

const EditForm: VFC<Props> = ({ editingServers, setEditingServers }) => {
  const [isEdited] = useState(new Map<number, boolean>());

  const onChange = (e: ChangeEvent<HTMLInputElement>) => {
    const target = e.currentTarget;
    const value = target.value;
    const index = Number(target.name);

    isEdited.set(index, true);

    let newServers = [...editingServers];
    newServers[index].address = value;
    setEditingServers(newServers);
  };

  const onSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const target = e.currentTarget;
    const index = Number(target.name);
    const serverNumber = index + 1;

    if (!isEdited.has(index)) return;

    axios
      .put(`/api/servers/${serverNumber}`, editingServers[index].address)
      .then(() => {
        alert("正常にアドレスがアップデートされました");
        isEdited.delete(index);
      })
      .catch(() => {
        alert("アップデートに失敗しました");
      });
  };

  return (
    <table css={styles.table}>
      <thead>
        <tr>
          <th css={styles.th}>Number</th>
          <th css={styles.th}>Private Address</th>
          <th css={styles.th}>Update</th>
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
              <td css={styles.td}>
                <form
                  name={(editingServer.server_number - 1).toString(10)}
                  onSubmit={onSubmit}>
                  <input
                    css={styles.button("#bce4c9")}
                    type="submit"
                    value="Update"
                  />
                </form>
              </td>
            </tr>
          );
        })}
      </tbody>
    </table>
  );
};

export default EditForm;
