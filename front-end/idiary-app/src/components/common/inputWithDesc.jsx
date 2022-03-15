import React from "react";

const InputWithDesc = ({ text, errors, label, id, ...rest }) => {
  let help = { id };
  help += "HelpInline";
  return (
    <div className="mb-3 row">
      <div className="form-floating col">
        <input
          {...rest}
          className={errors[id] ? "form-control is-invalid" : "form-control"}
          id={id}
          aria-describedby={help}
          placeholder={label}
        />
        <label htmlFor={id}>{label}</label>
      </div>
      <div className="col">
        <span id={help} className="form-text">
          {text}
        </span>
      </div>
    </div>
  );
};

export default InputWithDesc;
