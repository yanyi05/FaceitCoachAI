import { spawn } from "child_process";
import path from "path";
import { app } from "electron";

export interface ParseResult {
  success: boolean;
  message: string;
  match?: {
    map: string;
    tickRate: number;
    scoreCT: number;
    scoreT: number;
    winner: string;
  };
  players?: unknown[];
  kills?: unknown[];
  damages?: unknown[];
  rounds?: unknown[];
  stats?: unknown[];
  analysis?: {
    trades?: unknown[];
    rotations?: unknown[];
  };
}

function getParserPath() {
  if (app.isPackaged) {
    return path.join(process.resourcesPath, "parser", "parser.exe");
  }

  return path.join(process.env.APP_ROOT ?? process.cwd(), "parser", "parser.exe");
}

function parseParserOutput(stdout: string): ParseResult {
  const trimmed = stdout.trim();

  try {
    return JSON.parse(trimmed);
  } catch {
    const jsonStart = trimmed.lastIndexOf("\n{");
    const candidate = jsonStart >= 0 ? trimmed.slice(jsonStart + 1) : trimmed;

    return JSON.parse(candidate);
  }
}

export function parseDemo(demoPath: string): Promise<ParseResult> {
  return new Promise((resolve, reject) => {
    const parserPath = getParserPath();

    const child = spawn(parserPath, ["event", demoPath], {
      cwd: path.dirname(parserPath),
    });

    let stdout = "";
    let stderr = "";

    child.stdout.on("data", (data) => {
      stdout += data.toString();
    });

    child.stderr.on("data", (data) => {
      stderr += data.toString();
    });

    child.on("error", (err) => {
      reject(err);
    });

    child.on("close", (code) => {
      if (code !== 0) {
        resolve({
          success: false,
          message: stderr || `Parser exited with code ${code}`,
        });
        return;
      }

      try {
        resolve(parseParserOutput(stdout));
      } catch (err) {
        resolve({
          success: false,
          message:
            err instanceof Error
              ? `Could not read parser result: ${err.message}`
              : "Could not read parser result.",
        });
      }
    });
  });
}
