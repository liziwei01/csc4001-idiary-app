import React, { Component } from "react";
import Input from "./common/input";
import Button from "./common/button";
import { Link } from "react-router-dom";
import * as userService from "./../service/userService";

class LoginForm extends Component {
  state = { data: { username: "", password: "" }, errors: {} };
  handleSubmit = async (e) => {
    e.preventDefault();
    //server
    const data = { ...this.state.data };
    const response = await userService.login(data);
    console.log(response.data);
    const errors = { ...this.state.errors };
    if (response.data.errmsg != "Success") errors.username = response.data.data;
    this.setState({ errors });
  };
  validate = () => {
    const { data } = this.state;
    if (
      this.validateProperty({ id: "username", value: data.username }) ||
      this.validateProperty({
        id: "password",
        value: data.password,
      })
    )
      return true;
    else return false;
  };
  validateProperty = ({ id, value }) => {
    if (id === "username") {
      if (value.trim() === "") return "Username is required!";
    }
    if (id === "password") {
      if (value.trim() === "") return "Password is required!";
    }
  };
  handleChange = ({ target: input }) => {
    const errors = { ...this.state.errors };
    const errorMessage = this.validateProperty(input);
    if (errorMessage) errors[input.id] = errorMessage;
    else delete errors[input.id];

    const data = { ...this.state.data };
    data[input.id] = input.value;
    this.setState({ data, errors });
  };
  render() {
    const { data, errors } = this.state;
    return (
      <section class="sign-in sign">
        <div class="container">
          <div class="signin-content">
            <div class="signin-image">
              <figure>
                <img
                  src={require("../img/signin-image.jpg")}
                  alt="sing up image"
                />
              </figure>
            </div>
            <div className="signin-form">
              <h2 class="form-title">Sign up</h2>
              <form
                onSubmit={this.handleSubmit}
                method="POST"
                class="register-form"
                id="login-form"
              >
                <Input
                  onChange={this.handleChange}
                  value={data.username}
                  type="text"
                  id="username"
                  errors={errors.username}
                  label="fa fa-user"
                  placeholder="username"
                />
                <Input
                  onChange={this.handleChange}
                  value={data.password}
                  type="password"
                  id="password"
                  errors={errors.password}
                  label="fa fa-unlock-alt"
                  placeholder="password"
                />
                <Button disabled={this.validate()} label="Sign in" />
                <div className="row justify-content-center my-2">
                  <Link className="links" to="/forgetPasswordByEmail">
                    <small className="text-muted">Forget Password?</small>
                  </Link>
                </div>

                <hr class="my-4"></hr>
                <div class="text2 mt-4 d-flex flex-row align-items-center">
                  {" "}
                  <span>
                    Don't have an account?
                    <span class="register ">
                      <Link to="/register">Register here</Link>
                    </span>
                  </span>{" "}
                </div>
              </form>
            </div>
          </div>
        </div>
      </section>
    );
  }
}

export default LoginForm;
