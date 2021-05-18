import React from "react";
import { ReactQueryDevtools } from "react-query/devtools";

import "./App.css";
import { Container } from "semantic-ui-react";
import ItemsList from "./components/items-list";

function App() {
  return (
    <>
      <Container>
        <ItemsList />
      </Container>

      <ReactQueryDevtools initialIsOpen={false} />
    </>
  );
}

export default App;
