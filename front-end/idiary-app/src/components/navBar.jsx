import React from "react";
import { Link, NavLink } from "react-router-dom";

const NavBar = () => {
  var width = "280px";
  return (
    <div class="container" style={{ width: width }}>
      <nav>
        <ul class="mcd-menu">
          <li>
            <Link to="/about">
              <i class="fa fa-home"></i>
              <strong>iDiary</strong>
            </Link>
          </li>
          <li>
            <NavLink className="nav-link" aria-current="page" to="/myDiary">
              <i class="fa fa-edit"></i>
              <strong>My diary</strong>
            </NavLink>
          </li>
          <li>
            <NavLink className="nav-link" to="/world">
              <i class="fa fa-gift"></i>
              <strong>World</strong>
            </NavLink>
          </li>
          <li>
            <NavLink to="/personalCenter" className="nav-link">
              <i class="fa fa-globe"></i>
              <strong>PersonalCenter</strong>
            </NavLink>
          </li>
          <li>
            <NavLink to="/login" className="nav-link">
              <i class="fa fa-comments-o"></i>
              <strong>Login</strong>
            </NavLink>
          </li>
          <li>
            <NavLink to="/register" className="nav-link">
              <i class="fa fa-picture-o"></i>
              <strong>Register</strong>
            </NavLink>
          </li>
        </ul>
      </nav>
    </div>
  );
};

export default NavBar;
