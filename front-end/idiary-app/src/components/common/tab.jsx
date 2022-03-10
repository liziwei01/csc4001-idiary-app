import React from "react";
import { NavLink } from "react-router-dom";

const Tab = () => {
  return (
    <ul className="nav nav-tabs justify-content-center">
      <li className="nav-item">
        <NavLink
          className="nav-link"
          aria-current="page"
          to="/myDiary/personalDiary"
        >
          My Diary
        </NavLink>
      </li>
      <li className="nav-item">
        <NavLink className="nav-link" to="/myDiary/friendsDiary">
          Friends
        </NavLink>
      </li>
    </ul>
  );
};

export default Tab;
