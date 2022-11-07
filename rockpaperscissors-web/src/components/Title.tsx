import React from "react";

const Title = () => (
  <h1 className="my-3 text-center">
    <span className="text-red">Rock</span>,{" "}
    <span className="text-blue">Paper</span>,{" "}
    <span className="text-green">Scissors</span>!
  </h1>
);
Title.displayName = "Title";

export default Title;
