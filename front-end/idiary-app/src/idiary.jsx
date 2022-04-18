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
import ForgetPasswordByProblem from "./components/forgetPasswordByProblem";
import ForgetPasswordByEmail from "./components/forgetPasswordByEmail";
import ResetPassword from "./components/resetPassword";

const IDiary = () => {
  return (
    <div>
      <div className="Navbar">
        <NavBar />
      </div>
      <div className="mainpart">
        <Routes>
          <Route path="" exact element={<Navigate replace to="myDiary" />} />
          <Route path="myDiary/*" element={<MyDiary />} />
          <Route path="world" element={<World />} />
          <Route path="personalCenter" element={<PersonalCenter />} />
          <Route path="/not-found" element={<NotFound />} />
          <Route path="*" element={<Navigate replace to="/not-found" />} />
        </Routes>
      </div>
    </div>
  );
};

export default IDiary;
