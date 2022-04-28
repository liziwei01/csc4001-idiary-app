import React from "react";

const Button = ({ label, ...rest }) => {
  return (
    <button {...rest} type="submit" className="w-100 btn btn-lg btn-primary">
      {label}
    </button>
  );
};

export default Button;
