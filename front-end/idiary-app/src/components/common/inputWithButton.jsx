import React from "react";

const InputWithButton = ({ onClick, text, errors, label, id, ...rest }) => {
  return (
    <div className="row mb-3">
      <div className="form-floating col-md-6">
        <input
          {...rest}
          className={errors[id] ? "form-control is-invalid" : "form-control"}
          id={id}
          placeholder={label}
        />
        <label htmlFor={id}>{label}</label>
      </div>
      <div className="col-md-4">
        <button
          type="submit"
          className="btn btn-primary mb-3"
          onClick={onClick}
        >
          {text}
        </button>
      </div>
    </div>
  );
};

export default InputWithButton;
