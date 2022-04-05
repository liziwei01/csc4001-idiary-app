import React, { Component } from "react";
import ReactDOM from "react-dom";
import WeiBoList from "./list-wb.jsx";
import onePic from "./1.jpg";
import twoPic from "./2.jpg";
import threePic from "./3.jpg";
import Pagination from "./common/pagination.jsx";
import _ from "lodash";
import { paginate } from "../utils/paginate";

var dataList = [
  {
    myUrl: onePic,
    headUrl: onePic,
    nickName: "Robin",
    content:
      "test1111111111111111111111111111111111111111111111111111111111111111111111111",
    NoCollect: 132,
    NoForward: 202, //点赞
    NoComment: 142, //评论
    NoPointGreat: 423, //转发
    sendTime: "11月2日",
    contentImgUrls: [twoPic, threePic],
    isShowComment: false,
    isTransfer: false,
    id: 1,
    commentList: [
      {
        name: "user1",
        comment: "test123456",
      },
      {
        name: "user2",
        comment: "test123456",
      },
    ],
  },
  {
    myUrl: onePic,
    headUrl: onePic,
    nickName: "Robin",
    content: "test22222222222222222222222222222222222",
    NoCollect: 132,
    NoForward: 202,
    NoComment: 142,
    NoPointGreat: 423,
    sendTime: "11月2日",
    contentImgUrls: [twoPic, threePic],
    isShowComment: false,
    id: 2,
    commentList: [
      {
        name: "xxx",
        age: 18,
      },
    ],
  },
];

class personalDiary extends React.Component {
  state = {
    list: [...dataList],
    currentPage: 1,
    pageSize: 1,
    selectedGenre: null,
    sortColumn: { path: "sendTime", order: "asc" },
  };

  // 页面加载的钩子
  // react 生命周期
  componentDidMount() {
    // xxxx 请求列表的数据
    // this.setState({
    //     list : xxxx
    // })
  }
  getPagedData = () => {
    const { pageSize, currentPage, sortColumn, list: allDiaries } = this.state;

    let filtered = allDiaries;

    const sorted = _.orderBy(filtered, [sortColumn.path], [sortColumn.order]);

    const diaries = paginate(sorted, currentPage, pageSize);

    return { totalCount: filtered.length, data: diaries };
  };
  handlePageChange = (page) => {
    this.setState({ currentPage: page });
  };
  render() {
    const { pageSize, currentPage } = this.state;
    const { totalCount, data: diaries } = this.getPagedData();
    return (
      <div>
        <WeiBoList
          data={diaries}
          update={(list) => {
            this.setState({ list });
          }}
        />
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

export default personalDiary;
