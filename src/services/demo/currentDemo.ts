export interface DemoInfo {
  path: string;
  fileName: string;
}

let currentDemo: DemoInfo | null = null;

export function getCurrentDemo() {
  return currentDemo;
}

export function setCurrentDemo(demo: DemoInfo) {
  currentDemo = demo;
}

export function clearCurrentDemo() {
  currentDemo = null;
}