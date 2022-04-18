import React from "react";
import { Link, NavLink } from "react-router-dom";

const NavBar = () => {
  var width = "280px";
  return (
    <div className="nav_diary">
      <div class="container" style={{ width: width }}>
        <nav>
          <ul class="mcd-menu">
            <li>
              <Link to="/">
                <i class="fa fa-home"></i>
                <strong>iDiary</strong>
              </Link>
            </li>
            <li>
              <NavLink
                className="nav-link"
                aria-current="page"
                to="/idiary/myDiary"
              >
                <i class="fa fa-edit"></i>
                <strong>My diary</strong>
              </NavLink>
            </li>
            <li>
              <NavLink className="nav-link" to="/idiary/world">
                <i class="fa fa-gift"></i>
                <strong>World</strong>
              </NavLink>
            </li>
            <li>
              <NavLink to="/idiary/personalCenter" className="nav-link">
                <i class="fa fa-globe"></i>
                <strong>PersonalCenter</strong>
              </NavLink>
            </li>
          </ul>
        </nav>
      </div>
    </div>
  );
};

export default NavBar;
