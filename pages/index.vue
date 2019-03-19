<template>
  <v-container fluid>
    <v-layout>
      <v-flex pa-2>
        <h1>TODO</h1>
        <v-divider></v-divider>
        <v-layout column mt-2>
          <v-flex mb-2>
            <v-card>
              <v-card-title>
                <h2>Task 1</h2>
              </v-card-title>
            </v-card>
          </v-flex>
          <v-flex mb-2>
            <v-card>
              <v-card-title>
                <h2>Task 2</h2>
              </v-card-title>
            </v-card>
          </v-flex>
        </v-layout>
      </v-flex>
      <v-flex pa-2>
        <h1>DOING</h1>
        <v-divider></v-divider>
        <v-layout column mt-2>
          <v-flex mb-2>
            <v-card>
              <v-card-title>
                <h2>Task 3</h2>
              </v-card-title>
            </v-card>
          </v-flex>
        </v-layout>
      </v-flex>
      <v-flex pa-2>
        <h1>DONE</h1>
        <v-divider></v-divider>
        <v-layout column mt-2>
          <v-flex mb-2>
            <v-card>
              <v-card-title>
                <h2>Task 4</h2>
              </v-card-title>
            </v-card>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
</template>
