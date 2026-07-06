import { contextBridge, ipcRenderer } from "electron";

contextBridge.exposeInMainWorld("api", {
  dialog: {
    openDemo: () => ipcRenderer.invoke("dialog:open-demo"),
  },

  parser: {
    parseDemo: (demoPath: string) =>
      ipcRenderer.invoke("parser:parse-demo", demoPath),
  },
});