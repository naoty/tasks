<template>
  <v-dialog v-model="canShowDialog" width="360">
    <template v-slot:activator="{ on }">
      <slot name="activator" :on="on"></slot>
    </template>

    <v-card class="pa-4">
      <h1 class="mb-2">New task</h1>
      <v-form ref="form">
        <v-text-field v-model="title" :rules="titleRules" label="Title" required></v-text-field>
        <v-btn class="mt-2 ml-0" color="primary" @click="save">Save</v-btn>
      </v-form>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  data: () => ({
    canShowDialog: false,
    title: "",
    titleRules: [value => !!value || "Title is required"]
  }),
  methods: {
    async save() {
      const isValid = this.$refs.form.validate();
      if (!isValid) {
        return;
      }

      await this.$store.dispatch("addTask", this.title);
      this.canShowDialog = false;
    }
  }
};
</script>

