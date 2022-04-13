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
import PersonalDiary from "./components/personalDiary";
import FriendsDiary from "./components/friendsDiary";
import ResetPassword from "./components/resetPassword";
import Landing from "./landing";
import IDiary from "./idiary";
import Admin from "./components/admin";
import AAA from "./components/aaa";

function App() {
  return (
    <div>
      <Routes>
        <Route path="/" element={<Landing />} />
        <Route path="/aaa" element={<AAA />} />
        <Route path="/admin" element={<Admin />} />
        <Route
          path="/forgetPasswordByEmail"
          element={<ForgetPasswordByEmail />}
        />
        <Route
          path="/forgetPasswordByProblem"
          element={<ForgetPasswordByProblem />}
        />
        <Route path="/resetPassword" element={<ResetPassword />} />
        <Route path="/idiary/*" element={<IDiary />} />
        <Route path="/register" element={<RegisterForm />} />
        <Route path="/login" element={<LoginForm />} />

        <Route path="/not-found" element={<NotFound />} />
        <Route path="*" element={<Navigate replace to="/not-found" />} />
      </Routes>
    </div>
  );
}

export default App;
