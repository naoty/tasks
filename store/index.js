import Vuex from "vuex";

export default function () {
  return new Vuex.Store({
    state: () => ({
      statuses: {
        "1": { id: "1", name: "TODO", tasks: ["1", "2"] },
        "2": { id: "2", name: "DOING", tasks: ["3"] },
        "3": { id: "3", name: "DONE", tasks: ["4"] }
      },
      tasks: {
        "1": { id: "1", title: "Task 1" },
        "2": { id: "2", title: "Task 2" },
        "3": { id: "3", title: "Task 3" },
        "4": { id: "4", title: "Task 4" }
      }
    })
  });
}
