import { spawn } from "child_process";
import path from "path";

export interface ParseResult {
  success: boolean;
  message: string;
  path?: string;
}

export function parseDemo(demoPath: string): Promise<ParseResult> {
  return new Promise((resolve, reject) => {
    const parserPath = path.join(
      process.cwd(),
      "parser",
      "parser.exe"
    );

    const child = spawn(parserPath, [demoPath]);

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
        reject(new Error(stderr));
        return;
      }

      resolve(JSON.parse(stdout));
    });
  });
}