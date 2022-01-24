import { VFC } from "react";
import MeDisplay from "./Me";
import VersionDisplay from "./Version";

const App: VFC = () => {
  return (
    <>
      <MeDisplay />
      <VersionDisplay />
    </>
  );
};

export default App;
