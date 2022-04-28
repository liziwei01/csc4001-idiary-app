import React from "react";

const InputWithDesc = ({ text, errors, label, id, ...rest }) => {
  let help = { id };
  help += "HelpInline";
  return (
    <div className="mb-3">
      <div className="form-group mb-0">
        <input {...rest} id={id} aria-describedby={help} />
        <label htmlFor={id}>
          <i className={label}></i>
        </label>
      </div>
      <div>
        <span id={help} className="form-text">
          {text}
        </span>
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

export default InputWithDesc;
