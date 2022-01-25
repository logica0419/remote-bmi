import { css } from "@emotion/react";
import { Dispatch, SetStateAction, VFC } from "react";

const styles = {
  background: css`
    background-color: #000000;
    opacity: 0.5;
    width: 100vw;
    height: 100vh;
    position: absolute;
  `,
  modal: css`
    background-color: #ffffff;
    width: 400px;
    height: 200px;
    border-radius: 20px;
    opacity: 1;
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    margin: auto;
    font-size: 20px;
  `,
  button: (color: string) => {
    return css`
      margin: 0 0.2em;
      font-size: 25px;
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
  resetServers: () => void;
  setIsConfirming: Dispatch<SetStateAction<boolean>>;
}

const ConfirmModal: VFC<Props> = ({ resetServers, setIsConfirming }) => {
  const onConfirm = () => {
    resetServers();
    setIsConfirming(false);
  };

  const onCancel = () => {
    setIsConfirming(false);
  };

  return (
    <>
      <div css={styles.background} />
      <div css={styles.modal}>
        <h2>
          本当にサーバー設定を
          <br />
          リセットしますか？
        </h2>
        <button css={styles.button("#bce4c9")} onClick={onConfirm}>
          Confirm
        </button>
        <button css={styles.button("#fcb0b0")} onClick={onCancel}>
          Cancel
        </button>
      </div>
    </>
  );
};

export default ConfirmModal;
