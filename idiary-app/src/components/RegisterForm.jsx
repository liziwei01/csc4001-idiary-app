import React, { Component } from "react";

class RegisterForm extends Component {
  state = { data: { email: "", password: "", repeatpassword: "" }, errors: {} };
  handleSubmit = (e) => {
    e.preventDefault();
    //server
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
      })
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
      if (value.length < 6 || value.length > 20)
        return "Your password should between 6 and 20 characters!";
    }
    if (id === "repeatpassword") {
      if (value !== this.state.data.password) return "Different!";
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
        <h1>Register</h1>
        <form onSubmit={this.handleSubmit}>
          <div className="mb-3">
            <label htmlFor="email" className="form-label">
              Email address
            </label>
            <input
              onChange={this.handleChange}
              value={data.email}
              type="email"
              className="form-control"
              id="email"
              aria-describedby="emailHelp"
            />
            <div id="emailHelp" className="form-text">
              We'll never share your email with anyone else.
            </div>
            {errors.email && (
              <div className="alert alert-danger">{errors.email}</div>
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
          <div className="mb-3">
            <label htmlFor="repeatpassword" className="form-label">
              Repeat Your Password
            </label>
            <input
              onChange={this.handleChange}
              value={data.repeatpassword}
              type="password"
              className="form-control"
              id="repeatpassword"
            />
            {errors.repeatpassword && (
              <div className="alert alert-danger">{errors.repeatpassword}</div>
            )}
          </div>
          <div className="mb-3 form-check">
            <input type="checkbox" className="form-check-input" id="check" />
            <label className="form-check-label" htmlFor="check">
              Check me out
            </label>
            {errors.check && (
              <div className="alert alert-danger">{errors.check}</div>
            )}
          </div>
          <button
            disabled={this.validate()}
            type="submit"
            className="btn btn-primary"
          >
            Submit
          </button>
        </form>
      </div>
    );
  }
}

export default RegisterForm;
