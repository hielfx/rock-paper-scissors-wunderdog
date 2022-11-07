import React from "react";
import { act, fireEvent, render, screen } from "@testing-library/react";
import StartGameForm from "../StartGameForm";
import { BrowserRouter } from "react-router-dom";
import userEvent from "@testing-library/user-event";
import mockAxios from "jest-mock-axios";

const COMPONENT_TEST_ID = "start-game-form-component";
const SUBMIT_BTN_TEST_ID = "start-game-form-submit-btn";
const NICKNAME_ERROR = "start-game-form-nickname-error";
const NICKNAME_INPUT = "start-game-form-nickname-input";

const clickSubmitButton = async () => {
  await act(async () => {
    const submitBtn = screen.getByTestId(SUBMIT_BTN_TEST_ID);
    userEvent.click(submitBtn);
  });
};

/**
 * Shared test case between both rendering options
 */
const testWithDefaultData = async () => {
  await clickSubmitButton();

  const nicknameErr = screen.getByTestId(NICKNAME_ERROR);
  expect(nicknameErr).toBeInTheDocument();
  expect(nicknameErr).toHaveTextContent("Nickname is required");
};

describe(`<${StartGameForm.name} startType="create" />`, () => {
  beforeEach(() => {
    render(<StartGameForm startType="create" />, { wrapper: BrowserRouter });
  });
  afterEach(() => {
    mockAxios.reset();
  });

  it("renders create type without errors", () => {
    expect(screen.getByTestId(COMPONENT_TEST_ID)).toBeInTheDocument();
  });

  it("submit with default data", testWithDefaultData);

  it("create game successfully", async () => {
    const axiosMockedResponse = {
      data: {
        gameId: "testing-game-id",
      },
    };

    mockAxios.post.mockResolvedValueOnce(axiosMockedResponse);

    const nicknameInput = screen.getByTestId(NICKNAME_INPUT);
    expect(nicknameInput).toBeInTheDocument();
    await act(async () => {
      fireEvent.input(nicknameInput, {
        target: {
          value: "testing-player-id",
        },
      });
    });
    expect(nicknameInput).toHaveValue("testing-player-id");
    await clickSubmitButton();

    expect(mockAxios.post).toHaveBeenCalled();
  });
});

describe(`<${StartGameForm.name} startType="join" />`, () => {
  beforeEach(() => {
    render(<StartGameForm startType="join" />, { wrapper: BrowserRouter });
  });
  afterEach(() => {
    mockAxios.reset();
  });

  it("renders create type without errors", () => {
    expect(screen.getByTestId(COMPONENT_TEST_ID)).toBeInTheDocument();
  });

  it("submit with default data", testWithDefaultData);
});
