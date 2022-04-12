import React from "react";

const InputWithButton = ({ onClick, text, errors, label, id, ...rest }) => {
  return (
    <div>
      <div className="form-group mb-0">
        <input {...rest} id={id} />
        <label htmlFor={id}>
          <i className={label}></i>
        </label>
      </div>
      <div>
        {errors && (
          <p>
            <font color="#e52b50">{errors}</font>
          </p>
        )}
      </div>
      <div className="mt-3">
        <button
          type="submit"
          className="w-100 btn btn-secondary mb-3"
          onClick={onClick}
        >
          {text}
        </button>
      </div>
    </div>
  );
};

export default InputWithButton;
