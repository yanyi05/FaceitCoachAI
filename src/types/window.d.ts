export {};

interface ParseResult {
  success: boolean;
  message?: string;
  match?: {
    map: string;
    tickRate: number;
    scoreCT: number;
    scoreT: number;
    winner: string;
  };
  players: {
    name: string;
    steamId: number;
    team: string;
    teamGroup: string;
    playerId: number;
    isBot: boolean;
  }[];
  teams?: {
    id: string;
    name: string;
    startSide: string;
    players: {
      name: string;
      steamId: number;
      playerId: number;
    }[];
  }[];
  kills?: unknown[];
  damages?: unknown[];
  rounds?: unknown[];
  stats?: unknown[];
  analysis?: {
    support?: unknown[];
    trades?: unknown[];
    rotations?: unknown[];
  };
  facts?: {
    match: unknown;
    teams: unknown[];
    players: unknown[];
    rounds: unknown[];
    stats: unknown[];
    combat: {
      kills: unknown[];
      damages: unknown[];
      weaponFires: unknown[];
      shots: unknown[];
      trades: unknown[];
    };
    positioning: {
      playerStates: unknown[];
      support: unknown[];
      rotations: unknown[];
    };
    utility: {
      grenades: unknown[];
      flashes: unknown[];
      smokes: unknown[];
      molotovs: unknown[];
    };
    economy: {
      rounds: unknown[];
    };
    objective: {
      bombEvents: unknown[];
    };
  };
}

declare global {
  interface Window {
    api: {
      dialog: {
        openDemo(): Promise<string | null>;
      };

      parser: {
        parseDemo(demoPath: string): Promise<ParseResult>;
      };
    };
  }
}
