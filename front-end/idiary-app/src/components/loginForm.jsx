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
    await userService.login(data);
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
      <div className="form-signin">
        <form onSubmit={this.handleSubmit}>
          <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
          <Input
            onChange={this.handleChange}
            value={data.username}
            type="text"
            id="username"
            errors={errors}
            label="Username"
          />
          <Input
            onChange={this.handleChange}
            value={data.password}
            type="password"
            id="password"
            errors={errors}
            label="Password"
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
              <span class="register">
                <Link to="/register">Register here</Link>
              </span>
            </span>{" "}
          </div>
        </form>
      </div>
    );
  }
}

export default LoginForm;
