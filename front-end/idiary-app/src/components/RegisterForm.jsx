import React, { Component } from "react";
import * as userService from "./../service/userService";

class RegisterForm extends Component {
  state = {
    data: {
      email: "",
      username: "",
      password: "",
      repeatpassword: "",
      answer: "",
      problem: "Disabled",
      checked: false,
    },
    errors: {},
  };
  handleSubmit = async (e) => {
    e.preventDefault();
    //server
    const data = { ...this.state.data };
    delete data.repeatpassword;
    await userService.register(data);
  };
  handlechecked = (e) => {
    const checked = e.target.checked;
    const data = { ...this.state.data };
    if (checked) {
      data.checked = true;
    } else {
      data.checked = false;
    }
    this.setState({ data });
  };
  validate = () => {
    const { data } = this.state;
    if (
      this.validateProperty({ id: "email", value: data.email }) ||
      this.validateProperty({
        id: "password",
        value: data.password,
      }) ||
      this.validateProperty({
        id: "repeatpassword",
        value: data.repeatpassword,
      }) ||
      this.validateProperty({
        id: "username",
        value: data.username,
      }) ||
      this.validateProperty({
        id: "answer",
        value: data.answer,
      }) ||
      data.checked === false
    )
      return true;
    else return false;
  };
  validateProperty = ({ id, value }) => {
    if (id === "email") {
      if (value.trim() === "") return "Email is required!";
      var reg =
        /^\w+((-\w+)|(\.\w+))*\@[A-Za-z0-9]+((\.|-)[A-Za-z0-9]+)*\.[A-Za-z0-9]+$/;
      if (!reg.test(value)) return "Please enter correct email address!";
    }
    if (id === "password") {
      if (value.trim() === "") return "Password is required!";
      if (value.length < 8 || value.length > 20)
        return "Your password should between 8 and 20 characters!";
    }
    if (id === "repeatpassword") {
      if (value !== this.state.data.password) return "Different!";
    }
    if (id === "username") {
      if (value.trim() === "") return "Username is required!";
      if (value.length < 6 || value.length > 20)
        return "Your username should between 6 and 20 characters!";
    }
    if (id === "answer") {
      if (value.trim() === "" && this.state.data.problem !== "Disabled")
        return "Answer is required!";
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
        <h1 className="mb-3">Register</h1>
        <form onSubmit={this.handleSubmit}>
          <div className="mb-3 row">
            <div className="form-floating col">
              <input
                onChange={this.handleChange}
                value={data.email}
                type="email"
                className={
                  errors.email ? "form-control is-invalid" : "form-control"
                }
                id="email"
                placeholder="name@example.com"
              />
              <label htmlFor="email">Email address</label>
            </div>
            <div className="col">
              <button type="submit" class="btn btn-primary mb-3">
                Confirm identity
              </button>
            </div>
          </div>

          <div className="mb-3 row">
            <div className="form-floating col">
              <input
                onChange={this.handleChange}
                value={data.username}
                type="text"
                className={
                  errors.username ? "form-control is-invalid" : "form-control"
                }
                id="username"
                aria-describedby="usernameHelpInline"
                placeholder="Username"
              />
              <label htmlFor="username">Username</label>
            </div>
            <div class="col">
              <span id="usernameHelpInline" class="form-text">
                Must be 6-20 characters long.
              </span>
            </div>
          </div>

          <div className="mb-3 row">
            <div class="form-floating col">
              <input
                onChange={this.handleChange}
                value={data.password}
                type="password"
                className={
                  errors.password ? "form-control is-invalid" : "form-control"
                }
                id="password"
                aria-describedby="passwordHelpInline"
                placeholder="Password"
              />
              <label htmlfor="password">Password</label>
            </div>
            <div class="col">
              <span id="passwordHelpInline" class="form-text">
                Must be 8-20 characters long.
              </span>
            </div>
          </div>

          <div className="mb-3 row">
            <div className="form-floating col">
              <input
                onChange={this.handleChange}
                value={data.repeatpassword}
                type="password"
                className={
                  errors.repeatpassword
                    ? "form-control is-invalid"
                    : "form-control"
                }
                id="repeatpassword"
                placeholder="Repeat Password"
                aria-describedby="repeatPasswordHelpInline"
              />
              <label htmlFor="repeatpassword">Repeat Your Password</label>
            </div>
            <div class="col">
              <span id="repeatPasswordHelpInline" class="form-text">
                Please repeat your password.
              </span>
            </div>
          </div>

          <div className="row g-2 mb-3">
            <div className="col-md">
              <div className="form-floating">
                <input
                  type="text"
                  className={
                    errors.answer ? "form-control is-invalid" : "form-control"
                  }
                  id="answer"
                  placeholder="..."
                  value={data.problem === "Disabled" ? "" : data.answer}
                  onChange={this.handleChange}
                  disabled={data.problem === "Disabled" ? true : false}
                />
                <label htmlfor="answer">Your answer</label>
              </div>
            </div>
            <div className="col-md">
              <div className="form-floating">
                <select
                  className="form-select"
                  id="problem"
                  onChange={this.handleChange}
                >
                  <option selected>Disabled</option>
                  <option value="1">What's your mother's name?</option>
                  <option value="2">What's your birth day?</option>
                  <option value="3">What's your favourite color?</option>
                </select>
                <label htmlfor="problem">Optional secret problem</label>
              </div>
            </div>
          </div>

          <div className="mb-3 form-check">
            <input
              type="checkbox"
              className="form-check-input"
              id="check"
              onClick={(e) => {
                this.handlechecked(e);
              }}
            />
            <label
              className="form-check-label"
              htmlFor="check"
              data-bs-toggle="modal"
              data-bs-target="#exampleModal"
            >
              <a href="./common/modal"> Agree with the agreement.</a>
            </label>
          </div>
          <button
            disabled={this.validate()}
            type="submit"
            className="btn btn-primary"
          >
            Register
          </button>
        </form>
      </div>
    );
  }
}

export default RegisterForm;
