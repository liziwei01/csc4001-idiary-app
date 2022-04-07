import React, { Component } from "react";
import InputWithDesc from "./common/inputWithDesc";
import Button from "./common/button";
import * as userService from "./../service/userService";

class ResetPassword extends Component {
  state = { data: { password: "", repeatpassword: "" }, errors: {} };
  handleSubmit = async (e) => {
    e.preventDefault();
    //server
    const data = { ...this.state.data };
    await userService.login(data);
  };
  validate = () => {
    const { data } = this.state;
    if (
      this.validateProperty({
        id: "password",
        value: data.password,
      }) ||
      this.validateProperty({
        id: "repeatpassword",
        value: data.repeatpassword,
      })
    )
      return true;
    else return false;
  };
  validateProperty = ({ id, value }) => {
    if (id === "password") {
      if (value.trim() === "") return "Password is required!";
      if (value.length < 8 || value.length > 20)
        return "Your password should between 8 and 20 characters!";
    }
    if (id === "repeatpassword") {
      if (value !== this.state.data.password) return "Different!";
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
      <section class="signup">
        <div class="container">
          <div class="signup-content">
            <div className="form-register">
              <h1 className="h3 mb-3 fw-normal">
                Please enter your new password
              </h1>
              <form onSubmit={this.handleSubmit}>
                <InputWithDesc
                  onChange={this.handleChange}
                  value={data.password}
                  type="password"
                  id="password"
                  errors={errors}
                  text="Must be 8-20 characters long."
                  label="fa fa-unlock-alt"
                  placeholder="password"
                />
                <InputWithDesc
                  onChange={this.handleChange}
                  value={data.repeatpassword}
                  type="repeatpassword"
                  id="repeatpassword"
                  errors={errors}
                  text="Please repeat your password."
                  label="fa fa-repeat"
                  placeholder="repeat password"
                />
                <Button disabled={this.validate()} label="Reset Password" />
              </form>
            </div>
          </div>
        </div>
      </section>
    );
  }
}

export default ResetPassword;
