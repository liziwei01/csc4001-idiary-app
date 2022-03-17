import React, { Component } from "react";
import { Link } from "react-router-dom";
import InputWithDesc from "./common/inputWithDesc";
import Input from "./common/input";
import Button from "./common/button";
import * as userService from "./../service/userService";

class ForgetPasswordByProblem extends Component {
  state = {
    data: { username: "", answer1: "", answer2: "", answer3: "" },
    errors: {},
  };
  handleSubmit = async (e) => {
    e.preventDefault();
    //server
    const data = { ...this.state.data };
    await userService.send_problem(data);
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
  validate = () => {
    const { data } = this.state;
    if (
      this.validateProperty({ id: "username", value: data.username }) ||
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
      })
    )
      return true;
    else return false;
  };
  validateProperty = ({ id, value }) => {
    if (id === "username") {
      if (value.trim() === "") return "username is required!";
    }
    if (id === "answer1") {
      if (value.trim() === "") return "answer1 is required!";
    }
    if (id === "answer2") {
      if (value.trim() === "") return "answer2 is required!";
    }
    if (id === "answer3") {
      if (value.trim() === "") return "answer3 is required!";
    }
  };

  render() {
    const { data, errors } = this.state;
    return (
      <div className="form">
        <h1 className="mb-3">We need to verify your identity</h1>
        <form onSubmit={this.handleSubmit}>
          <InputWithDesc
            onChange={this.handleChange}
            value={data.username}
            type="text"
            id="username"
            errors={errors}
            text="Pleaser enter your user name"
            label="Username"
          />
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
          <Button disabled={this.validate()} label="Continue" />
          <Link className="ms-4 mt-4" to="/forgetPasswordByEmail">
            By email
          </Link>
        </form>
      </div>
    );
  }
}

export default ForgetPasswordByProblem;
