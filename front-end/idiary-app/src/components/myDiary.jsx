import React, { Component, PropTypes } from "react";
import { Route, Routes } from "react-router-dom";
import axios from "axios";
import PersonalDiary from "./personalDiary";
import FriendsDiary from "./friendsDiary";

import Tab from "./common/tab";

// import PicturesWall from "./upload_img";
import { Upload, Button, message } from "antd";
import { UploadOutlined } from "@ant-design/icons";

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
  };
  writeChange = (e) => {
    this.setState({
      writeVal: e.target.value,
    });
  };

  write = () => {
    // 拿到后发送给服务端
    var t = (new Date()).getTime();
    axios({
        url: "/api/diary/addDiary",
        method: "post",
        data:{
           user_id: 1,
           content: this.state.writeVal,
           image_list: "",
           authority: this.state.diary_type,
           time: t,
        }
      }).then((res) => {
          
        if(res.status === 200) {
            console.log("success")
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
    axios({
      method: "post",
      url: `/api/upload/image`,
      headers: {
        // "access-uid": jsCookie.get("accessUid"),
        // Authorization: jsCookie.get("adminToken"),
      },
    }).then((res) => {
      console.log(res);
      if (res.status === 200) {
        // this.setState({
        //   articleData: { ...this.setState.articleData, img: res.data.data.url },
        // });
      }

      // if (res.status === 200) {
      //   if (!res.data.data[0]) return message.error(res);
      //   const data = res.data.data[0];
      //   axios({
      //     method: "put",
      //     url: data.url,
      //     data: file,
      //     headers: {
      //       "Content-Type": "binary/octet-stream",
      //     },
      //   })
      //     .then((res) => {
      //       if (res.status === 200) {
      //         const uploadData = {
      //           fileName: file.name,
      //           fileKey: data.key,
      //         };
      //         onSuccess(true);
      //       }
      //     })
      //     .catch(() => {
      //       onError(true);
      //     });
      // } else {
      //   message.error(res.status);
      // }
    });
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
              <Upload
                name="file"
                // beforeUpload={this.beforeUpload}
                //customRequest={this.customRequest}
                listType="image"
                // accept=".jpg.png.jpeg"
                // onRemove={this.onRemove}
                // showUploadList={{ showPreviewIcon: false, showDownloadIcon: false }}
              >
                <Button icon={<UploadOutlined />}>{"upload image"}</Button>
                {/* {dataLength > 0 && (
                  <span style={{ marginLeft: "10px" }}>
                    {t("共")} {dataLength} {t("数据")}
                  </span>
                )} */}
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
