import { schema } from "normalizr";

export const taskSchema = new schema.Entity(
  "tasks",
  {},
  { idAttribute: "taskId" }
);

export const statusSchema = new schema.Entity(
  "statuses",
  { tasks: [taskSchema] },
  { idAttribute: "statusId" }
);
