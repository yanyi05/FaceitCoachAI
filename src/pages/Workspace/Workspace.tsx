import { useState } from "react";
import styles from "./Workspace.module.css";
import Button from "../../components/ui/Button/Button";
import { getCurrentDemo } from "../../services/demo/demoService";
import {
  ParseResult,
  parseDemo,
} from "../../services/parser/parserService";

function Workspace() {
  const demo = getCurrentDemo();

  const [status, setStatus] = useState("Ready");
  const [parsing, setParsing] = useState(false);
  const [result, setResult] = useState<ParseResult | null>(null);

  const stats = result?.stats ?? [];
  const kills = result?.kills ?? [];
  const trades = result?.analysis?.trades ?? [];
  const teams = result?.teams ?? [];
  const facts = result?.facts;

  const dataCategories = facts
    ? [
        { label: "Match", count: 1 },
        { label: "Teams", count: facts.teams.length },
        { label: "Players", count: facts.players.length },
        { label: "Rounds", count: facts.rounds.length },
        { label: "Kills", count: facts.combat.kills.length },
        { label: "Damages", count: facts.combat.damages.length },
        { label: "Shots", count: facts.combat.shots.length },
        { label: "Trades", count: facts.combat.trades.length },
        { label: "Support", count: facts.positioning.support.length },
        { label: "Positions", count: facts.positioning.playerStates.length },
        { label: "Rotations", count: facts.positioning.rotations.length },
        { label: "Utility", count: facts.utility.grenades.length },
        { label: "Economy", count: facts.economy.rounds.length },
        { label: "Objective", count: facts.objective.bombEvents.length },
      ]
    : [];

  async function handleParseDemo() {
    if (!demo) {
      return;
    }

    setParsing(true);
    setStatus("Parsing...");
    setResult(null);

    const result = await parseDemo(demo.path);

    if (result.success) {
      setResult(result);
      setStatus("Parsed successfully");
    } else {
      setStatus(result.message ?? "Failed");
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

        <p><strong>Demo name</strong></p>
        <p>{demo.fileName}</p>

        <hr />

        <p><strong>Full path</strong></p>
        <p>{demo.path}</p>

        <hr />

        <p><strong>Status</strong></p>
        <p>{status}</p>

        <br />

        <Button onClick={handleParseDemo} disabled={parsing}>
          {parsing ? "Parsing..." : "Parse Demo"}
        </Button>
      </div>

      {result?.success && (
        <>
          <section className={styles.summaryGrid}>
            <div className={styles.metric}>
              <span>Players</span>
              <strong>{result.players.length}</strong>
            </div>

            <div className={styles.metric}>
              <span>Rounds</span>
              <strong>{result.rounds?.length ?? 0}</strong>
            </div>

            <div className={styles.metric}>
              <span>Kills</span>
              <strong>{kills.length}</strong>
            </div>

            <div className={styles.metric}>
              <span>Trades</span>
              <strong>{trades.length}</strong>
            </div>
          </section>

          <section className={styles.card}>
            <h2>Data categories</h2>
            <div className={styles.categoryGrid}>
              {dataCategories.map((category) => (
                <div className={styles.categoryItem} key={category.label}>
                  <span>{category.label}</span>
                  <strong>{category.count}</strong>
                </div>
              ))}
            </div>
          </section>

          <section className={styles.card}>
            <h2>Stable teams</h2>
            {teams.length === 2 ? (
              <div className={styles.teamGrid}>
                {teams.map((team) => (
                  <div className={styles.teamPanel} key={team.id}>
                    <div className={styles.teamHeader}>
                      <strong>{team.name}</strong>
                      <span>Started {team.startSide}</span>
                    </div>

                    <div className={styles.playerList}>
                      {team.players.map((player) => (
                        <div className={styles.playerPill} key={player.steamId}>
                          {player.name}
                        </div>
                      ))}
                    </div>
                  </div>
                ))}
              </div>
            ) : (
              <p>Could not detect two complete 5-player teams from this demo.</p>
            )}
          </section>

          <section className={styles.card}>
            <h2>Match overview</h2>
            <div className={styles.detailGrid}>
              <div>
                <span>Tick rate</span>
                <strong>{result.match?.tickRate ?? "-"}</strong>
              </div>
              <div>
                <span>Score</span>
                <strong>
                  CT {result.match?.scoreCT ?? 0} : {result.match?.scoreT ?? 0} T
                </strong>
              </div>
              <div>
                <span>Winner</span>
                <strong>{result.match?.winner ?? "-"}</strong>
              </div>
            </div>
          </section>

          <section className={styles.card}>
            <h2>Player stats</h2>
            <div className={styles.tableWrap}>
              <table>
                <thead>
                  <tr>
                    <th>Player</th>
                    <th>K</th>
                    <th>D</th>
                    <th>ADR</th>
                    <th>HS%</th>
                    <th>ACC%</th>
                    <th>Entry</th>
                  </tr>
                </thead>
                <tbody>
                  {stats.map((player) => (
                    <tr key={player.steamId}>
                      <td>{player.name}</td>
                      <td>{player.kills}</td>
                      <td>{player.deaths}</td>
                      <td>{player.adr.toFixed(1)}</td>
                      <td>{player.headshotPercent.toFixed(1)}</td>
                      <td>{player.accuracy.toFixed(1)}</td>
                      <td>
                        {player.entryKills}/{player.entryDeaths}
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </section>

          <section className={styles.card}>
            <h2>Recent kills</h2>
            <div className={styles.eventList}>
              {kills.slice(-12).reverse().map((kill) => (
                <div
                  className={styles.eventRow}
                  key={`${kill.round}-${kill.tick}-${kill.killer}-${kill.victim}`}
                >
                  <span>R{kill.round}</span>
                  <strong>{kill.killer}</strong>
                  <span>killed</span>
                  <strong>{kill.victim}</strong>
                  <span>{kill.weapon}{kill.headshot ? " / HS" : ""}</span>
                </div>
              ))}
            </div>
          </section>

          <section className={styles.card}>
            <h2>Trade facts</h2>
            <div className={styles.eventList}>
              {trades.slice(0, 12).map((trade) => (
                <div
                  className={styles.eventRow}
                  key={`${trade.round}-${trade.tick}-${trade.victim}-${trade.trader}`}
                >
                  <span>R{trade.round}</span>
                  <strong>{trade.trader || "-"}</strong>
                  <span>traded</span>
                  <strong>{trade.killer}</strong>
                  <span>{trade.tradeTimeTicks} ticks</span>
                </div>
              ))}
              {trades.length === 0 && (
                <p>No trade facts were found in this demo.</p>
              )}
            </div>
          </section>
        </>
      )}
    </div>
  );
}

export default Workspace;
