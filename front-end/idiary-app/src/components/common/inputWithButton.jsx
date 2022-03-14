import React from "react";

const InputWithButton = ({ text, errors, label, id, ...rest }) => {
  return (
    <div className="mb-3 row">
      <div className="form-floating col">
        <input
          {...rest}
          className={errors[id] ? "form-control is-invalid" : "form-control"}
          id={id}
          placeholder={label}
        />
        <label htmlFor={id}>{label}</label>
      </div>
      <div className="col">
        <button type="submit" class="btn btn-primary mb-3">
          {text}
        </button>
      </div>
    </div>
  );
};

export default InputWithButton;
