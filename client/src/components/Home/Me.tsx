import { css } from "@emotion/react";
import { FC } from "react";
import { useSelector } from "react-redux";
import { RootState } from "../../store";

const styles = {
  container: css`
    display: flex;
  `,
  bold: css`
    font-weight: bold;
  `,
};

const MeDisplay: FC = () => {
  const myName = useSelector((state: RootState) => state.me.name);

  return (
    <div css={styles.container}>
      Welcome,&nbsp;
      <div css={styles.bold}>{myName}</div>
    </div>
  );
};

export default MeDisplay;
