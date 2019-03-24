import Vuex from "vuex";

export default function () {
  return new Vuex.Store({
    state: () => ({
      statuses: {
      },
      tasks: {
      },
      nextTaskId: "1",
      initialTaskStatusId: "1"
    }),
    mutations: {
      addTask(state, { title }) {
        state.tasks[state.nextTaskId] = {
          id: state.nextTaskId,
          title,
          status: state.initialTaskStatusId
        };
        state.statuses[state.initialTaskStatusId].taskIds.push(state.nextTaskId);
        state.nextTaskId = String(Number(state.nextTaskId) + 1);
      },
      moveTask(state, { oldStatusId, oldIndex, newStatusId, newIndex }) {
        const oldStatus = state.statuses[oldStatusId];
        const taskId = oldStatus.taskIds[oldIndex];
        oldStatus.taskIds.splice(oldIndex, 1);

        const task = state.tasks[taskId];
        task.status = newStatusId;

        const newStatus = state.statuses[newStatusId];
        newStatus.taskIds.splice(newIndex, 0, taskId);
      },
      removeTask(state, { taskId }) {
        const task = state.tasks[taskId];
        const status = state.statuses[task.status];
        status.taskIds.splice(status.taskIds.indexOf(taskId), 1);

        delete state.tasks[taskId];
      },
      setStatuses(state, { statuses }) {
        state.statuses = statuses;
      }
    }
  });
}
