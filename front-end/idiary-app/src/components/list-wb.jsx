import React, { Component, PropTypes } from "react";
import ContentImg from "./content-img.jsx";

export default class WeiBoList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      inputVal: "",
      transferVal: "",
    };
  }

  /**
   * 渲染顶部View
   */
  _renderHeadView(data) {
    var contentImgs;
    if (data.contentImgUrls) {
      //若后端给的数据中有图片url，则展示
      contentImgs = <ContentImg content-img-urls={data.contentImgUrls} />;
    }
    return (
      <div className="item">
        <div className="topRightView">
          <div className="nickNameAndSendTime">
            <div>
              <img className="nick-img" src={data.headUrl} />
              <span style={{ marginLeft: "10px" }}>{data.nickName}</span>
            </div>

            <span>{data.sendTime}</span>
          </div>
          <div>
            <p style={{ marginTop: "10px" }}>{data.content}</p>
            <div>{contentImgs}</div>
          </div>
        </div>
      </div>
    );
  }

  like = (id, number) => {
    const arr = this.props.data.map((ele) => {
      if (ele.id === id) {
        return { ...ele, NoForward: number + 1 };
      }
      return { ...ele };
    });
    this.props.update(arr);
  };

  comment = (id) => {
    const arr = this.props.data.map((ele) => {
      if (ele.id === id) {
        return { ...ele, isShowComment: true };
      }
      return { ...ele };
    });
    this.props.update(arr);
  };

  intChange = (e) => {
    this.setState({
      inputVal: e.target.value,
    });
  };

  btnSure = (id, name) => {
    // 拿到后发送给服务端
    const arr = this.props.data.map((ele) => {
      if (ele.id === id) {
        return {
          ...ele,
          commentList: [
            ...ele.commentList,
            {
              name: name,
              comment: this.state.inputVal,
            },
          ],
        };
      }
      return { ...ele };
    });
    this.props.update(arr);
  };

  transferdata = (id) => {
    const arr = this.props.data.map((ele) => {
      if (ele.id === id) {
        return { ...ele, isTransfer: true };
      }
      return { ...ele };
    });
    this.props.update(arr);
  };

  transferChange = (e) => {
    this.setState({
      transferVal: e.target.value,
    });
  };

  transfer = (id, number) => {
    // 增加转发数量
    const arr = this.props.data.map((ele) => {
      if (ele.id === id) {
        return { ...ele, NoPointGreat: number + 1 };
      }
      return { ...ele };
    });
    this.props.update(arr);

    //自己发送日记
  };

  render() {
    //渲染列表
    return (
      <div className="listRootViewStyle">
        {this.props.data.map((ele, index) => {
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
                    onClick={() => this.like(ele.id, ele.NoForward)}
                  >
                    点赞:{ele.NoForward}
                  </li>

                  <div className="shuxian"></div>
                  <li className="liStyle" onClick={() => this.comment(ele.id)}>
                    评论:{ele.NoComment}
                  </li>

                  <div className="shuxian"></div>
                  <li
                    className="liStyle"
                    onClick={() => this.transferdata(ele.id)}
                  >
                    转发:{ele.NoPointGreat}
                  </li>
                </ul>

                {ele.isTransfer && (
                  <div>
                    <input
                      style={{ width: "1000px" }}
                      type="text"
                      onChange={this.transferChange}
                      value={this.state.transferVal}
                    />
                    <pre></pre>
                    <button
                      onClick={() => this.transfer(ele.id, ele.NoPointGreat)}
                    >
                      转发
                    </button>
                  </div>
                )}

                {ele.isShowComment && (
                  <div>
                    <input
                      style={{ width: "1000px" }}
                      type="text"
                      onChange={this.intChange}
                      value={this.state.inputVal}
                    />
                    <pre></pre>
                    <button onClick={() => this.btnSure(ele.id, ele.nickName)}>
                      评论
                    </button>
                  </div>
                )}
                {ele.isShowComment &&
                  ele.commentList.map((subEle, subIndex) => {
                    return (
                      <div key={subIndex}>
                        {subEle.name} : {subEle.comment}
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

WeiBoList.propTypes = {};

//export default ListWB;
