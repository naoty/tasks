<template>
  <v-app>
    <v-toolbar>
      <v-toolbar-side-icon></v-toolbar-side-icon>
      <v-toolbar-title>Tasks</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-dialog>
        <template v-slot:activator="{ on }">
          <v-btn depressed v-on="on">add</v-btn>
        </template>

        <h1>Hello, v-dialog!</h1>
      </v-dialog>
    </v-toolbar>
    <v-content>
      <nuxt/>
    </v-content>
  </v-app>
</template>
