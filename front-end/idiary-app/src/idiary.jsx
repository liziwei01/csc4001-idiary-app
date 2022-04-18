import { Route, Routes, Navigate } from "react-router-dom";
import MyDiary from "./components/myDiary";
import React, { Component } from "react";
import World from "./components/world";
import NotFound from "./components/notFound";
import PersonalCenter from "./components/personalCenter";
import NavBar from "./components/navBar";
import "./App.css";
import auth from "./service/authService";
import RequireAuth from "./components/common/protectRoute";

class IDiary extends Component {
  render() {
    return (
      <div>
        <div className="Navbar">
          <NavBar />
        </div>
        <div className="mainpart">
          <Routes>
            <Route
              path=""
              exact
              element={
                <RequireAuth>
                  <Navigate replace to="myDiary" />
                </RequireAuth>
              }
            />
            <Route path="myDiary/*" element={<MyDiary />} />
            <Route path="world" element={<World />} />
            <Route path="personalCenter" element={<PersonalCenter />} />
            <Route path="/not-found" element={<NotFound />} />
            <Route path="*" element={<Navigate replace to="/not-found" />} />
          </Routes>
        </div>
      </div>
    );
  }
}

export default IDiary;
