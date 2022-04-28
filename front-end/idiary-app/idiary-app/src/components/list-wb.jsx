import React,{ Component, PropTypes } from 'react';
//import OneWB from './one-wb.jsx';//标签的名字根据这个变量名来决定
import WeiBoListItem from './ListItem';
// class ListWB extends React.Component{
//   render() {
//     // 遍历后端给的数据，并且插入
//     var oneWBNodes = this.props.data.map(function(aWB,index){
//       return <OneWB oneData={aWB} key={index}></OneWB>;
//     });
//     return <div>
//         {oneWBNodes}
//     	</div>;
//   }
// }
export default class WeiBoList extends Component {
    constructor(props) {
      super(props);
    }
  
    render() {
      //遍历微博列表
      var ItemView = this.props.data.map(function(item,index) {
          return <WeiBoListItem itemData={item} key={index}/>

    }) ;
      //渲染列表
      return (
      <div className="listRootViewStyle">
          {ItemView}
      </div>);
    }
  }
  
  WeiBoList.propTypes = {
  };

//export default ListWB;