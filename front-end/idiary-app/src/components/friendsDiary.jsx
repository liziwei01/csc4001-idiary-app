import React, { Component, PropTypes } from "react";
import axios from "axios";

class FriendsDiary extends Component {
  state = {
    friendsDiaryList: [],
  };

  componentDidMount() {
    console.log("FriendsDiary====FriendsDiary");

    // 上传 : /api/upload/image
    // 获取列表 /api/getxxList
    // 
    // axios({
    //   url: "/api/get/list",
    //   method: "get",
    // }).thent((res) => {
    //   if(res.status === 200) {
    //     this.setState({
    //       list: res.xxxx.x
    //     })
    //   }
    // });
    setTimeout(() => {
      const res = {
        list: [
          { name: "ch", age: 27 },
          { name: "zhangsan", age: 20 },
          { name: "xx", age: 18 },
        ],
        success: true,
        code: 200,
      };
      if (res.code === 200) {
        this.setState({
          friendsDiaryList: res.list,
        });
      }
    }, 2000);
  }

  render() {
    const { friendsDiaryList } = this.state;
    return (
      <div>
        <h1> Friends' Diary</h1>

        {friendsDiaryList.map((ele, i) => {
          return <div key={i}>{`name: ${ele.name} , age: ${ele.age}`}</div>;
        })}
      </div>
    );
  }
}

export default FriendsDiary;
