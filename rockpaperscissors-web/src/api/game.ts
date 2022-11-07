import axios from "./axios";
import { Game } from "src/types/game";
import { CallbackFunction } from "src/types/api";
import { AxiosError } from "axios";

export interface CreateGameBody {
  rounds: number;
}
/**
 * Calls the API to create a game based on the given data
 */
export const createGame = (
  data: CreateGameBody,
  cb: CallbackFunction<Game> = () => {}
) => {
  console.log("createGame DATA: ", data);
  axios
    .post<Game>("/games", data)
    .then((response) => {
      console.log("Response data: ", response.data);
      cb(response.data, null);
    })
    .catch((error: Error | AxiosError) => {
      console.log("Error: ", error);
      cb(null, error);
    });
};

/**
 * Calls the API to join a player to a game
 */
export const joinGame = (
  gameId: string,
  nickname: string,
  cb: CallbackFunction<Game> = () => {}
) => {
  axios
    .post(`/games/${gameId}`, { nickname })
    .then((response) => {
      console.log("Response data: ", response.data);
      cb(response.data, null);
    })
    .catch((error: Error | AxiosError) => {
      console.log("Error: ", error);
      cb(null, error);
    });
};
