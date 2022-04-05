import React, { Component } from "react";
import * as userService from "./../service/userService";
import InputWithButton from "./common/inputWithButton";
import InputWithDesc from "./common/inputWithDesc";
import Button from "./common/button";
import Input from "./common/input";

class RegisterForm extends Component {
  state = {
    data: {
      email: "",
      verification: "",
      username: "",
      password: "",
      repeatpassword: "",
      answer1: "",
      answer2: "",
      answer3: "",
      agreement: false,
      secretproblem: false,
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
        id: "answer1",
        value: data.answer1,
      }) ||
      this.validateProperty({
        id: "answer2",
        value: data.answer2,
      }) ||
      this.validateProperty({
        id: "answer3",
        value: data.answer3,
      }) ||
      data.agreement === false
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
    if (id === "answer1" && this.state.data.secretproblem === true) {
      if (value.trim() === "") return "Answer is required!";
    }
    if (id === "answer2" && this.state.data.secretproblem === true) {
      if (value.trim() === "") return "Answer is required!";
    }
    if (id === "answer3" && this.state.data.secretproblem === true) {
      if (value.trim() === "") return "Answer is required!";
    }
  };
  handlechecked = (e) => {
    const checked = e.target.checked;
    const data = { ...this.state.data };
    if (checked) {
      data[e.target.id] = true;
    } else {
      data[e.target.id] = false;
    }
    this.setState({ data });
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
  handleClicked = async () => {
    console.log("send");
    const response = await userService.send_email(this.state.data);
  };

  render() {
    const { data, errors } = this.state;
    return (
      <div className="form">
        <h1 className="mb-3">Register</h1>
        <form onSubmit={this.handleSubmit}>
          <InputWithButton
            onChange={this.handleChange}
            value={data.email}
            type="email"
            id="email"
            errors={errors}
            text="Send"
            label="Email Address"
            onClick={this.handleClicked}
          />
          <InputWithDesc
            onChange={this.handleChange}
            value={data.verification}
            type="text"
            id="verification"
            errors={errors}
            text=""
            label="Verification Code"
          />

          <InputWithDesc
            onChange={this.handleChange}
            value={data.username}
            type="text"
            id="username"
            errors={errors}
            text="Must be 6-20 characters long."
            label="Username"
          />
          <InputWithDesc
            onChange={this.handleChange}
            value={data.password}
            type="password"
            id="password"
            errors={errors}
            text="Must be 8-20 characters long."
            label="Password"
          />
          <InputWithDesc
            onChange={this.handleChange}
            value={data.repeatpassword}
            type="repeatpassword"
            id="repeatpassword"
            errors={errors}
            text="Please repeat your password."
            label="Repeat Password"
          />

          <div className="form-check form-switch mb-3">
            <input
              className="form-check-input"
              type="checkbox"
              role="switch"
              id="secretproblem"
              onClick={this.handlechecked}
            />
            <label className="form-check-label" htmlFor="secretproblem">
              Answer Secret Protect Problems
            </label>
          </div>

          {data.secretproblem && (
            <div>
              <Input
                onChange={this.handleChange}
                value={data.answer1}
                type="text"
                id="answer1"
                errors={errors}
                label="What's your birth day?"
              />
              <Input
                onChange={this.handleChange}
                value={data.answer2}
                type="text"
                id="answer2"
                errors={errors}
                label="What's your mother's name?"
              />
              <Input
                onChange={this.handleChange}
                value={data.answer3}
                type="text"
                id="answer3"
                errors={errors}
                label="What's your favorite color?"
              />
            </div>
          )}

          <div className="mb-3 form-check">
            <input
              type="checkbox"
              className="form-check-input"
              id="agreement"
              onClick={this.handlechecked}
            />
            <label
              className="form-check-label"
              htmlFor="agreement"
              data-bs-toggle="modal"
              data-bs-target="#staticBackdrop"
            >
              Agree with the agreement.
            </label>
          </div>

          <Button disabled={this.validate()} label="Register" />
        </form>
      </div>
    );
  }
}

export default RegisterForm;
