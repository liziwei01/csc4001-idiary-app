import React from 'react';
import CommentForm from './comment-form.jsx';
import ContentImg from './content-img.jsx';
class OneWB extends React.Component{
  constructor(props){
    super(props);
    this.state={
      isComment:false,
    };
  }
  render() {
   var oneData=this.props.oneData;
   var commentForm;
   var contentImgs;
   if(this.state.isComment) {
     //控制评论框是否展现,因为是动态的，所以放在state而不是props
     commentForm =<CommentForm data-my-head-img={oneData.headUrl}/>;
   }
   if(oneData.contentImgUrls){
     //若后端给的数据中有图片url，则展示
     contentImgs = <ContentImg content_img_urls={oneData.contentImgUrls} />
   }
   return <div className="big-center" >
     <div className="one-wb ">
       <div className="clearfix">
         <div className="ow-left">
            <img className="nick-img" src={oneData.headUrl}/>
         </div>
         <div className="ow-right">
           <div className="ow-nick row">
             <span>{oneData.nickName}</span>
           </div>
           <div className="ow-content_row">{oneData.content}</div>
           {contentImgs}
         </div>
       </div>
       <div className="ow-footer row ">
         <ul className="clearfix" onClick={this.handlerForwardClick.bind(this)}>
           <li className="li-side-border"><button>收藏</button> {oneData.NoCollect}</li>
           <li className="li-side-border"><button>转发</button> {oneData.NoForward}</li>
           <li className="li-side-border"><button>评论</button> {oneData.NoComment}</li>
           <li><span>赞</span> {oneData.NoPointGreat}</li>
         </ul>
       </div>
     </div>
     {commentForm}
   </div>;
  }
  handlerForwardClick(event) {
    if(event.target.innerText == '评论'){
      this.setState({isComment:true});
    }else{
      this.setState({isComment:false});
    }
  }
}
export default OneWB;
