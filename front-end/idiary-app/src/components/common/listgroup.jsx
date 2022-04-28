import React from "react";

const ListGroup = ({
  items,
  selectedItem,
  onItemSelect,
}) => {
  return (
    <ul className="list-group" style={{ borderRadius: "20px", marginTop: "20px" }}>
      {items.map((item) => (
        <li
          onClick={() => onItemSelect(item)}
          key={item._id}
          className={
            item === selectedItem ? "list-group-item active clickable" : "list-group-item clickable"
          }
        >
          <i class={item.icon} aria-hidden="true"></i>{item.name}
        </li>
      ))}
    </ul>
  );
};

ListGroup.defaultProps = {
  textProperty: "name",
  valueProperty: "_id",
};

export default ListGroup;
