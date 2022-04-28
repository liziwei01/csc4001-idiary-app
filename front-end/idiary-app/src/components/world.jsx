import React, { Component, PropTypes } from "react";
import axios from "axios";
import ContentImg from "./content-img.jsx";
import moment from "moment";
import { Button } from "antd";
import { CheckOutlined } from "@ant-design/icons";
import auth from "../service/authService";
import * as userService from "../service/userService";
import { LikeOutlined, LikeFilled, CommentOutlined } from "@ant-design/icons";
import Pagination from "./common/pagination.jsx";
import _ from "lodash";
import { paginate } from "../utils/paginate";
class WorldDiary extends Component {
  state = {
    WorldDiaryList: [],
    user_id: null,
    user_name: null,
    currentPage: 1,
    pageSize: 6,
    sortColumn: { path: "sendTime", order: "asc" },
  };

  componentDidMount = async () => {
    // console.log("WorldDiary====WorldDiary");

    // 上传 : /diary/world

    const user = auth.getCurrentUser();
    const response = await userService.getinfobyemail(user);
    const user_id = response.data.data.user_id;
    const user_name = response.data.data.nickname;
    this.setState({ user_id, user_name });

    axios({
      url: "/api/diary/world",
      method: "get",
      params: {
        user_id: this.state.user_id,
      },
    }).then((res) => {
      console.log(res);
      console.log(res.data.data.diaries[0].user_profile);
      console.log("WorldDiary====WorldDiary");
      if (res.status === 200) {
        this.setState({
          WorldDiaryList: (res.data.data && res.data.data.diaries) || [],
        });
      }
    });
  }

  follow = (id) => {
    const arr = this.state.WorldDiaryList.map((ele) => {
      if (ele.user_id === id) {
        var follow = ele.has_followed;
        axios({
          url: "/api/user/follow/follow",
          method: "post",
          params: {
            user_id: this.state.user_id,
            following_id: id,
            should_unfollow: follow,
          },
        }).then((res) => {
          console.log("post/follow/follow");
        });
        return { ...ele, has_followed: !follow };
      }
      return { ...ele };
    });
    this.setState({ WorldDiaryList: arr });
  };

  _renderHeadView(data) {
    //console.log(data)
    var follow = "follow";
    return (
      <div className="item container">
        <div className="topRightView">
          <div className="nameandtime">
            <div style={{ marginright: "10px" }}>
              <img className="nick-img" src={data.user_profile} />
              <span style={{ marginLeft: "10px", fontWeight: "600", fontSize: "16px" }}>{data.nickname}</span>

              {!(data.user_id === this.state.user_id) &&
                <button
                  className="btn btn-sm ms-3"
                  type="button"
                  onClick={() => this.follow(data.user_id)}
                  style={{ backgroundColor: "red", borderRadius: "15px" }}
                >
                  {data.has_followed ? <div> <i class="fa fa-check" style={{ color: "white", display: "inline-block" }}></i><p className="mb-0 ms-3" style={{ color: "white", display: "inline-block" }}>cancel</p></div> : <div><i class="fa fa-heart" aria-hidden="true" style={{ color: "white", display: "inline-block" }}></i><p className="ms-3 mb-0" style={{ color: "white", display: "inline-block" }}>follow</p></div>}
                </button>}

            </div>
            {/* <div style={{marginLeft:"600px"}}>{data.sendTime}</div> */}
            <div style={{ flex: 1, textAlign: "right", color: "grey" }}>
              {moment(data.db_time * 1000).format("YYYY-MM-DD HH:mm:ss")}
            </div>
          </div>
          <div>
            <p style={{ marginTop: "20px" }}>{data.content}</p>
            <ContentImg contentImgUrls={data.image_list || []} />
          </div>
        </div>
      </div >
    );
  }
  like = (id, number) => {
    const arr = this.state.WorldDiaryList.map((ele) => {
      var vote = ele.has_voted;
      if (ele.diary_id === id) {
        axios({
          url: "/api/diary/like",
          method: "post",
          params: {
            user_id: this.state.user_id,
            diary_id: id,
            should_unlike: vote,
          },
        }).then((res) => {

        });
        var result;
        if (vote) {
          result = number - 1;
        } else {
          result = number + 1;
        }
        return { ...ele, vote_count: result, has_voted: !vote };
      }
      return { ...ele };
    });
    //传给后端一个点赞的消息
    this.setState({ WorldDiaryList: arr });
  };

  comment = (id, number) => {
    const arr = this.state.WorldDiaryList.map((ele) => {
      if (ele.diary_id === id) {
        var isshow = ele.isShowComment;
        return { ...ele, isShowComment: !isshow };
      }
      return { ...ele };
    });
    //this.props.update(arr);
    //传给后端新评论
    this.setState({ WorldDiaryList: arr });
  };

  transferdata = (id) => {
    const arr = this.state.WorldDiaryList.map((ele) => {
      if (ele.diary_id === id) {
        var transfer = ele.isTransfer;
        return {
          ...ele,
          isTransfer: !transfer,
        };
      }
      return { ...ele };
    });
    //this.props.update(arr);
    //传给后端转发的消息
    this.setState({ WorldDiaryList: arr });
  };

  transferChange = (e, id) => {
    const arr = this.state.WorldDiaryList.map((ele) => {
      if (ele.diary_id === id) {
        return { ...ele, transferValue: e.target.value };
      }
      return { ...ele };
    });
    this.setState({ WorldDiaryList: arr });
  };

  transfer = (id, number) => {
    // 增加转发数量
    const arr = this.state.WorldDiaryList.map((ele) => {
      if (ele.diary_id === id) {
        return {
          ...ele,
          transferValue: "",
          share_count: number + 1,
        };
      }
      return { ...ele };
    });

    //this.props.update(arr);
    //传给后端
    //自己发送日记
    this.setState({ WorldDiaryList: arr });
  };

  intChange = (e, id) => {
    const arr = this.state.WorldDiaryList.map((ele) => {
      if (ele.diary_id === id) {
        return { ...ele, inputValue: e.target.value };
      }
      return { ...ele };
    });
    //this.props.update(arr);
    this.setState({ WorldDiaryList: arr });
  };

  btnSure = (id, name, number) => {
    // 拿到后发送给服务端
    var content;
    const arr = this.state.WorldDiaryList.map((ele) => {
      if (ele.diary_id === id) {
        if (ele.comment_count === 0) {
          var commentlist = [{ nickname: this.state.user_name, content: ele.inputValue },];
        } else {
          var commentlist = [
            ...ele.comment_list,
            {
              nickname: this.state.user_name,
              content: ele.inputValue,
            },
          ];
        }
        content = ele.inputValue;
        return {
          ...ele,
          inputValue: "",
          comment_list: commentlist,
          comment_count: number + 1,
        };
      }

      return { ...ele };
    });
    //this.props.update(arr)

    console.log("comment", this.state.user_id)

    this.setState({ WorldDiaryList: arr });
    axios({
      url: "/api/diary/comment",
      method: "post",
      params: {
        user_id: this.state.user_id,
        diary_id: id,
        content: content,
      },
    }).then((res) => { });
  };
  getPagedData = () => {
    const {
      pageSize,
      currentPage,
      sortColumn,
      WorldDiaryList: allDiaries,
    } = this.state;

    let filtered = allDiaries;
    const sorted = _.orderBy(filtered, [sortColumn.path], [sortColumn.order]);

    const diaries = paginate(sorted, currentPage, pageSize);

    return { totalCount: filtered.length, data: diaries };
  };

  handlePageChange = (page) => {
    this.setState({ currentPage: page });
  };

  render() {
    //渲染列表
    const { WorldDiaryList } = this.state;
    const { pageSize, currentPage } = this.state;
    const { totalCount, data: diaries } = this.getPagedData();
    return (
      <div className="listRootViewStyle">
        {diaries.map((ele, index) => {
          return (
            <div key={index} style={{ marginTop: 10 }}>
              {this._renderHeadView(ele)}
              <hr className="hrStyle" />
              {/* footer */}
              <div className="commentViewStyle">
                <ul className="ulStyle">
                  <div className="shuxian"></div>
                  <li
                    className="liStyle"
                    onClick={() => this.like(ele.diary_id, ele.vote_count)}
                  >
                    {!ele.has_voted ? <LikeOutlined style={{ marginTop: "5px", marginRight: "6px" }} /> : <LikeFilled style={{ marginTop: "5px", marginRight: "6px" }} />}: {ele.vote_count}
                  </li>

                  <div className="shuxian"></div>
                  <li
                    className="liStyle"
                    onClick={() =>
                      this.comment(ele.diary_id, ele.comment_count)
                    }
                  >
                    <CommentOutlined style={{ marginTop: "5px", marginRight: "6px" }} />comments : {ele.comment_count}
                  </li>

                </ul>

                {ele.isTransfer && (
                  <div>
                    <input
                      style={{ width: "500px" }}
                      type="text"
                      onChange={(e) => this.transferChange(e, ele.diary_id)}
                      value={ele.transferValue}
                    />
                    <pre></pre>
                    <button
                      onClick={() =>
                        this.transfer(ele.diary_id, ele.share_count)
                      }
                    >
                      转发
                    </button>
                  </div>
                )}

                {ele.isShowComment && (
                  <div className="row">
                    <div className="col-auto">
                      <button
                        className="btn btn-sm ms-3"
                        onClick={() =>
                          this.btnSure(
                            ele.diary_id,
                            ele.nickname,
                            ele.comment_count
                          )
                        }
                        style={{ backgroundColor: "#ffbb49", color: "white", borderRadius: "8px" }}
                      >
                        post
                      </button>
                    </div>
                    <div className="col">
                      <input
                        style={{ width: "500px" }}
                        type="text"
                        onChange={(e) => this.intChange(e, ele.diary_id)}
                        value={ele.inputValue}
                        placeholder="Say something..."
                        className="commentinput"
                      ></input>
                    </div>
                    <pre></pre>

                  </div>
                )}
                {ele.isShowComment &&
                  !(ele.comment_count === 0) &&
                  ele.comment_list.map((subEle, subIndex) => {
                    return (
                      <div key={subIndex} style={{ backgroundColor: "white", padding: "10px" }}>
                        {subEle.nickname} : {subEle.content}
                      </div>
                    );
                  })}
              </div>
            </div>
          );
        })}
        <Pagination
          itemsCount={totalCount}
          pageSize={pageSize}
          currentPage={currentPage}
          onPageChange={this.handlePageChange}
        />
      </div>
    );
  }
}

export default WorldDiary;
