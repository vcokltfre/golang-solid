import { Route, Routes } from "@solidjs/router";
import { Component } from "solid-js";

const App: Component = () => {
  return (
    <Routes>
      <Route path="/" element={<h1>Home</h1>} />
    </Routes>
  );
};

export default App;
