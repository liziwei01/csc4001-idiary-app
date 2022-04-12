import React, { Component } from "react";
import * as userService from "./../service/userService";
import InputWithButton from "./common/inputWithButton";
import InputWithDesc from "./common/inputWithDesc";
import Button from "./common/button";
import Input from "./common/input";
import { Link } from "react-router-dom";

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
    const response = await userService.send_email(this.state.data);
    console.log(response);
    let obj = JSON.parse(response.data);

    console.log(obj);
  };

  render() {
    const { data, errors } = this.state;
    return (
      <section class="signup sign">
        <div class="container">
          <div class="signup-content">
            <div class="signup-form">
              <h2 class="form-title">Sign up</h2>
              <form
                onSubmit={this.handleSubmit}
                method="POST"
                class="register-form"
                id="register-form"
              >
                <InputWithButton
                  onChange={this.handleChange}
                  value={data.email}
                  type="email"
                  id="email"
                  errors={errors.email}
                  text="Send"
                  onClick={this.handleClicked}
                  label="fa fa-envelope"
                  placeholder="email"
                />
                <Input
                  onChange={this.handleChange}
                  value={data.verification}
                  type="text"
                  id="verification"
                  errors={errors.verification}
                  label="fa fa-keyboard-o"
                  placeholder="verifaction code"
                />

                <InputWithDesc
                  onChange={this.handleChange}
                  value={data.username}
                  type="text"
                  id="username"
                  errors={errors.username}
                  text="Must be 6-20 characters long."
                  label="fa fa-user"
                  placeholder="username"
                />
                <InputWithDesc
                  onChange={this.handleChange}
                  value={data.password}
                  type="password"
                  id="password"
                  errors={errors.password}
                  text="Must be 8-20 characters long."
                  label="fa fa-unlock-alt"
                  placeholder="password"
                />
                <InputWithDesc
                  onChange={this.handleChange}
                  value={data.repeatpassword}
                  type="repeatpassword"
                  id="repeatpassword"
                  errors={errors.repeatpassword}
                  text="Please repeat your password."
                  label="fa fa-repeat"
                  placeholder="repeat password"
                />

                <div className="form-group mb-3">
                  <input
                    type="checkbox"
                    class="agree-term"
                    id="secretproblem"
                    onClick={this.handlechecked}
                  />
                  <label htmlFor="secretproblem" class="label-agree-term">
                    <span>
                      <span></span>
                    </span>
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
                      errors={errors.answer1}
                      placeholder="What's your birth day?"
                      label="fa fa-question"
                    />
                    <Input
                      onChange={this.handleChange}
                      value={data.answer2}
                      type="text"
                      id="answer2"
                      errors={errors.answer2}
                      placeholder="What's your mother's name?"
                      label="fa fa-question"
                    />
                    <Input
                      onChange={this.handleChange}
                      value={data.answer3}
                      type="text"
                      id="answer3"
                      errors={errors.answer3}
                      placeholder="What's your favorite color?"
                      label="fa fa-question"
                    />
                  </div>
                )}

                <div className="form-group">
                  <input
                    type="checkbox"
                    class="agree-term"
                    id="agreement"
                    onClick={this.handlechecked}
                  />
                  <label class="label-agree-term" htmlFor="agreement">
                    <span>
                      <span></span>
                    </span>
                    I agree all statements in{" "}
                    <a
                      href="#!"
                      data-bs-toggle="modal"
                      data-bs-target="#staticBackdrop"
                      class="term-service"
                    >
                      Terms of service
                    </a>
                  </label>
                </div>

                <Button
                  disabled={this.validate()}
                  label="Create your account!"
                />
              </form>
              <div
                class="modal fade"
                id="staticBackdrop"
                data-bs-backdrop="static"
                data-bs-keyboard="false"
                tabindex="-1"
                aria-labelledby="staticBackdropLabel"
                aria-hidden="true"
              >
                <div class="modal-dialog">
                  <div class="modal-content">
                    <div class="modal-header">
                      <h5 class="modal-title" id="staticBackdropLabel">
                        Modal title
                      </h5>
                      <button
                        type="button"
                        class="btn-close"
                        data-bs-dismiss="modal"
                        aria-label="Close"
                      ></button>
                    </div>
                    <div class="modal-body">...</div>
                    <div class="modal-footer">
                      <button
                        type="button"
                        class="btn btn-secondary"
                        data-bs-dismiss="modal"
                      >
                        Close
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="signup-image">
              <figure>
                <img
                  src={require("../img/signup-image.jpg")}
                  alt="sing up image"
                />
              </figure>
              <Link to="/login" class="signup-image-link">
                I am already member
              </Link>
            </div>
          </div>
        </div>
      </section>
    );
  }
}

export default RegisterForm;
