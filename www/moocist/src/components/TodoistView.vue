<template>
  <v-container fluid fill-height class="todist-view">
    <v-layout justify-right align-left column pa-5>
      <v-card max-width="600" class="mx-auto">
        <v-toolbar color="red" dark>
          <v-btn icon>
            <v-icon>mdi-check-underline-circle</v-icon>
          </v-btn>

          <v-spacer></v-spacer>
          <v-toolbar-title>{{title}}</v-toolbar-title>
        </v-toolbar>

        <v-list>
          <v-list-group
            v-for="week in course.weeks"
            :key="week.title"
            v-model="week.active"
            no-action
          >
            <template v-slot:activator>
              <v-list-item-content>
                <v-list-item-title v-text="week.title"></v-list-item-title>
              </v-list-item-content>
            </template>

            <v-list-item v-for="(task, i) in week.tasks" :key="i">
              <v-list-item-content>
                <v-list-item-title v-text="task.title"></v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list-group>
        </v-list>
      </v-card>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: "TodoistView",
  props: {
    course: {
      type: Object,
      default: () => ({
        weeks: []
      })
    }
  },
  data: () => ({
    title: "Course"
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