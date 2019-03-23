<template>
  <v-card>
    <v-card-title @mouseenter="isHovered = true" @mouseleave="isHovered = false">
      <h2>{{ task.title }}</h2>
      <v-spacer></v-spacer>
      <v-btn class="ma-0" icon v-show="isHovered" @click="handleClearButtonClick">
        <v-icon>clear</v-icon>
      </v-btn>
    </v-card-title>
  </v-card>
</template>

<script>
export default {
  data: () => ({
    isHovered: false
  }),
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
  },
  methods: {
    handleClearButtonClick(event) {
      this.$store.commit("removeTask", { taskId: this.taskId });
    }
  }
};
</script>
