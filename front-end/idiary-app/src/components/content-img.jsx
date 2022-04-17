import React from "react";
import "../css/image.css";
class ContentImg extends React.Component {
  constructor(props) {
    super(props);
  }
  render() {
    //console.log(this.props["contentImgUrls"]);
    var imgNodes = this.props["contentImgUrls"].map(function (oneImg, index) {
      return (
        <div key={index} className="aa">
          <img
            style={{ width: "240px", height: "240px" }}
            src={oneImg}
            alt="微博配图"
          />
        </div>
      );
    });
    return <div className="box">{imgNodes}</div>;
  }
}
export default ContentImg;
