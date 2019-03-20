import Vuex from "vuex";

export default function () {
  return new Vuex.Store({
    state: () => ({
      statuses: [
        { id: "1", name: "TODO" },
        { id: "2", name: "DOING" },
        { id: "3", name: "DONE" }
      ],
      tasks: [
        { id: "1", title: "Task 1", status: "1" },
        { id: "2", title: "Task 2", status: "1" },
        { id: "3", title: "Task 3", status: "2" },
        { id: "4", title: "Task 4", status: "3" }
      ]
    })
  });
}
