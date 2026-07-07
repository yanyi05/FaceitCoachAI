import { useState } from "react";
import styles from "./Workspace.module.css";
import Button from "../../components/ui/Button/Button";
import { getCurrentDemo } from "../../services/demo/demoService";
import { parseDemo } from "../../services/parser/parserService";

function Workspace() {
  const demo = getCurrentDemo();

  const [status, setStatus] = useState("Ready");
  const [parsing, setParsing] = useState(false);
  const [players, setPlayers] = useState<{ name: string }[]>([]);

  async function handleParseDemo() {
    if (!demo) {
      return;
    }

    setParsing(true);
    setStatus("Parsing...");

    const result = await parseDemo(demo.path);

    console.log(result);

    if (result.success) {
      setPlayers(result.players);
      setStatus("Parsed");
    } else {
      setStatus("Failed");
    }

    setParsing(false);
  }

  if (!demo) {
    return (
      <div className={styles.workspace}>
        <h1>Workspace</h1>

        <div className={styles.card}>
          <h2>No Demo Loaded</h2>
          <p>Please import a CS2 demo from the Dashboard.</p>
        </div>
      </div>
    );
  }

  return (
    <div className={styles.workspace}>
      <h1>Workspace</h1>

      <div className={styles.card}>
        <h2>Demo Information</h2>

        <p><strong>Demo Name</strong></p>
        <p>{demo.fileName}</p>

        <hr />

        <p><strong>Full Path</strong></p>
        <p>{demo.path}</p>

        <hr />

        <p><strong>Status</strong></p>
        <p>{status}</p>

        <br />

        <Button onClick={handleParseDemo} disabled={parsing}>
          {parsing ? "Parsing..." : "Parse Demo"}
        </Button>
        <br />
        <br />

        <h3>Players</h3>

        {players.map((player) => (
          <p key={player.name}>{player.name}</p>
        ))}
      </div>
    </div>
  );
}

export default Workspace;