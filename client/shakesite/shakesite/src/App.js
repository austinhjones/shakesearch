import React from "react";
import "./App.css";
// import the Container Component from the semantic-ui-react
import { Container } from "react-bootstrap";
// import the ToDoList component
import ShakeSite from "./shakesite.js";
function App() {
  return (
      <Container fluid>
        <ShakeSite />
      </Container>
  );
}
export default App;