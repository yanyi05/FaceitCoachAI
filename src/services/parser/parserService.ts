export interface Player {
  name: string;
}

export interface ParseResult {
  success: boolean;
  players: Player[];
  message?: string;
}

export async function parseDemo(
  demoPath: string
): Promise<ParseResult> {
  return await window.api.parser.parseDemo(demoPath);
}