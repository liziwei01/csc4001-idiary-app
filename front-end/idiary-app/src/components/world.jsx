import React, { Component, PropTypes } from "react";
import axios from "axios";
import ContentImg from "./content-img.jsx";
import moment from 'moment';


class WorldDiary extends Component {
  state = {
    WorldDiaryList: [],
  };

  componentDidMount() {
    console.log("WorldDiary====WorldDiary");

    // 上传 : /diary/world

    axios({
      url: "/api/diary/world",
      method: "get",
    }).then((res) => {
        console.log(res)
      if(res.status === 200) {
        this.setState({
          WorldDiaryList: res.data.diaries
        })
      }
    });
  }

  render() {
    const { WorldDiaryList } = this.state;
    return (
      <div>
        <h1> WorldDiary</h1>

        {WorldDiaryList.map((ele, i) => {
          return <div key={i}>{`name: ${ele.name} , age: ${ele.age}`}</div>;
        })}
      </div>
    );
  }
// _renderHeadView(data) {
//     var contentImgs;
//     // if (data.contentImgUrls) {
//     //   //若后端给的数据中有图片url，则展示
//     //   contentImgs = <ContentImg content-img-urls={data.contentImgUrls} />;
//     // }
//     contentImgs = <ContentImg content-img-urls={data.image_list} />;
//     return (
//       <div className="item">
//         <div className="topRightView">
//           <div className="nameandtime">
//             <div style={{marginright:"10px"}}>
//               {/* <img className="nick-img" src={data.headUrl} /> */}
//               <span style={{ marginLeft: "10px" }}>{data.user_id}</span>
              
//             </div>
//             {/* <div style={{marginLeft:"600px"}}>{data.sendTime}</div> */}
//             <div style={{marginLeft:"500px"}}>{moment(data.db_time).format("YYYY-MM-DD HH:mm:ss")}</div>
            
//           </div>
//           <div>
//             <p style={{ marginTop: "10px" }}>{data.content}</p>
//             <div>{contentImgs}</div>
//           </div>
//         </div>
//       </div>
//     );
//   }
//   like = (id, number) => {
//     const arr = this.state.WorldDiaryList.map((ele) => {
//       if (ele.id === id) {
//         return { ...ele, vote_count: number + 1 };
//       }
//       return { ...ele };
//     });
//     //传给后端一个点赞的消息

//   };

//   comment = (id,number) => {
//     const arr = this.state.WorldDiaryList.map((ele) => {
//       if (ele.id === id) {
//           var isshow = ele.isShowComment;
//         return { ...ele, isShowComment: !isshow, share_count: number+1 };
//       }
//       return { ...ele };
//     });
//     //this.props.update(arr);
//     //传给后端新评论
//   };

//   transferdata = (id) => {
//     const arr = this.state.WorldDiaryList.map((ele) => {
//       if (ele.id === id) {
//         var istransfer = ele.isTransfer;
//         return { 
//             ...ele, 
//             isTransfer: !istransfer
//             };
//       }
//       return { ...ele };
//     });
//     //this.props.update(arr);
//     //传给后端转发的消息
//   };

//   transferChange = (e,id) => {
//     const arr = this.state.WorldDiaryList.map((ele) => {
//         if (ele.diary_id === id) {
//           return { ...ele, transferValue: e.target.value };
//         }
//         return { ...ele };
//       });
//     //this.props.update(arr);
//     console.log(arr);
//     this.setState({WorldDiaryList:arr})
    
//   };

//   transfer = (id, number) => {
//     // 增加转发数量
//     const arr = this.state.WorldDiaryList.map((ele) => {
//       if (ele.diary_id === id) {
//         return { 
//             ...ele, 
//             transferValue:"",
//             report_count: number + 1 };
//       }
//       return { ...ele };
//     });
//     console.log(arr);
//     //this.props.update(arr);
//     //传给后端
//     //自己发送日记
//   };

//   intChange = (e, id) => {
//     const arr = this.state.WorldDiaryList.map((ele) => {
//       if (ele.diary_id === id) {
//         return { ...ele, inputValue: e.target.value };
//       }
//       return { ...ele };
//     });
//     //this.props.update(arr);
//   };

//   btnSure = (id, name) => {
//     // 拿到后发送给服务端
//     const arr = this.state.WorldDiaryList.map((ele) => {
//       if (ele.diary_id === id) {
//         return {
//           ...ele,
//           inputValue: "",
//           commentList: [
//             ...ele.commentList,
//             {
//               user_name: name,
//               comment: ele.inputValue,
//             },
//           ],
//         };
//       }
//       return { ...ele };
//     });
//     //this.props.update(arr);
//   };

// render() {
//     //渲染列表
//     const { WorldDiaryList } = this.state;
//     return (
//       <div className="listRootViewStyle">
        
//         {WorldDiaryList.map((ele, index) => {
//           return (
//             <div key={index} style={{ marginTop: 20 }}>
//               {this._renderHeadView(ele)}
//               <hr className="hrStyle" />
//               {/* footer */}
//               <div className="commentViewStyle">
//                 <ul className="ulStyle">
//                   <div className="shuxian"></div>
//                   <li
//                     className="liStyle"
//                     onClick={() => this.like(ele.diary_id, ele.vote_count)}>
//                     点赞:{ele.vote_count}
//                   </li>

//                   <div className="shuxian"></div>
//                   <li className="liStyle" onClick={() => this.comment(ele.diary_id,ele.share_count)}>
//                     评论:{ele.share_count}
//                   </li>

//                   <div className="shuxian"></div>
//                   <li
//                     className="liStyle"
//                     onClick={() => this.transferdata(ele.diary_id)}
//                   >
//                     转发:{ele.report_count}
//                   </li>
//                 </ul>

//                 {ele.isTransfer && (
//                   <div>
//                     <input
//                       style={{ width: "500px" }}
//                       type="text"
//                       onChange={(e) => this.transferChange(e, ele.diary_id)}
//                       value={ele.transferValue}
//                     />
//                     <pre></pre>
//                     <button
//                       onClick={() => this.transfer(ele.diary_id, ele.report_count)}
//                     >
//                       转发
//                     </button>
//                   </div>
//                 )}

//                 {ele.isShowComment && (
//                   <div>
//                     <input
//                       style={{ width: "500px" }}
//                       type="text"
//                       onChange={(e) => this.intChange(e, ele.diary_id)}
//                       value={ele.inputValue}
//                     ></input>
//                     <pre></pre>
//                     <button onClick={() => this.btnSure(ele.diary_id, ele.user_name)}>
//                       评论
//                     </button>
//                   </div>
//                 )}
//                 {ele.isShowComment &&
//                   ele.commentList.map((subEle, subIndex) => {
//                     return (
//                       <div key={subIndex}>
//                         {subEle.user_name} : {subEle.comment}
//                       </div>
//                     );
//                   })}
//               </div>
//             </div>
//           );
//         })}
//       </div>
//     );
//   }

}

export default WorldDiary;
