import { Route, Routes, Navigate } from "react-router-dom";
import MyDiary from "./components/myDiary";
import React from "react";
import World from "./components/world";
import About from "./components/about";
import NotFound from "./components/notFound";
import PersonalCenter from "./components/personalCenter";
import NavBar from "./components/navBar";
import "./App.css";

function App() {
  return (
    <React.Fragment>
      <NavBar />
      <Routes>
        <Route path="/myDiary/*" element={<MyDiary />} />
        <Route path="/world" element={<World />} />
        <Route path="/personalCenter" element={<PersonalCenter />} />
        <Route path="/about" element={<About />} />
        <Route path="/not-found" element={<NotFound />} />
        <Route path="*" element={<Navigate replace to="/not-found" />} />
        <Route
          path="/myDiary"
          exact
          element={<Navigate replace to="/myDiary/personalDiary" />}
        />
      </Routes>
    </React.Fragment>
  );
}

export default App;
