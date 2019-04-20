export const state = () => ({
  statuses: {}
});

export const mutations = {
  addTask(state, { statusId, taskId }) {
    state.statuses[statusId].tasks.push(taskId);
  },
  set(state, statuses) {
    state.statuses = statuses || {};
  }
};
