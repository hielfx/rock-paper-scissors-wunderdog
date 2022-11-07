import React from "react";
import { Container } from "react-bootstrap";
import { Route, Routes } from "react-router-dom";
import routes from "../routes";
import Title from "./Title";

const DefaultLayout = () => (
  <div id="app">
    <Container>
      <Title />
      <Routes>
        {routes.map((route, idx) => {
          return (
            route.element && (
              <Route key={idx} path={route.path} element={route.element} />
            )
          );
        })}
      </Routes>
    </Container>
  </div>
);

export default DefaultLayout;
