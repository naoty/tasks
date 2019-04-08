import axios from "axios";
import Vuex from "vuex";

export default function () {
  return new Vuex.Store({
    state: () => ({
      statuses: {
      },
      tasks: {
      },
    }),
    mutations: {
      addTask(state, { taskId, title, statusId }) {
        state.statuses[statusId].tasks.push(taskId);
        state.tasks[taskId] = {
          id: taskId,
          title,
          status: statusId
        };
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
      async addTask({ commit }, task) {
        const { data } = await axios.post(`${process.env.BACKEND_BASE_URL}/tasks`, { task });
        commit("addTask", data.task);
      }
    }
  });
}
