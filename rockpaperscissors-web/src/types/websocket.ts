import { Game as GameType } from "./game";

/**
 * Websocket payload
 */
export interface Payload {
  nickname: string;
  gameId: string;
  game: GameType;
  // value: number;
  command: Command;
  message: string;
}

/**
 * Command to send to the websocket server so it can manage the game status
 */
export type Command = "connect" | "play" | "ok" | "error" | "close";
