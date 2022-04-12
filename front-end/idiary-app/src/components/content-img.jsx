import React from 'react';
import "../css/image.css"
class ContentImg extends React.Component{
    constructor(props) {
        super(props);
      }
  render() {
    var imgNodes=this.props['content-img-urls'].map(function(oneImg,index){
      return <div key={index} className="aa">
            <img style={{width:"240px",height:"240px"}} src={oneImg} alt="微博配图" />
          </div>;
    });
    return <div  className="box">
      
        {imgNodes}
      
    </div>
    // var arr=[]
    // var badge = this.props.content_img_urls;
    // if(badge){
    //     for (var i=0; i<badge.length; i++){
    //       arr.push(
    //           <div className="outerViewStyle">
    //               <div className="viewStyle">
    //                   <img source={badge[i]}/> 
    //               </div>
    //           </div>
    //       )
    //     }
    //     return arr;
    // }
    //return <div>test</div>
  }
}
export default ContentImg;