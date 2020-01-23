<template>
  <v-container fluid fill-height class="todist-view">
    <v-layout justify-right align-left column pa-5>
      <v-card max-width="600" class="mx-auto">
        <v-toolbar color="red" dark>
          <v-btn icon>
            <v-icon>mdi-check-underline-circle</v-icon>
          </v-btn>

          <v-toolbar-title>{{title}}</v-toolbar-title>
          <v-spacer></v-spacer>
        </v-toolbar>

        <v-list>
          <template v-for="item in items">
            <v-list-item :key="item">
              <template>
                <v-btn icon>
                  <v-icon>mdi-checkbox-blank-circle-outline</v-icon>
                </v-btn>

                <v-list-item-content>
                  <v-list-item-title>{{item}}</v-list-item-title>
                </v-list-item-content>
              </template>
            </v-list-item>
          </template>
        </v-list>
      </v-card>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: "TodoistView",
  data: () => ({
    title: "Course Title",
    items: ["Video 1", "Video 2", "Video 3"],
    count: 3
  })
};
</script>

<style scoped>
.todoist-view {
  background-size: cover;
  background-color: red;
  width: 100%;
  height: 100%;
}
</style>