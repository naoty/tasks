<template>
  <v-container fluid>
    <v-layout>
      <v-flex v-for="status in $store.state.statuses" :key="status.id" pa-2>
        <h1>{{ status.name }}</h1>
        <v-divider></v-divider>
        <v-layout column mt-2>
          <v-flex
            v-for="task in $store.state.tasks.filter(task => task.status === status.id)"
            :key="task.id"
            mb-2
          >
            <v-card>
              <v-card-title>
                <h2>{{ task.title }}</h2>
              </v-card-title>
            </v-card>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
</template>
