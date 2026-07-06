import { ipcMain, dialog } from "electron";
ipcMain.handle("dialog:open-demo", async () => {
  const result = await dialog.showOpenDialog({
    title: "Select CS2 Demo",
    properties: ["openFile"],
    filters: [
      {
        name: "CS2 Demo",
        extensions: ["dem"]
      }
    ]
  });
  if (result.canceled) {
    return null;
  }
  return result.filePaths[0];
});
ipcMain.handle("parser:parse-demo", async (_, demoPath) => {
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
      "ZywOo"
    ]
  };
});
