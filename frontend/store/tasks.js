export const state = () => ({
  tasks: {}
});

export const mutations = {
  add(state, { id, title, status }) {
    state.tasks[id] = { id, title, status };
  },
  set(state, tasks) {
    state.tasks = tasks || {};
  }
};
