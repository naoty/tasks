<template>
  <v-container fluid>
    <v-layout>
      <v-flex v-for="(status, statusId) in $store.state.statuses" :key="statusId" pa-2>
        <h1>{{ status.name }}</h1>
        <v-divider></v-divider>
        <v-layout column mt-2>
          <draggable group="status" @end="handleDragEnd" :data-status-id="statusId">
            <v-flex v-for="taskId in status.tasks" :key="taskId" mb-2>
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
import { normalize } from "normalizr";
import TaskCard from "../components/TaskCard";
import { statusSchema } from "../schema";

export default {
  components: {
    draggable,
    "task-card": TaskCard
  },
  async fetch({ env, store }) {
    const url = `${process.env.BACKEND_BASE_URL}/statuses`;
    const { data } = await axios.get(url);
    const normalizedData = normalize(data, [statusSchema]);
    store.commit("setStatusesAndTasks", normalizedData.entities);
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
