import {
  getCurrentDemo,
  setCurrentDemo,
} from "./currentDemo";

export async function openDemo() {
  const path = await window.api.dialog.openDemo();

  if (!path) {
    return null;
  }

  const fileName = path.split("\\").pop() ?? "";

  const demo = {
    path,
    fileName,
  };

  setCurrentDemo(demo);

  return demo;
}

export { getCurrentDemo };