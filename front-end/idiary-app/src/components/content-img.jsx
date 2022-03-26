import React from 'react';
class ContentImg extends React.Component{
    constructor(props) {
        super(props);
      }
  render() {
    var imgNodes=this.props['content-img-urls'].map(function(oneImg,index){
      return <div key={index} style={{display:"flex",marginLeft:"20px",flexDirection:"row"}}><img style={{width:"100px",height:"100px"}} src={oneImg} alt="微博配图" /></div>;
    });
    return <div  className="row_extra-content_clearfix">
      <ul style={{display:"flex",flexDirection:"row"}}>
        {imgNodes}
      </ul>
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