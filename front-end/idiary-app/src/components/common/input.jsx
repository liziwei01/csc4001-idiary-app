import React from "react";

const Input = ({ label, errors, id, ...rest }) => {
  return (
    <div className="mb-3">
      <div className="form-group mb-0">
        <label htmlFor={id}>
          <i className={label}></i>
        </label>
        <input {...rest} id={id} />
      </div>
      <div>
        {errors && (
          <p>
            <font color="#e52b50">{errors}</font>
          </p>
        )}
      </div>
    </div>
  );
};

export default Input;
