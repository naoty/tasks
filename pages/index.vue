<template>
  <v-container fluid>
    <v-layout>
      <v-flex v-for="(status, statusId) in $store.state.statuses" :key="statusId" pa-2>
        <h1>{{ status.name }}</h1>
        <v-divider></v-divider>
        <v-layout column mt-2>
          <draggable group="status.id">
            <v-flex v-for="(task, taskId) in tasksWithStatus(status)" :key="taskId" mb-2>
              <v-card>
                <v-card-title>
                  <h2>{{ task.title }}</h2>
                </v-card-title>
              </v-card>
            </v-flex>
          </draggable>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import draggable from "vuedraggable";

export default {
  components: {
    draggable
  },
  methods: {
    tasksWithStatus: function(status) {
      return status.tasks.map(taskId => this.$store.state.tasks[taskId]);
    }
  }
};
</script>
