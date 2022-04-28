import React from "react";

const Button = ({ label, ...rest }) => {
  return (
    <button {...rest} type="submit" className="btn btn-primary">
      {label}
    </button>
  );
};

export default Button;
