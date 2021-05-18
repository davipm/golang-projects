import React, { useState } from "react";
import { Card, Header, Form, Input } from "semantic-ui-react";
import { useQuery, useQueryClient, useMutation } from "react-query";

import ItemCard from "../item-card";
import api from "../../api";

export default function ItemsList() {
  const queryClient = useQueryClient();
  const [task, setTask] = useState("");

  const { data } = useQuery("todos", async () => {
    const { data } = await api.get(`/api/task`);
    return data;
  });

  const mutation = useMutation((task) => api.post(`/api/task`, { task }), {
    onSuccess: () => {
      queryClient.invalidateQueries("todos");
      setTask("");
    },
  });

  const deleteMutation = useMutation(
    (id) => api.delete(`/api/deleteTask/${id}`),
    {
      onSuccess: () => queryClient.invalidateQueries("todos"),
    }
  );

  const updateMutation = useMutation((id) => api.put(`/api/task/${id}`), {
    onMutate: async () => {
      const { data: previousValue } = queryClient.getQueryData("todos");

      return previousValue;
    },

    onSuccess: (_, variables) => {
      queryClient.setQueryData("todos", (old) =>
        old.map((item) => {
          if (item._id === variables && !item.status) {
            return {
              ...item,
              status: !item.status,
            };
          }
          return item;
        })
      );
    },

    onError: (error) => {
      console.log(error);
    },
  });

  const undoMutation = useMutation((id) => api.put(`/api/undoTask/${id}`), {
    onMutate: (variables) => {
      queryClient.setQueryData("todos", (old) =>
        old.map((item) => {
          if (item._id === variables && item.status) {
            return {
              ...item,
              status: !item.status,
            };
          }
          return item;
        })
      );
    },
  });

  async function onSubmit() {
    if (!task) return;
    mutation.mutate(task);
  }

  return (
    <>
      <div className="row">
        <Header className="header" as="h2">
          TO DO LIST
        </Header>
      </div>

      <div className="row">
        <Form onSubmit={onSubmit}>
          <Input
            input="text"
            name="task"
            placeholder="Create Task"
            onChange={(event) => setTask(event.target.value)}
            value={task}
            fluid
          />
        </Form>
      </div>

      <div className="row">
        <Card.Group>
          {data?.map((item) => (
            <ItemCard
              key={item._id}
              item={item}
              onDelete={deleteMutation}
              onUndo={undoMutation}
              onUpdate={updateMutation}
            />
          ))}
        </Card.Group>
      </div>
    </>
  );
}
