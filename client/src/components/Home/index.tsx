import { VFC } from "react";
import Benchmark from "./Benchmark";
import Logs from "./Log";
import MeDisplay from "./Me";
import ServerContainer from "./Server";
import VersionDisplay from "./Version";

const Home: VFC = () => {
  return (
    <>
      <MeDisplay />
      <VersionDisplay />
      <ServerContainer />
      <Benchmark />
      <Logs />
    </>
  );
};

export default Home;
