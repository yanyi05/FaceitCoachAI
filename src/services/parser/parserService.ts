export interface Player {
  name: string;
  steamId: number;
  team: string;
  teamGroup: string;
  playerId: number;
  isBot: boolean;
}

export interface TeamMember {
  name: string;
  steamId: number;
  playerId: number;
}

export interface TeamGroup {
  id: string;
  name: string;
  startSide: string;
  players: TeamMember[];
}

export interface MatchInfo {
  map: string;
  tickRate: number;
  scoreCT: number;
  scoreT: number;
  winner: string;
}

export interface Kill {
  tick: number;
  round: number;
  killer: string;
  killerTeam: string;
  victim: string;
  victimTeam: string;
  assister: string;
  weapon: string;
  headshot: boolean;
}

export interface PlayerStats {
  name: string;
  steamId: number;
  kills: number;
  deaths: number;
  assists: number;
  headshots: number;
  adr: number;
  kast: number;
  rating: number;
  accuracy: number;
  headshotPercent: number;
  damage: number;
  entryKills: number;
  entryDeaths: number;
  averageTTD: number;
  tradeKills: number;
  tradeDeaths: number;
  tradeSuccess: number;
}

export interface TradeAnalysis {
  round: number;
  tick: number;
  victim: string;
  killer: string;
  trader: string;
  expectedTrader: string;
  success: boolean;
  tradeTimeTicks: number;
  supportDistance: number;
  heightDifference: number;
  tradeQuality: string;
}

export interface Facts {
  match: MatchInfo;
  teams: TeamGroup[];
  players: Player[];
  rounds: unknown[];
  stats: PlayerStats[];
  combat: {
    kills: Kill[];
    damages: unknown[];
    weaponFires: unknown[];
    shots: unknown[];
    trades: TradeAnalysis[];
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
}

export interface ParseResult {
  success: boolean;
  match?: MatchInfo;
  players: Player[];
  teams?: TeamGroup[];
  kills?: Kill[];
  damages?: unknown[];
  rounds?: unknown[];
  stats?: PlayerStats[];
  analysis?: {
    support?: unknown[];
    trades?: TradeAnalysis[];
    rotations?: unknown[];
  };
  facts?: Facts;
  message?: string;
}

export async function parseDemo(
  demoPath: string
): Promise<ParseResult> {
  return (await window.api.parser.parseDemo(demoPath)) as ParseResult;
}
