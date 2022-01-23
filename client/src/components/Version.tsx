import axios from "axios";
import { useEffect, useState, VFC } from "react";

const VersionDisplay: VFC = () => {
  const [version, setVersion] = useState("loading...");

  useEffect(() => {
    axios
      .get<string>("/api/version")
      .then(({ data }) => {
        setVersion(data);
      })
      .catch(() => {
        setVersion("unknown");
        alert("Failed to get version");
      });
  }, []);

  return <>ISUCON version: {version}</>;
};

export default VersionDisplay;
