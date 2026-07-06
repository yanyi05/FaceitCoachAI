export {};

declare global {
  interface Window {
    api: {
      dialog: {
        openDemo(): Promise<string | null>;
      };

      parser: {
        parseDemo(demoPath: string): Promise<{
          success: boolean;
          map: string;
          tickRate: number;
          duration: string;
          players: string[];
        }>;
      };
    };
  }
}