import React, { Component, PropTypes } from "react";
import { Route, Routes } from "react-router-dom";
import axios from "axios";
import PersonalDiary from "./personalDiary";
import FriendsDiary from "./friendsDiary";

import Tab from "./common/tab";

import { Upload, Button, message } from "antd";
import { LoadingOutlined, PlusOutlined } from "@ant-design/icons";

export default class MyDiary extends Component {
  state = {
    Tabflag: "",
    writeVal: "",
    diary_type: "",
    list: [],
    articleData: {
      img: "",
      text: "",
    },
    imageUrl: "",
    loading: false,
  };
  writeChange = (e) => {
    this.setState({
      writeVal: e.target.value,
    });
  };

  write = () => {
    // 拿到后发送给服务端
    var t = new Date().getTime();
    axios({
      url: "/api/diary/addDiary",
      method: "post",
      data: {
        user_id: 1,
        content: this.state.writeVal,
        image_list: "",
        authority: this.state.diary_type,
      },
    }).then((res) => {
      if (res.status === 200) {
        console.log("success");
      }
    });
  };
  // react一个生命周期

  handleChange = (event) => {
    const { value } = event.currentTarget;
    this.setState({ diary_type: value });
    // console.log(this.state.diary_type);
  };

  customRequest = async (obj) => {
    const { file, onSuccess, onError } = obj;
    const suffix = file.name.split(".").reverse()[0];
    // console.log(file, "==file");
    // console.log(suffix, "===suffix");
    const values = {
      user_id: 1,
      file,
    };
    const formData = new FormData();
    Object.keys(values).forEach((key) => {
      formData.append(key, values[key]);
    });
    axios({
      method: "post",
      url: `/api/upload/image`,
      data: formData,
    }).then((res) => {
      // console.log(res);
      if (res.status === 200) {
        this.setState({
          imageUrl: res.data.data.url,
        });
        onSuccess();
      }
    });
  };

  render() {
    const { imageUrl, loading } = this.state;
    const uploadButton = (
      <div>
        {loading ? <LoadingOutlined /> : <PlusOutlined />}
        <div style={{ marginTop: 8 }}>Upload</div>
      </div>
    );
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
              <Upload
                name="file"
                // action="www.baidu.com/api/upload/image"
                // beforeUpload={this.beforeUpload}
                customRequest={this.customRequest}
                // listType="picture-card"
                showUploadList={false}
                // accept=".jpg.png.jpeg"
                // onRemove={this.onRemove}
                // showUploadList={{ showPreviewIcon: false, showDownloadIcon: false }}
              >
                {imageUrl ? (
                  <img
                    src={imageUrl}
                    alt="avatar"
                    style={{ width: 100, height: 100 }}
                  />
                ) : null}
                {uploadButton}
              </Upload>
            </div>
            <div style={{ display: "flex" }}>
              <div>
                <button
                  onClick={() => this.write()}
                  style={{ marginTop: "10px" }}
                >
                  publish
                </button>
              </div>
              <div style={{ margin: "10px" }}>
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
        {/* <div id="upload-image"></div> */}

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
