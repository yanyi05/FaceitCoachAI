import { ipcMain } from "electron";
import { parseDemo } from "../services/parserService";

ipcMain.handle(
  "parser:parse-demo",
  async (_event, demoPath: string) => {
    return await parseDemo(demoPath);
  }
);