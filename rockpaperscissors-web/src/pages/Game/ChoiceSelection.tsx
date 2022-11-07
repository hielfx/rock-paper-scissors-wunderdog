import React from "react";
import { Button, Col, Row } from "react-bootstrap";
import { useGameManagerContext } from "src/context/GameManager";
import { Move, Moves } from "src/types/game";
import { Commands } from "src/types/gameCommands";

/**
 * Choice seleccion for the players
 */
const ChoiceSelection = () => {
  const { ws, currentPlayer, game } = useGameManagerContext();

  const currentRound = game.rounds[game.currentRound];
  const selectedMove = currentRound?.playerSelections
    ? currentRound.playerSelections[currentPlayer]
    : "";

  const handleChoice = (move: Move) => {
    ws.send(
      JSON.stringify({
        gameId: game.id,
        nickname: currentPlayer,
        command: Commands.play,
        value: move,
      })
    );
  };

  return (
    <div>
      <Row>
        <Col>
          <Button
            disabled={!!selectedMove}
            variant={
              selectedMove === Moves.rock ? "primary" : "outline-primary"
            }
            style={{ width: "100%" }}
            onClick={() => handleChoice(Moves.rock)}
          >
            Rock
          </Button>
        </Col>
        <Col>
          <Button
            disabled={!!selectedMove}
            variant={
              selectedMove === Moves.paper ? "primary" : "outline-primary"
            }
            style={{ width: "100%" }}
            onClick={() => handleChoice(Moves.paper)}
          >
            Paper
          </Button>
        </Col>
        <Col>
          <Button
            disabled={!!selectedMove}
            variant={
              selectedMove === Moves.scissors ? "primary" : "outline-primary"
            }
            style={{ width: "100%" }}
            onClick={() => handleChoice(Moves.scissors)}
          >
            Scissors
          </Button>
        </Col>
      </Row>
    </div>
  );
};

export default ChoiceSelection;
