import { dialog, ipcMain } from "electron";

ipcMain.handle("dialog:open-demo", async () => {
  const result = await dialog.showOpenDialog({
    title: "Select CS2 Demo",
    properties: ["openFile"],
    filters: [
      {
        name: "CS2 Demo",
        extensions: ["dem"],
      },
    ],
  });

  if (result.canceled) {
    return null;
  }

  return result.filePaths[0];
});