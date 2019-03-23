<template>
  <v-app>
    <v-toolbar>
      <v-toolbar-side-icon></v-toolbar-side-icon>
      <v-toolbar-title>Tasks</v-toolbar-title>
      <v-spacer></v-spacer>
      <task-form>
        <template v-slot:activator="{ on }">
          <v-btn depressed v-on="on">add</v-btn>
        </template>
      </task-form>
    </v-toolbar>
    <v-content>
      <nuxt/>
    </v-content>
  </v-app>
</template>

<script>
import TaskForm from "../components/TaskForm";

export default {
  components: {
    "task-form": TaskForm
  }
};
</script>

