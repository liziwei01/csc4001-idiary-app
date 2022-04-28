import React, { Component } from "react";
import ReactDOM from 'react-dom';
//import ListWB from './list-wb.jsx';//微博列表组件
import WeiBoList from './list-wb.jsx';
import onePic from './1.jpg'
import twoPic from './2.jpg'
import threePic from './3.jpg'

// var dataList = [
//     {
//         headUrl: onePic,
//         nickName: 'Robin',
//         content: '拿快递拿快递3号小邮局爆仓啦',
//         NoCollect: 132,
//         NoForward: 202,
//         NoComment: 142,
//         NoPointGreat: 423,
//         contentImgUrls: [
//             twoPic,
//             threePic
//         ]
//     },
    
// ];

var dataList = [
    {
        myUrl: onePic,
        headUrl: onePic,
        nickName: 'Robin',
        content: 'test1111111111111111111111111111111111111111111111111111111111111111111111111',
        NoCollect: 132,
        NoForward: 202,
        NoComment: 142,
        NoPointGreat: 423,
        sendTime:"11月2日",
        contentImgUrls: [
            twoPic,
            threePic
        ]
    },
    {
        myUrl: onePic,
        headUrl: onePic,
        nickName: 'Robin',
        content: 'test22222222222222222222222222222222222',
        NoCollect: 132,
        NoForward: 202,
        NoComment: 142,
        NoPointGreat: 423,
        sendTime:"11月2日",
        contentImgUrls: [
            twoPic,
            threePic
        ]
    },
    
];

// class personalDiary extends React.Component {
//     render() {
//         return <div>
//             <ListWB data={dataList} />
//         </div>;
//     }
    
// }
class personalDiary extends React.Component {
    render() {
        return <div>
            <WeiBoList data={dataList} />
        </div>;
    }
    
}

export default personalDiary;