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

          <template v-for="week in courses.weeks">
            <v-list-item :key="week">
              <template>
                  <v-list-item-title>{{week.title}}</v-list-item-title>
              </template>
              <!-- <v-list> -->
            <v-list>
              <template v-for="task in week.tasks">
                <v-list-item :key="task">
                  <template>
                    <v-btn icon>
                      <v-icon>mdi-checkbox-blank-circle-outline</v-icon>
                    </v-btn>

                    <v-list-item-content>
                      <v-list-item-title>{{task.title}}</v-list-item-title>
                    </v-list-item-content>
                  </template>
                </v-list-item>
              </template>
            </v-list>
              <!-- </v-list> -->
            </v-list-item>
          </template>
      </v-card>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: "TodoistView",
  props: ["courses"],
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