import { Route, Routes } from "react-router-dom";
import PersonalDiary from "./personalDiary";
import FriendsDiary from "./friendsDiary";
import Tab from "./common/tab";
import React, { Component, PropTypes } from "react";
import PicturesWall from "./upload_img";
import ReactDOM from "react-dom";
export default class MyDiary extends Component {
  state = {
    writeVal: "",
    diary_type: "",
  };
  writeChange = (e) => {
    this.setState({
      writeVal: e.target.value,
    });
  };

  write = () => {
    // 拿到后发送给服务端
  };

  handleChange = (event) => {
    const { value } = event.currentTarget;
    this.setState({ diary_type: value });
    // console.log(this.state.diary_type);
  };

  render() {
    return (
      <React.Fragment>
        <div className="wrap">
          <div className="says">
            <h4>write a new diary</h4>
            <div style={{ display: "flex" }}>
              <textarea
                style={{ width: "500px", height: "150px" }}
                type="text"
                onChange={this.writeChange}
                value={this.state.writeVal}
              />
              
              
            </div>
            <div style={{display:"flex"}}>
            <div>
              <button
                onClick={() => this.write()}
                style={{ marginTop: "10px" }}
              >
                publish
              </button>
            </div>
            <div style={{ margin:"10px" }}>
                <select
                  value={this.state.diary_type}
                  onChange={this.handleChange}
                >
                  <option value="Public">Public</option>
                  <option value="Protected">Protected</option>
                  <option value="Private">Private</option>
                </select>
              </div>
            </div>
            
          </div>
        </div>
        <div id="upload-image"></div>

        <Tab />
        <Routes>
          <Route path="personalDiary" element={<PersonalDiary />} />
          <Route path="friendsDiary" element={<FriendsDiary />} />
        </Routes>
      </React.Fragment>
    );
  }
}
//ReactDOM.render(<PicturesWall />, document.getElementById("upload-image"));
