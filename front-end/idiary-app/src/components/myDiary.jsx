
import { Route, Routes } from "react-router-dom";
import PersonalDiary from "./personalDiary";
import FriendsDiary from "./friendsDiary";
import Tab from "./common/tab";
import React, { Component, PropTypes } from "react";
export default class MyDiary extends Component {
    state = {
        writeVal:""
      };
      writeChange = (e) => {
        this.setState({
          writeVal: e.target.value,
        });
      };
    
      write = () => {
        // 拿到后发送给服务端
      };
  render(){

    return (
        <React.Fragment>
          
          <div className="wrap">
              <div className="says">
                <h4>write a new diary</h4>
                <textarea
                  style={{ width: "500px", height: "150px" }}
                  type="text"
                  onChange={this.writeChange}
                  value={this.state.writeVal}
                />
                <div>
                  <button
                    onClick={() => this.write()}
                    style={{ marginTop: "10px" }}>
                    publish
                  </button>
                </div>
              </div>
          </div>
          
          <Tab />
          <Routes>
            <Route path="personalDiary" element={<PersonalDiary />} />
            <Route path="friendsDiary" element={<FriendsDiary />} />
          </Routes>
        </React.Fragment>
      );
  }  
  
};


