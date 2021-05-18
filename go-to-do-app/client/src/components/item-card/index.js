import React, { memo } from "react";
import { Card, Icon } from "semantic-ui-react";

function ItemCard({ item, onDelete, onUpdate, onUndo }) {
  return (
    <Card key={item._id} color={item.status ? "green" : "yellow"} fluid>
      <Card.Content>
        <Card.Header textAlign="left">
          <div className={`title ${item.status ? "active" : ""}`}>
            {item.task}
          </div>
        </Card.Header>

        <Card.Meta textAlign="right">
          <span className="btn" onClick={() => onUpdate.mutate(item._id)}>
            <Icon name="check circle" color="green" />
            <span className="pr-10">Done</span>
          </span>
          <span className="btn" onClick={() => onUndo.mutate(item._id)}>
            <Icon name="undo" color="yellow" />
            <span className="pr-10">Undo</span>
          </span>
          <span className="btn" onClick={() => onDelete.mutate(item._id)}>
            <Icon name="delete" color="red" />
            <span className="pr-10">Delete</span>
          </span>
        </Card.Meta>
      </Card.Content>
    </Card>
  );
}

export default memo(ItemCard);
