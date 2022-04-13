import React, { Component } from "react";
import Table from "./common/table";

class AdminTable extends Component {
  columns = [
    {
      path: "username",
      label: "Username",
    },
    { path: "email", label: "Email" },
    { path: "password", label: "Password" },
    {
      key: "reset",
      content: (user) => (
        <button
          onClick={() => this.props.onReset(user)}
          className="btn btn-danger btn-sm "
          data-bs-toggle="modal"
          data-bs-target="#exampleModal"
        >
          Reset Password
        </button>
      ),
    },
  ];

  render() {
    const { users } = this.props;

    return <Table columns={this.columns} data={users} />;
  }
}

export default AdminTable;
