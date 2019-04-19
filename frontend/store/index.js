import axios from "axios";
import Vuex from "vuex";
import gql from "graphql-tag";
import { normalize, schema } from "normalizr";
import { statusesSchema } from "../schema";

export default function () {
  return new Vuex.Store({
    state: () => ({
      statuses: {
      },
      tasks: {
      },
    }),
    mutations: {
      addTask(state, { id, title, status }) {
        state.statuses[status].tasks.push(id);
        state.tasks[id] = { id, title, status };
      },
      moveTask(state, { oldStatusId, oldIndex, newStatusId, newIndex }) {
        const oldStatus = state.statuses[oldStatusId];
        const taskId = oldStatus.tasks[oldIndex];
        oldStatus.tasks.splice(oldIndex, 1);

        const task = state.tasks[taskId];
        task.statusId = newStatusId;

        const newStatus = state.statuses[newStatusId];
        newStatus.tasks.splice(newIndex, 0, taskId);
      },
      removeTask(state, { taskId }) {
        const task = state.tasks[taskId];
        const status = state.statuses[task.statusId];
        status.tasks.splice(status.tasks.indexOf(taskId), 1);

        delete state.tasks[taskId];
      },
      setStatusesAndTasks(state, { statuses, tasks }) {
        state.statuses = statuses || {};
        state.tasks = tasks || {};
      }
    },
    actions: {
      async addTask({ commit }, title) {
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
        commit("addTask", task);
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
        commit("setStatusesAndTasks", normalizedData.entities);
      }
    }
  });
}
