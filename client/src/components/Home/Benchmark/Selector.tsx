import { css } from "@emotion/react";
import { Dispatch, SetStateAction, VFC } from "react";
import { useSelector } from "react-redux";
import { RootState } from "../../../store";

const styles = {
  selector: css`
    font-size: min(25px, calc(5px + 2vw));
    border-radius: 0.5em;
  `,
};

interface Props {
  serverNumber: number;
  setServerNumber: Dispatch<SetStateAction<number>>;
}

const Selector: VFC<Props> = ({ serverNumber, setServerNumber }) => {
  const servers = useSelector((state: RootState) => state.servers);

  const onChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setServerNumber(Number(e.target.value));
  };

  return (
    <select css={styles.selector} value={serverNumber} onChange={onChange}>
      {servers.map((server) => (
        <option key={server.id} value={server.server_number}>
          {server.server_number}
        </option>
      ))}
    </select>
  );
};

export default Selector;
