import { createContext, useContext } from "react";
import { Game } from "src/types/game";

export interface GameManagerContextValue {
  ws: WebSocket;
  game: Game;
  currentPlayer: string;
}
export const GameManagerContext = createContext<
  GameManagerContextValue | undefined
>(undefined);

export const useGameManagerContext = () => {
  const context = useContext(GameManagerContext);

  if (!context) {
    throw new Error(
      "useGameManagerContext must be used inside a GameManagerContext.Provider"
    );
  }

  return context;
};
