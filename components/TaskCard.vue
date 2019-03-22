<template>
  <v-card>
    <v-card-title>
      <h2>{{ task.title }}</h2>
    </v-card-title>
  </v-card>
</template>

<script>
export default {
  props: {
    taskId: {
      type: String,
      required: true
    }
  },
  computed: {
    task: function() {
      return this.$store.state.tasks[this.taskId];
    }
  }
};
</script>

