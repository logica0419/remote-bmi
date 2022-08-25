import { FC } from "react";
import Benchmark from "./Benchmark";
import Logs from "./Log";
import MeDisplay from "./Me";
import ServerContainer from "./Server";
import VersionDisplay from "./Version";

const Home: FC = () => {
  return (
    <>
      <MeDisplay />
      <VersionDisplay />
      <ServerContainer />
      <Benchmark />
      <Logs />
      <br />
    </>
  );
};

export default Home;
