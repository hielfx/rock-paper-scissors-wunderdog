/**
 * The game object
 */
export interface Game {
  id: string;
  roundsNumber: number;
  players: Record<string, Player>;
  rounds: Array<Round>;
  status: string;
  currentRound: number;
  winner: string;
  isDraw: boolean;
  pausedBy: string;
}

/**
 * The player object
 */
export interface Player {
  nickname: string;
}

/**
 * Round information
 */
export interface Round {
  /**
   * Map of nickname - move so we know what the players choose
   */
  playerSelections: Record<string, Move> | null;
  /**
   * The player that won the round
   */
  winner: string;
  isDraw: boolean;
}

/**
 * Selected move from the players
 */
export type Move = "ROCK" | "PAPER" | "SCISSORS";

export const Moves: Record<string, Move> = {
  rock: "ROCK",
  paper: "PAPER",
  scissors: "SCISSORS",
};
