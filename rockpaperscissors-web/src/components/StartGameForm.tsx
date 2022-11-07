import React from "react";
import Card from "react-bootstrap/Card";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import { SubmitHandler, useForm } from "react-hook-form";
import { createGame, CreateGameBody } from "src/api/game";
import { useNavigate } from "react-router-dom";
import { apiErrorHandler } from "src/api/axios";

/**
 * Game type start used to display the creation or joining form
 */
type StartType = "create" | "join";

interface StartGameFormProps {
  startType: StartType;
}

type GameStartButtonTextType = {
  [key in StartType]: string;
};
const GAME_START_BUTTON_TEXT: GameStartButtonTextType = {
  create: "Create",
  join: "Join",
};

/**
 * Form values when submitting. It uses both create and join game values
 */
type FormValues = {
  nickname: string;
  rounds: number;
  gameId: string;
};

/**
 * StartGameForm component renders a form to start or join a game.
 */
const StartGameForm = (props: StartGameFormProps) => {
  const { startType } = props;

  const isCreate = startType === "create";
  const isJoin = startType === "join";

  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>({ defaultValues: { rounds: 1 } });

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    if (data.gameId) {
      console.log("joining game: ", data);
      navigate(`/game/${data.gameId}`, {
        state: {
          nickname: data.nickname,
        },
      });
    } else {
      console.log("creating game: ", data);
      const payload: CreateGameBody = {
        rounds: data.rounds,
      };
      createGame(payload, (game, error) => {
        if (error) {
          apiErrorHandler(error);
        }
        if (game) {
          navigate(`/game/${game.id}`, {
            state: {
              nickname: data.nickname,
              game,
            },
          });
        }
      });
    }
  };

  return (
    <Card className="p-3" data-testid="start-game-form-component">
      <h2 className="mb-3 text-center text-uppercase">{`${GAME_START_BUTTON_TEXT[startType]} Game`}</h2>
      <Form id={`${startType}-game-form`} onSubmit={handleSubmit(onSubmit)}>
        <Row>
          <Col>
            <Form.Group>
              <Form.Label>Nickname</Form.Label>
              <Form.Control
                {...register("nickname", { required: "Nickname is required" })}
                placeholder="Nickname"
                data-testid="start-game-form-nickname-input"
              />
              {errors?.nickname && (
                <p
                  className="text-danger"
                  role="alert"
                  data-testid="start-game-form-nickname-error"
                >
                  {errors.nickname.message}
                </p>
              )}
            </Form.Group>
          </Col>
          {isCreate && (
            <Col md="3">
              <Form.Group>
                <Form.Label>Rounds</Form.Label>
                <Form.Control
                  {...register("rounds", {
                    valueAsNumber: true,
                    required: "Round number is required",
                    min: { value: 1, message: "It should be at least 1 round" },
                  })}
                  type="number"
                  placeholder="Rounds"
                  min={1}
                />
                {errors?.rounds && (
                  <p className="text-danger" role="alert">
                    {errors.rounds.message}
                  </p>
                )}
              </Form.Group>
            </Col>
          )}
          {isJoin && (
            <Col>
              <Form.Group>
                <Form.Label>Game ID</Form.Label>
                <Form.Control
                  {...register("gameId", { required: "Game ID is required" })}
                  type="text"
                  placeholder="Game ID"
                />
                {errors?.gameId && (
                  <p className="text-danger">{errors.gameId.message}</p>
                )}
              </Form.Group>
            </Col>
          )}
        </Row>
        <div className="d-grid">
          <Button
            data-testid="start-game-form-submit-btn"
            type="submit"
            form={`${startType}-game-form`}
            size="lg"
            className="mt-3"
          >{`${GAME_START_BUTTON_TEXT[startType]} Game!`}</Button>
        </div>
      </Form>
    </Card>
  );
};

export default StartGameForm;
