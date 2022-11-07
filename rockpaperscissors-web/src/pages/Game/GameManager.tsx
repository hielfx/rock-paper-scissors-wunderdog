import React, { useEffect } from "react";
import Score from "./Score";
import ChoiceSelection from "./ChoiceSelection";
import { useGameManagerContext } from "src/context/GameManager";
import { useNavigate } from "react-router-dom";
import Swal, { SweetAlertIcon } from "sweetalert2";
import { Button } from "react-bootstrap";
import { Commands } from "src/types/gameCommands";

/**
 * Manages the game based on the status
 */
const GameManager = () => {
  const { game, currentPlayer, ws } = useGameManagerContext();
  const navigate = useNavigate();

  useEffect(() => {
    let title: string, icon: SweetAlertIcon;
    switch (game.status) {
      case "FINISHED":
        // If one player is missing, don't show the status again
        if (Object.keys(game.players).length === 2) {
          if (game.isDraw) {
            title = "It's a draw!";
            icon = "info";
          } else if (game.winner === currentPlayer) {
            title = "You won the game!";
            icon = "success";
          } else if (game.winner && game.winner !== currentPlayer) {
            title = "You lost the game!";
            icon = "error";
          } else if (!game.isDraw && !game.winner) {
            title = "A player left the game!";
            icon = "warning";
          } else {
            title = "Game finished abnormaly";
            icon = "warning";
          }
          Swal.fire({
            title,
            icon,
            allowEnterKey: true,
            allowOutsideClick: false,
            allowEscapeKey: false,
          });
        }
        break;
      case "STARTED":
        Swal.close();
        if (game.currentRound > 0 && game.roundsNumber > 1) {
          const currentRound = game.rounds[game.currentRound];
          //Display the previous round results if there's no move in the current one
          if (
            !currentRound?.isDraw &&
            !currentRound?.winner &&
            currentRound?.playerSelections === null
          ) {
            const round = game.rounds[game.currentRound - 1];
            if (round.isDraw) {
              title = "It's a draw!";
              icon = "info";
            } else if (round.winner === currentPlayer) {
              title = "You won the round!";
              icon = "success";
            } else {
              title = "You lost the round!";
              icon = "error";
            }
            Swal.fire({
              title,
              icon,
              allowEnterKey: true,
              allowOutsideClick: false,
              allowEscapeKey: false,
            });
          }
        }
        break;
      case "PAUSED":
        if (game.pausedBy === currentPlayer) {
          Swal.fire({
            title: "You paused the game",
            text: "Your opponent is waiting for the game to resume",
            confirmButtonText: "Resume game",
            // icon: "warning",
            allowEnterKey: false,
            allowOutsideClick: false,
            allowEscapeKey: false,
          }).then(() => {
            handleUnpause();
          });
        } else {
          Swal.fire({
            title: "Opponent paused the game",
            text: "Waiting for game to resume...",
            allowOutsideClick: false,
            allowEnterKey: false,
            allowEscapeKey: false,
            didOpen: () => {
              Swal.showLoading();
            },
          });
        }
        break;
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [game]);

  const handlePause = () => {
    ws.send(
      JSON.stringify({
        gameId: game.id,
        nickname: currentPlayer,
        command: Commands.pause,
      })
    );
  };

  const handleUnpause = () => {
    ws.send(
      JSON.stringify({
        gameId: game.id,
        nickname: currentPlayer,
        command: Commands.unpause,
      })
    );
  };

  return (
    <div>
      {/* <p>{JSON.stringify(game)}</p> */}
      {game.id && Object.keys(game.players).length < 2 && (
        <p>
          Send this ID to your friend! <b>{game.id}</b>
        </p>
      )}
      {Object.keys(game.players).length < 2 && game.status !== "FINISHED" ? (
        <p>Waiting for other player to join....</p>
      ) : (
        <>
          <Score />
          <br />
          {game.status === "STARTED" && <ChoiceSelection />}
          <br />
          {game.status !== "FINISHED" && (
            <Button onClick={handlePause}>Pause game</Button>
          )}
        </>
      )}

      {game.status === "FINISHED" && (
        <Button onClick={() => navigate("/")}>New game</Button>
      )}
    </div>
  );
};

export default GameManager;
