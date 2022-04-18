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
