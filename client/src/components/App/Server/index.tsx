import { css } from "@emotion/react";
import axios from "axios";
import { useEffect, useState, VFC } from "react";
import { PostServersRequest, Server } from "../../../utils/types";
import RegisterForm from "./Register";
import ServerList from "./List";
import EditForm from "./Edit";
import ConfirmModal from "./ConfirmModal";

const styles = {
  container: css`
    margin-top: 1em;
    max-height: min(500px, calc(20vw + 100px));
    align-items: center;
  `,
  title: css`
    line-height: min(10px, calc(2vw - 10px));
  `,
  button: (color: string, isConfirming?: boolean) => {
    return css`
      margin: 0.5em 0.2em 0;
      font-size: min(25px, calc(5px + 2vw));
      border: 0;
      border-radius: 0.4em;
      padding: 0.3em 0.5em;
      background: ${color};
      :hover {
        ${!isConfirming && "opacity: 0.6;"}
        cursor: pointer;
      }
    `;
  },
};

const ServerContainer: VFC = () => {
  const [isFetching, setIsFetching] = useState(true);
  const [editingServers, setEditingServers] = useState<Server[]>([]);
  const [servers, setServers] = useState<Server[]>([]);
  const [isEditing, setIsEditing] = useState(false);
  const [isConfirming, setIsConfirming] = useState(false);

  useEffect(() => {
    setIsFetching(true);

    axios
      .get<Server[]>("/api/servers")
      .then(({ data }) => {
        setServers(data);
        setIsFetching(false);
      })
      .catch(() => {
        setServers([]);
        setIsFetching(false);
      });
  }, []);

  const startEdit = () => {
    setEditingServers(servers);
    setIsEditing(true);
  };

  const startRegister = () => {
    setEditingServers([
      { id: "", server_number: 1, address: "" },
      { id: "", server_number: 2, address: "" },
      { id: "", server_number: 3, address: "" },
    ]);
    setIsEditing(true);
  };

  const finishEdit = () => {
    setIsEditing(false);
  };

  const registerServer = () => {
    let newServers: PostServersRequest[] = [];

    for (const editingServer of editingServers) {
      if (!editingServer.address) {
        break;
      }

      newServers.push({
        server_number: editingServer.server_number,
        address: editingServer.address,
      });
    }

    if (!newServers.length) {
      alert("登録するサーバーがありません");
      return;
    }

    setIsEditing(false);
    setIsFetching(true);

    axios
      .post<Server[]>("/api/servers", newServers)
      .then(({ data }) => {
        setServers(data);
      })
      .catch(() => {
        alert("サーバー設定の作成に失敗しました");
      });

    setIsFetching(false);
  };

  const resetServers = () => {
    setIsFetching(true);

    axios
      .delete("/api/servers")
      .then(() => {
        setServers([]);
      })
      .catch(() => {
        alert("サーバー設定の削除に失敗しました");
      });

    setIsFetching(false);
  };

  const onConfirm = () => {
    setIsConfirming(true);
  };

  return (
    <>
      {isConfirming && (
        <ConfirmModal
          resetServers={resetServers}
          setIsConfirming={setIsConfirming}
        />
      )}
      <div css={styles.container}>
        {isEditing ? (
          !servers.length ? (
            <h2 css={styles.title}>Registering Servers...</h2>
          ) : (
            <h2 css={styles.title}>Editing Servers...</h2>
          )
        ) : (
          <h2 css={styles.title}>Your Servers</h2>
        )}
        {isEditing ? (
          !servers.length ? (
            <>
              <RegisterForm
                editingServers={editingServers}
                setEditingServers={setEditingServers}
              />
              <button css={styles.button("#bce4c9")} onClick={registerServer}>
                Register
              </button>
              <button css={styles.button("#fcb0b0")} onClick={finishEdit}>
                Cancel
              </button>
            </>
          ) : (
            <>
              <EditForm
                editingServers={editingServers}
                setEditingServers={setEditingServers}
              />
              <button css={styles.button("#e0e0e0")} onClick={finishEdit}>
                Quit
              </button>
            </>
          )
        ) : isFetching ? (
          <h2>Loading Servers...</h2>
        ) : servers.length ? (
          <>
            <ServerList servers={servers} />
            <button css={styles.button("#e0e0e0")} onClick={startEdit}>
              Edit
            </button>
            <button
              css={styles.button("#fcb0b0", isConfirming)}
              onClick={onConfirm}>
              Reset Servers
            </button>
          </>
        ) : (
          <>
            No Servers Registered.
            <br />
            <button css={styles.button("#e0e0e0")} onClick={startRegister}>
              Register
            </button>
          </>
        )}
      </div>
    </>
  );
};

export default ServerContainer;
