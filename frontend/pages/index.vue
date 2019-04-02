<template>
  <v-container fluid>
    <v-layout>
      <v-flex v-for="(status, statusId) in $store.state.statuses" :key="statusId" pa-2>
        <h1>{{ status.name }}</h1>
        <v-divider></v-divider>
        <v-layout column mt-2>
          <draggable group="status" @end="handleDragEnd" :data-status-id="statusId">
            <v-flex v-for="taskId in status.taskIds" :key="taskId" mb-2>
              <task-card :taskId="taskId"></task-card>
            </v-flex>
          </draggable>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import axios from "axios";
import draggable from "vuedraggable";
import TaskCard from "../components/TaskCard";

export default {
  components: {
    draggable,
    "task-card": TaskCard
  },
  async fetch({ env, store }) {
    const host = process.client ? env.clientBackendHost : env.serverBackendHost;
    const url = `http://${host}:${env.backendPort}/statuses`;
    const { data } = await axios.get(url);
    const statuses = data.reduce((result, status) => {
      result[status.statusId] = status;
      return result;
    }, {});
    store.commit("setStatuses", { statuses });
  },
  methods: {
    handleDragEnd: function(event) {
      this.$store.commit("moveTask", {
        oldStatusId: event.from.dataset.statusId,
        oldIndex: event.oldIndex,
        newStatusId: event.to.dataset.statusId,
        newIndex: event.newIndex
      });
    }
  }
};
</script>
