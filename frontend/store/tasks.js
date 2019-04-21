export const state = () => ({
  tasks: {}
});

export const mutations = {
  add(state, { id, title, status }) {
    state.tasks[id] = { id, title, status };
  },
  remove(state, id) {
    delete state.tasks[id];
  },
  set(state, tasks) {
    state.tasks = tasks || {};
  }
};
