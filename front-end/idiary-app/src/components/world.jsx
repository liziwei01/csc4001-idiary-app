import React, { Component, PropTypes } from "react";
import axios from "axios";
import ContentImg from "./content-img.jsx";
import moment from "moment";

class WorldDiary extends Component {
  state = {
    WorldDiaryList: [],
  };

  componentDidMount() {
    // console.log("WorldDiary====WorldDiary");

    // 上传 : /diary/world

    axios({
      url: "/api/diary/world",
      method: "get",
      params: {
        user_id: 1,
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

  _renderHeadView(data) {
    //console.log(data)
    return (
      <div className="item">
        <div className="topRightView">
          <div className="nameandtime">
            <div style={{ marginright: "10px" }}>
              <img className="nick-img" src={data.user_profile} />
              <span style={{ marginLeft: "10px" }}>{data.nickname}</span>
            </div>
            {/* <div style={{marginLeft:"600px"}}>{data.sendTime}</div> */}
            <div style={{ marginLeft: "700px" }}>
              {moment(data.db_time).format("YYYY-MM-DD HH:mm:ss")}
            </div>
          </div>
          <div>
            <p style={{ marginTop: "10px" }}>{data.content}</p>
            <ContentImg contentImgUrls={data.image_list || []} />
          </div>
        </div>
      </div>
    );
  }
  like = (id, number) => {
    const arr = this.state.WorldDiaryList.map((ele) => {
      if (ele.diary_id === id) {
        return { ...ele, vote_count: number + 1 };
      }
      return { ...ele };
    });
    //传给后端一个点赞的消息
    this.setState({ WorldDiaryList: arr });
  };

  comment = (id, number) => {
    const arr = this.state.WorldDiaryList.map((ele) => {
      if (ele.diary_id === id) {
        //var isshow = ele.isShowComment;
        return { ...ele, isShowComment: true};
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
        return {
          ...ele,
          isTransfer: true,
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
    const arr = this.state.WorldDiaryList.map((ele) => {
      if (ele.diary_id === id) {
        return {
          ...ele,
          inputValue: "",
          comment_list: [
            ...ele.comment_list,
            {
              nickname: name,
              content: ele.inputValue,
            },
          ],
          comment_count: number + 1,
        };
      }
      return { ...ele };
    });
    //this.props.update(arr);
  };

  render() {
    //渲染列表
    const { WorldDiaryList } = this.state;
    return (
      <div className="listRootViewStyle">
        {WorldDiaryList.map((ele, index) => {
          return (
            <div key={index} style={{ marginTop: 20 }}>
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
                    点赞:{ele.vote_count}
                  </li>

                  <div className="shuxian"></div>
                  <li
                    className="liStyle"
                    onClick={() =>
                      this.comment(ele.diary_id, ele.comment_count)
                    }
                  >
                    评论:{ele.comment_count}
                  </li>

                  <div className="shuxian"></div>
                  <li
                    className="liStyle"
                    onClick={() => this.transferdata(ele.diary_id)}
                  >
                    转发:{ele.share_count}
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
                  <div>
                    <input
                      style={{ width: "500px" }}
                      type="text"
                      onChange={(e) => this.intChange(e, ele.diary_id)}
                      value={ele.inputValue}
                    ></input>
                    <pre></pre>
                    <button
                      onClick={() =>
                        this.btnSure(
                          ele.diary_id,
                          ele.nickname,
                          ele.comment_count
                        )
                      }
                    >
                      评论
                    </button>
                  </div>
                )}
                {ele.isShowComment &&
                  ele.comment_list.map((subEle, subIndex) => {
                    return (
                      <div key={subIndex}>
                        {subEle.nickname} : {subEle.content}
                      </div>
                    );
                  })}
              </div>
            </div>
          );
        })}
      </div>
    );
  }
}

export default WorldDiary;
