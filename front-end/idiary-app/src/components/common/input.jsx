import React from "react";

const Input = ({ label, errors, id, ...rest }) => {
  return (
    <div className="form-group">
      <label htmlFor={id}>
        <i className={label}></i>
      </label>
      <input {...rest} id={id} />
    </div>
  );
};

export default Input;
