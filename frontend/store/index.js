import Vuex from "vuex";

export default function () {
  return new Vuex.Store({
    state: () => ({
      statuses: {
        "1": { id: "1", name: "TODO", tasks: [] },
        "2": { id: "2", name: "DOING", tasks: [] },
        "3": { id: "3", name: "DONE", tasks: [] }
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
        state.statuses[state.initialTaskStatusId].tasks.push(state.nextTaskId);
        state.nextTaskId = String(Number(state.nextTaskId) + 1);
      },
      moveTask(state, { oldStatusId, oldIndex, newStatusId, newIndex }) {
        const oldStatus = state.statuses[oldStatusId];
        const taskId = oldStatus.tasks[oldIndex];
        oldStatus.tasks.splice(oldIndex, 1);

        const task = state.tasks[taskId];
        task.status = newStatusId;

        const newStatus = state.statuses[newStatusId];
        newStatus.tasks.splice(newIndex, 0, taskId);
      },
      removeTask(state, { taskId }) {
        const task = state.tasks[taskId];
        const status = state.statuses[task.status];
        status.tasks.splice(status.tasks.indexOf(taskId), 1);

        delete state.tasks[taskId];
      }
    }
  });
}
