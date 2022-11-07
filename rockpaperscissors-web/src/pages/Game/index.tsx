import React, { useEffect, useState } from "react";
import {
  Navigate,
  useLocation,
  useParams,
  useNavigate,
} from "react-router-dom";
import { Game as GameType } from "src/types/game";
import { joinGame } from "src/api/game";
import { apiErrorHandler } from "src/api/axios";
import { WS_URL } from "src/constants";
import { Commands } from "src/types/gameCommands";
import GameManager from "./GameManager";
import { Payload } from "src/types/websocket";
import {
  GameManagerContext,
  GameManagerContextValue,
} from "src/context/GameManager";

/**
 * Game component to play
 */
const Game = () => {
  const navigate = useNavigate();
  const { gameId } = useParams();
  const { state } = useLocation();
  console.log("LOCATION: ", state);
  const [game, setGame] = useState<GameType | null>(null);
  const [ws, setWS] = useState<WebSocket | null>(null);

  /**
   * Fetches the game if we didn't receive it from the navigation state
   * (if we joined the game without creating it)
   */
  // useEffect(() => {
  //   if (!game) {
  //     //TODO: GET GAME
  //   } else {
  //     console.log("CURRENT GAME: ", game);
  //   }
  // }, []);

  /**
   * Joins the game
   */
  useEffect(() => {
    console.log("join game");
    joinGame(gameId as string, state.nickname, (_game, error) => {
      if (error) {
        console.log("Error: ", error);
        apiErrorHandler(error);
        navigate("/");
      }
      if (_game) {
        console.log("_gamne: ", _game);
        setGame(_game);
        setWS(new WebSocket(WS_URL));
      }
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [gameId]);

  /**
   * Connects with the webook
   */
  useEffect(() => {
    if (ws) {
      ws.onopen = function (this: WebSocket) {
        this.send(
          JSON.stringify({
            gameId,
            nickname: state.nickname,
            command: Commands.connect,
          })
        );
      };
      ws.onclose = function (this: WebSocket) {
        this.send(
          JSON.stringify({
            gameId,
            nickname: state.nickname,
            command: Commands.close,
          })
        );
      };
      ws.onmessage = function (this: WebSocket, event: MessageEvent<string>) {
        const payload: Payload = JSON.parse(event.data);
        setGame(payload.game);
      };

      return () => {
        ws.close();
      };
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [ws]);

  const gameManagerContextValue: GameManagerContextValue = {
    ws: ws as WebSocket,
    currentPlayer: state?.nickname,
    game: game as GameType,
  };

  return !gameId || !state?.nickname ? (
    <Navigate to="/" replace={true} />
  ) : !game ? (
    <div>{"Loading..."}</div>
  ) : (
    <GameManagerContext.Provider value={gameManagerContextValue}>
      <GameManager />
    </GameManagerContext.Provider>
  );
};

export default Game;
