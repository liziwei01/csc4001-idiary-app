import React from "react";

const Input = ({ label, errors, id, ...rest }) => {
  return (
    <div className="mb-3">
      <label htmlFor={id} className="form-label">
        {label}
      </label>
      <input
        {...rest}
        className={errors[id] ? "form-control is-invalid" : "form-control"}
        id={id}
      />
    </div>
  );
};

export default Input;
