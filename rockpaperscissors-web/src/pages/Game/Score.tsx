import React from "react";
import { Col, Row } from "react-bootstrap";
import { useGameManagerContext } from "src/context/GameManager";
import { Round } from "src/types/game";

/**
 * displays the players score and rounds
 */
export const Score = () => {
  const { game, currentPlayer } = useGameManagerContext();

  const opponentPlayer = Object.keys(game.players).find(
    (e) => e !== currentPlayer
  ) as string;

  /**
   * returns a color based on the player choices and the result
   */
  const getRoundChoiceColor = (round: Round, player: string) => {
    let color = "inherit";
    const bothPlayersChoose =
      round.playerSelections &&
      Object.keys(round.playerSelections).length === 2;
    if (round.winner === player && bothPlayersChoose) {
      color = "#81C784";
    } else if (round.isDraw) {
      color = "#64B5F6";
    } else if (round.winner !== player && bothPlayersChoose) {
      color = "#E57373";
    }

    return color;
  };

  return (
    <div>
      <Row>
        <Col style={{ textAlign: "center" }}>
          <div>{`${currentPlayer} - ${game.rounds.reduce(
            (prev, cur) => prev + (cur.winner === currentPlayer ? 1 : 0),
            0
          )}`}</div>
          {game.status === "FINISHED" && game.winner === currentPlayer && (
            <div>{"(YOU WON)"}</div>
          )}
        </Col>
        <Col style={{ textAlign: "center" }}>{`Round ${game.currentRound + 1}/${
          game.roundsNumber
        }`}</Col>
        <Col style={{ textAlign: "center" }}>
          <div>
            {`${game.rounds.reduce(
              (prev, cur) => prev + (cur.winner === opponentPlayer ? 1 : 0),
              0
            )} - ${opponentPlayer}`}
          </div>
          {game.status === "FINISHED" && game.winner === opponentPlayer && (
            <div>{"(WINNER)"}</div>
          )}
        </Col>
      </Row>
      <hr />
      {game.rounds.map(
        (round, idx) =>
          round.playerSelections &&
          Object.keys(round.playerSelections).length > 0 && (
            <Row>
              <Col
                style={{
                  textAlign: "center",
                  color: getRoundChoiceColor(round, currentPlayer),
                }}
              >
                {round.playerSelections[currentPlayer] ||
                  "Waiting for choice..."}
              </Col>
              <Col
                style={{
                  textAlign: "center",
                }}
              >{`${idx + 1}/${game.roundsNumber}`}</Col>
              <Col
                style={{
                  textAlign: "center",
                  color: getRoundChoiceColor(round, opponentPlayer),
                }}
              >
                {round.playerSelections[opponentPlayer]
                  ? round.isDraw || round.winner
                    ? round.playerSelections[opponentPlayer]
                    : "Opponent made a choice"
                  : "Waiting for choice..."}
              </Col>
            </Row>
          )
      )}
    </div>
  );
};

export default Score;
