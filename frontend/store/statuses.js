export const state = () => ({
  statuses: {}
});

export const mutations = {
  addTask(state, { statusId, taskId }) {
    state.statuses[statusId].tasks.push(taskId);
  },
  removeTask(state, { statusId, taskId }) {
    const status = state.statuses[statusId];
    status.tasks.splice(status.tasks.indexOf(taskId), 1);
  },
  set(state, statuses) {
    state.statuses = statuses || {};
  }
};
