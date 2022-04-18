import React, { Component } from "react";
import { Link } from "react-router-dom";
import * as userService from "./../service/userService";
import AdminTable from "./adminTable";
import Pagination from "./common/pagination";
import { paginate } from "../utils/paginate";
import _ from "lodash";
import SearchBox from "./common/searchBox";

class Admin extends Component {
  state = {
    users: [],
    resetuser: "",
    resetpassword: "",
    currentPage: 1,
    pageSize: 4,
    searchQuery: "",
  };
  async componentDidMount() {
    const user = localStorage.getItem("token");
    const response = await userService.getallinfo(user);
    const data = response.data.data;
    let users = [];
    data.map((item) => users.push(item));

    this.setState({ users });
  }
  handleReset = (user) => {
    this.setState({ resetuser: user.email });
  };

  handlePageChange = (page) => {
    this.setState({ currentPage: page });
  };

  handleSearch = (query) => {
    this.setState({ searchQuery: query, currentPage: 1 });
  };

  handleChange = ({ target: input }) => {
    this.setState({ resetpassword: input.value });
  };

  handleClick = async () => {
    console.log(this.state.resetpassword);
    let data = { email: "", password: "" };
    data.email = this.state.resetuser;
    data.password = this.state.resetpassword;
    const response = await userService.resetpassword(data);
  };

  getPagedData = () => {
    const { pageSize, currentPage, searchQuery, users: allUsers } = this.state;

    let filtered = allUsers;
    if (searchQuery)
      filtered = allUsers.filter((m) =>
        m.email.toLowerCase().startsWith(searchQuery.toLowerCase())
      );

    const users = paginate(filtered, currentPage, pageSize);

    return { totalCount: filtered.length, data: users };
  };
  render() {
    const { length: count } = this.state.users;
    const { pageSize, currentPage, searchQuery, resetpassword, resetuser } =
      this.state;
    const { user } = this.props;

    if (count === 0) return <p>There are no movies in the database.</p>;

    const { totalCount, data: users } = this.getPagedData();

    return (
      <div>
        <div>
          {user && (
            <Link
              to="/movies/new"
              className="btn btn-primary"
              style={{ marginBottom: 20 }}
            >
              New Movie
            </Link>
          )}
          <p>Showing {totalCount} users in the database.</p>
          <SearchBox value={searchQuery} onChange={this.handleSearch} />
          <AdminTable users={users} onReset={this.handleReset} />
          <Pagination
            itemsCount={totalCount}
            pageSize={pageSize}
            currentPage={currentPage}
            onPageChange={this.handlePageChange}
          />
        </div>
        <div
          class="modal fade"
          id="exampleModal"
          tabindex="-1"
          aria-labelledby="exampleModalLabel"
          aria-hidden="true"
        >
          <div class="modal-dialog">
            <div class="modal-content">
              <div class="modal-header">
                <h5 class="modal-title" id="exampleModalLabel">
                  Reset password for user {resetuser}
                </h5>
                <button
                  type="button"
                  class="btn-close"
                  data-bs-dismiss="modal"
                  aria-label="Close"
                ></button>
              </div>
              <div class="modal-body">
                <label for="exampleInputEmail1" class="form-label">
                  New password
                </label>
                <input
                  type="text"
                  class="form-control"
                  id="exampleInputEmail1"
                  aria-describedby="emailHelp"
                  onChange={this.handleChange}
                />
                <div id="emailHelp" class="form-text">
                  The new password should between 8-20 characters.
                </div>
              </div>
              <div class="modal-footer">
                <button
                  type="button"
                  class="btn btn-secondary"
                  data-bs-dismiss="modal"
                >
                  Close
                </button>
                <button
                  type="button"
                  class="btn btn-primary"
                  onClick={this.handleClick}
                >
                  Reset
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default Admin;
