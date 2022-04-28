import React from "react";
import { Navigate } from "react-router-dom";

function RequireAuth({ children }) {
  const authed = localStorage.getItem("token");

  return authed ? children : <Navigate to="/login" replace />;
}
export default RequireAuth;
