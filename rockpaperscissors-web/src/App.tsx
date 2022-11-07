import React, { Suspense } from "react";
import { BrowserRouter } from "react-router-dom";
import DefaultLayout from "src/components/DefaultLayout";

function App() {
  return (
    <BrowserRouter>
      <Suspense fallback="Loading...">
        <DefaultLayout />
      </Suspense>
    </BrowserRouter>
  );
}

export default App;
