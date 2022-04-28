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
import auth from "./service/authService";
import RequireAuth from "./components/common/protectRoute";

class App extends Component {
  state = { user: null };
  componentDidMount() {
    const user = auth.getCurrentUser();
    this.setState({ user });
  }
  render() {
    const user = { ...this.state.user };
    return (
      <div>
        <Routes>
          <Route path="/" element={<Landing />} />
          <Route
            path="/admin"
            element={
              <RequireAuth>
                <Admin />
              </RequireAuth>
            }
          />
          <Route
            path="/forgetPasswordByEmail"
            element={<ForgetPasswordByEmail />}
          />
          <Route
            path="/forgetPasswordByProblem"
            element={<ForgetPasswordByProblem />}
          />
          <Route path="/resetPassword" element={<ResetPassword />} />
          <Route path="/register" element={<RegisterForm />} />
          <Route path="/login" element={<LoginForm />} />
          <Route
            path="/idiary/*"
            element={
              <RequireAuth>
                <IDiary />
              </RequireAuth>
            }
          />
          <Route path="/not-found" element={<NotFound />} />

          <Route path="*" element={<Navigate replace to="/not-found" />} />
        </Routes>
      </div>
    );
  }
}

export default App;
