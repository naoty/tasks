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
        state.tasks[state.nextTaskId] = { id: state.nextTaskId, title };
        state.statuses[state.initialTaskStatusId].tasks.push(state.nextTaskId);
        state.nextTaskId = String(Number(state.nextTaskId) + 1);
      },
      moveTask(state, { oldStatusId, oldIndex, newStatusId, newIndex }) {
        const oldStatus = state.statuses[oldStatusId];
        const taskId = oldStatus.tasks[oldIndex];
        oldStatus.tasks.splice(oldIndex, 1);

        const newStatus = state.statuses[newStatusId];
        newStatus.tasks.splice(newIndex, 0, taskId);
      }
    }
  });
}
