/// <reference types="vite/client" />
interface Window {
  api: {
    dialog: {
      openDemo: () => Promise<string | null>;
    };
  };
}