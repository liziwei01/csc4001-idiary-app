import React from "react";
import { Route, Routes } from "react-router-dom";
import PersonalDiary from "./personalDiary";
import FriendsDiary from "./friendsDiary";
import Tab from "./common/tab";

const MyDiary = () => {
  return (
    <React.Fragment>
      <div className="content">
        <p>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Dolorum nisi
          vitae nesciunt ipsa maiores eum labore consectetur vel quam,
          voluptates, sunt voluptatibus optio placeat nulla ullam debitis
          blanditiis inventore eius officia possimus ipsum ea. Repellendus
          molestiae enim iste a optio error, voluptate, facere dolorum adipisci
          porro rerum corrupti, architecto labore.
        </p>
      </div>
      <Tab />
      <Routes>
        <Route path="personalDiary" element={<PersonalDiary />} />
        <Route path="friendsDiary" element={<FriendsDiary />} />
      </Routes>
    </React.Fragment>
  );
};

export default MyDiary;
