import { ipcMain } from "electron";

ipcMain.handle("parser:parse-demo", async (_, demoPath: string) => {
  console.log("Parsing Demo:", demoPath);

  return {
    success: true,
    map: "Mirage",
    tickRate: 128,
    duration: "35:42",
    players: [
      "donk",
      "m0NESY",
      "NiKo",
      "ropz",
      "ZywOo",
    ],
  };
});