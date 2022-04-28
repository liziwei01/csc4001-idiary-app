import React from "react";

const Input = ({ label, errors, id, ...rest }) => {
  return (
    <div className="mb-3">
      <label htmlFor={id} className="form-label">
        {label}
      </label>
      <input {...rest} className="form-control" id={id} />
      {errors[id] && <div className="alert alert-danger">{errors[id]}</div>}
    </div>
  );
};

export default Input;
