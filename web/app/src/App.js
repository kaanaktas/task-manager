import React from "react";
import "./App.css";
import { Container } from "semantic-ui-react";
import TaskList from "./TaskList";

function App() {
  return (
      <div>
        <Container>
          <TaskList />
        </Container>
      </div>
  );
}
export default App;