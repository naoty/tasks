import { schema } from "normalizr";

export const taskSchema = new schema.Entity("tasks");
export const tasksSchema = new schema.Array(taskSchema);
export const statusSchema = new schema.Entity("statuses", { tasks: tasksSchema });
export const statusesSchema = new schema.Array(statusSchema);
