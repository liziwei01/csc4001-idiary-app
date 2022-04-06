import React, { Component } from "react";
import InputWithButton from "./common/inputWithButton";
import * as userService from "../service/userService";
import InputWithDesc from "./common/inputWithDesc";
import Button from "./common/button";
import { Link } from "react-router-dom";

class ForgetPasswordByEmail extends Component {
  state = { data: { email: "", verification: "" }, errors: {} };
  handleSubmit = async (e) => {
    e.preventDefault();
    //server
    const data = { ...this.state.data };
    await userService.send_email(data);
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
      this.validateProperty({ id: "email", value: data.email }) ||
      this.validateProperty({
        id: "verification",
        value: data.verification,
      })
    )
      return true;
    else return false;
  };
  validateProperty = ({ id, value }) => {
    if (id === "email") {
      if (value.trim() === "") return "email is required!";
    }
    if (id === "verification") {
      if (value.trim() === "") return "verification is required!";
    }
  };
  handleClicked = async () => {
    console.log("send");
    const response = await userService.send_email(this.state.data);
  };
  render() {
    const { data, errors } = this.state;
    return (
      <div className="form-register">
        <h1 className="h3 mb-3 fw-normal">
          We need to verify your identity ...
        </h1>
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
          <Button disabled={this.validate()} label="Continue" />
          <hr class="my-4"></hr>
          <h2 class="fs-5 fw-bold mb-3">Or you can choose</h2>
          <button class="w-100 py-2 mb-2 btn btn-outline-secondary rounded-4">
            <Link className="links" to="/forgetPasswordByProblem">
              By Secret Protection Problem
            </Link>
          </button>
        </form>
      </div>
    );
  }
}

export default ForgetPasswordByEmail;
