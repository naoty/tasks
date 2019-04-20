import gql from "graphql-tag";
import { normalize, schema } from "normalizr";
import { statusesSchema } from "../schema";

export const state = () => ({});

export const actions = {
  async createTask({ commit }, title) {
    const client = this.app.apolloProvider.defaultClient;
    const { data } = await client.mutate({
      mutation: gql`mutation ($title: String!) {
        createTask(input: {
          title: $title
        }) {
          task {
            id
            title
            status {
              id
            }
          }
        }
      }`,
      variables: {
        title
      }
    });
    const statusSchema = new schema.Entity("statuses");
    const taskSchema = new schema.Entity("tasks", { status: statusSchema });
    const rootSchema = new schema.Object({ createTask: { task: taskSchema } });
    const normalizedData = normalize(data, rootSchema);
    const task = normalizedData.entities.tasks[normalizedData.result.createTask.task];
    commit("tasks/add", task);
    commit("statuses/addTask", { statusId: task.status, taskId: task.id })
  },
  async fetchStatuses({ commit }) {
    const client = this.app.apolloProvider.defaultClient;
    const { data } = await client.query({
      query: gql`query {
        statuses {
          id
          name
          tasks {
            id
            title
          }
        }
      }`
    });
    const rootSchema = new schema.Object({ statuses: statusesSchema });
    const normalizedData = normalize(data, rootSchema);
    commit("tasks/set", normalizedData.entities.tasks);
    commit("statuses/set", normalizedData.entities.statuses);
  }
};
