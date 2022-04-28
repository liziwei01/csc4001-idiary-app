import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter } from "react-router-dom";
import reportWebVitals from "./reportWebVitals";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap/dist/js/bootstrap.bundle";
import "font-awesome/css/font-awesome.css";
import "bootstrap-icons/font/bootstrap-icons.css";
import "./index.css";
import "./css/profile.css";
import "./css/style.css";
import "./css/nav.css";
import "./css/aos.css";
import "./css/glightbox.min.css";
import "./css/swiper-bundle.min.css";
import "./css/style.min.css";
import "./css/materialdesignicons.min.css";
import "./css/bootstrap.min.css";
import "./css/jstable.css";
import App from "./App";

ReactDOM.render(
  <BrowserRouter>
    <App />
  </BrowserRouter>,
  document.getElementById("root")
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
