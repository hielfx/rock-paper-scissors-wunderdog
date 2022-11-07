import React from "react";
import CreateGame from "src/pages/CreateGame";
import Home from "src/pages/Home";
import JoinGame from "src/pages/JoinGame";
import Game from "src/pages/Game";
import { Navigate, RouteProps } from "react-router-dom";

// const Home = React.lazy(() => import("src/pages/Home"));
// const CreateGame = React.lazy(() => import("src/pages/CreateGame"));
// const JoinGame = React.lazy(() => import("src/pages/JoinGame"));

const routes: Array<RouteProps> = [
  { path: "*", element: <Navigate to="/" /> },
  { path: "/", element: <Home /> },
  { path: "/game/:gameId", element: <Game /> },
  { path: "/create-game", element: <CreateGame /> },
  { path: "/join-game/:gameId", element: <JoinGame /> },
];

export default routes;
