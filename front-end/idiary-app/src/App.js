import { Route, Routes, Navigate } from "react-router-dom";
import React, { Component } from "react";
import NotFound from "./components/notFound";
import "./App.css";
import LoginForm from "./components/loginForm";
import RegisterForm from "./components/RegisterForm";
import ForgetPasswordByProblem from "./components/forgetPasswordByProblem";
import ForgetPasswordByEmail from "./components/forgetPasswordByEmail";
import ResetPassword from "./components/resetPassword";
import Landing from "./landing";
import IDiary from "./idiary";
import Admin from "./components/admin";


class App extends Component {
  render() {
    return (
      <div>
        <Routes>
          <Route path="/" element={<Landing />} />
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
}

export default App;
