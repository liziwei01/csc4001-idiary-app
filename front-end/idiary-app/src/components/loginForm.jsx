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
      <div className="form">
        <h1>Login</h1>
        <form onSubmit={this.handleSubmit}>
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
          <Button disabled={this.validate()} label="Log in" />
          <Link className="ms-4 mt-4" to="/register">
            Register
          </Link>
          <Link className="ms-4 mt-4" to="/forgetPasswordByEmail">
            Forget Password
          </Link>
        </form>
      </div>
    );
  }
}

export default LoginForm;
