import { Route, Routes, Navigate, Switch } from "react-router-dom";
import MyDiary from "./components/myDiary";
import React from "react";
import World from "./components/world";
import About from "./components/about";
import NotFound from "./components/notFound";
import PersonalCenter from "./components/personalCenter";
import NavBar from "./components/navBar";
import "./App.css";
import LoginForm from "./components/loginForm";
import RegisterForm from "./components/RegisterForm";
import PersonalDiary from "./components/personalDiary";
import FriendsDiary from "./components/friendsDiary";

function App() {
  return (
    <React.Fragment>
      <NavBar />
      <Routes>
        <Route path="/" exact element={<Navigate replace to="/about" />} />
        <Route path="/register" element={<RegisterForm />} />
        <Route path="/login" element={<LoginForm />} />
        <Route path="/myDiary/*" element={<MyDiary />} />
        <Route path="/world" element={<World />} />
        <Route path="/personalCenter" element={<PersonalCenter />} />
        <Route path="/about" element={<About />} />
        <Route path="/not-found" element={<NotFound />} />
        <Route path="*" element={<Navigate replace to="/not-found" />} />
      </Routes>
    </React.Fragment>
  );
}

export default App;
