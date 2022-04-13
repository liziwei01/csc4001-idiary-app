import React, { Component, PropTypes} from "react";
import ContentImg from "./content-img.jsx";
import "../css/listWb.css";
import PicturesWall from "./upload_img.jsx"
import moment from 'moment';
export default class WeiBoList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      inputVal: "",
      transferVal: "",
      writeVal: "",
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
          <div className="nameandtime">
            <div style={{marginright:"10px"}}>
              <img className="nick-img" src={data.headUrl} />
              <span style={{ marginLeft: "10px" }}>{data.nickName}</span>
              
            </div>
            {/* <div style={{marginLeft:"600px"}}>{data.sendTime}</div> */}
            <div style={{marginLeft:"500px"}}>{moment(data.sendTime).format("YYYY-MM-DD HH:mm:ss")}</div>
            
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
          var isshow = ele.isShowComment;
        return { ...ele, isShowComment: !isshow };
      }
      return { ...ele };
    });
    this.props.update(arr);
  };

  intChange = (e, id) => {
    const arr = this.props.data.map((ele) => {
      if (ele.id === id) {
        return { ...ele, inputValue: e.target.value };
      }
      return { ...ele };
    });
    this.props.update(arr);
  };

  btnSure = (id, name) => {
    // 拿到后发送给服务端
    const arr = this.props.data.map((ele) => {
      if (ele.id === id) {
        return {
          ...ele,
          inputValue: "",
          commentList: [
            ...ele.commentList,
            {
              name: name,
              comment: ele.inputValue,
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
        var istransfer = ele.isTransfer;
        return { 
            ...ele, 
            isTransfer: !istransfer };
      }
      return { ...ele };
    });
    this.props.update(arr);
  };

  transferChange = (e,id) => {
    const arr = this.props.data.map((ele) => {
        if (ele.id === id) {
          return { ...ele, transferValue: e.target.value };
        }
        return { ...ele };
      });
      this.props.update(arr);
    
  };

  transfer = (id, number) => {
    // 增加转发数量
    const arr = this.props.data.map((ele) => {
      if (ele.id === id) {
        return { 
            ...ele, 
            transferValue:"",
            NoPointGreat: number + 1 };
      }
      return { ...ele };
    });
    console.log(arr);
    this.props.update(arr);

    //自己发送日记
  };

  writeChange = (e) => {
    this.setState({
      writeVal: e.target.value,
    });
  };

  write = () => {
    // 拿到后发送给服务端
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
                    onClick={() => this.like(ele.id, ele.NoForward)}>
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
                      style={{ width: "500px" }}
                      type="text"
                      onChange={(e) => this.transferChange(e, ele.id)}
                      value={ele.transferValue}
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
                      style={{ width: "500px" }}
                      type="text"
                      onChange={(e) => this.intChange(e, ele.id)}
                      value={ele.inputValue}
                    ></input>
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
