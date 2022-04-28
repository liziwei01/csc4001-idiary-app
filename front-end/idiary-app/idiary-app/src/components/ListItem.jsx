import React, { Component, PropTypes } from 'react';
import ContentImg from './content-img.jsx';

/**
 * 微博评论列表组件
 */
export default class WeiBoListItem extends Component {
  constructor(props) {
    super(props);
    this.state = {
      //是否点击评论按钮标志
      isComment:false,
      //默认的条目数据
      itemData:this.props.itemData,
      //默认的点赞数
      zanNum:this.props.itemData.NoCollect
    }
  }

  render() {
    //取得传递过来的数据
    let data = this.props.itemData ;
    return (
      <div>
          {this._renderHeadView(data)}
          <hr className="hrStyle"/>
          {this._renderFooterView(data)} 
        </div>
     );
  }

  /**
   * 渲染顶部View
   */
  _renderHeadView(data){
    var contentImgs;
    if(data.contentImgUrls){
        //若后端给的数据中有图片url，则展示
        contentImgs = <ContentImg content-img-urls={data.contentImgUrls} />
      }
    return(
      <div className="item">
        <div className="topRightView">
          <div className="nickNameAndSendTime">
            <div>
              <img className="nick-img" src={data.headUrl}/>
              <span style={{marginLeft:'10px'}}>{data.nickName}</span>
            </div>
            
            <span>{data.sendTime}</span>
          </div>
          <div>
            <p style={{marginTop:'10px'}}>{data.content}</p>
            <div>{contentImgs}</div>
          </div>
          

        </div>

        
      </div>
    )
  }

  /**
   * 渲染底部View
   */
    like = ()=>{
      
    }
    comment = ()=>{
      //新的框，里面是现有评论
    }
    transferdata = ()=>{
    
    }
   _renderFooterView(data){
       return(
         <div className="commentViewStyle">
           <ul className="ulStyle">
             <li className="liStyle" onClick={this.like}>点赞:{this.state.zanNum}</li><div className="shuxian"></div>
             <li className="liStyle" onClick={this.comment}>评论:{data.NoComment}</li><div className="shuxian"></div>
             <li className="liStyle" onClick={this.transferdata}>转发:{data.NoPointGreat}</li>
           </ul>
         </div>
       );
   }
 }

