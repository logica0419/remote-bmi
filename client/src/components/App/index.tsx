import { VFC } from "react";
import MeDisplay from "./Me";
import ServerContainer from "./Server";
import VersionDisplay from "./Version";

const App: VFC = () => {
  return (
    <>
      <MeDisplay />
      <VersionDisplay />
      <ServerContainer />
    </>
  );
};

export default App;
