import { css } from "@emotion/react";
import { ChangeEvent, Dispatch, FC, FormEvent, SetStateAction } from "react";

const styles = {
  input: css`
    font-size: min(20px, calc(2vw));
  `,
  button: (color: string) => {
    return css`
      font-size: min(22px, calc(2vw + 2px));
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

type Props = {
  action: string;
  onSubmit: (e: FormEvent<HTMLFormElement>) => void;
  name: string;
  setName: Dispatch<SetStateAction<string>>;
};

const Form: FC<Props> = ({ action, onSubmit, name, setName }) => {
  const onChange = (e: ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setName(value);
  };

  return (
    <form onSubmit={onSubmit}>
      <label>
        Name:&nbsp;
        <input
          css={styles.input}
          type="text"
          value={name}
          onChange={onChange}
        />
      </label>
      &nbsp;
      <input css={styles.button("#e0e0e0")} type="submit" value={action} />
    </form>
  );
};

export default Form;
