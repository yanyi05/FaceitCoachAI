"use strict";
const electron = require("electron");
electron.contextBridge.exposeInMainWorld("api", {
  dialog: {
    openDemo: () => electron.ipcRenderer.invoke("dialog:open-demo")
  },
  parser: {
    parseDemo: (demoPath) => electron.ipcRenderer.invoke("parser:parse-demo", demoPath)
  }
});
