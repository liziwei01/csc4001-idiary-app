import React, { Component } from "react";

class LoginForm extends Component {
  state = { data: { username: "", password: "" }, errors: {} };
  handleSubmit = (e) => {
    e.preventDefault();
    //server
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
  handleChange = ({ currentTarget: input }) => {
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
          <div className="mb-3">
            <label htmlFor="username" className="form-label">
              Username
            </label>
            <input
              onChange={this.handleChange}
              value={data.username}
              type="username"
              className="form-control"
              id="username"
            />
            {errors.username && (
              <div className="alert alert-danger">{errors.username}</div>
            )}
          </div>
          <div className="mb-3">
            <label htmlFor="password" className="form-label">
              Password
            </label>
            <input
              onChange={this.handleChange}
              value={data.password}
              type="password"
              className="form-control"
              id="password"
            />
            {errors.password && (
              <div className="alert alert-danger">{errors.password}</div>
            )}
          </div>
          <button
            disabled={this.validate()}
            type="submit"
            className="btn btn-primary"
          >
            Login in
          </button>
        </form>
      </div>
    );
  }
}

export default LoginForm;
