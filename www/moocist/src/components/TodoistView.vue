<template>
  <v-container fluid fill-height class="todist-view">
    <v-layout justify-right align-left column pa-5>
      <v-card max-width="600" class="mx-auto">
        <v-toolbar color="red" dark>
          <v-app-bar-nav-icon></v-app-bar-nav-icon>
          <v-toolbar-title>Todoist Preview</v-toolbar-title>
          <v-spacer></v-spacer>
        </v-toolbar>

        <v-list two-line>
          <v-list-item-group>
            <template v-for="item in items">
              <v-list-item :key="item.title">
                <template>
                  <v-btn icon>
                    <v-icon>mdi-checkbox-blank-circle-outline</v-icon>
                  </v-btn>

                  <v-list-item-content>
                    <v-list-item-title v-text="item"></v-list-item-title>
                    <!-- <v-list-item-subtitle class="text--primary" v-text="item.headline"></v-list-item-subtitle>
                    <v-list-item-subtitle v-text="item.subtitle"></v-list-item-subtitle> -->
                  </v-list-item-content>
                </template>
              </v-list-item>

              <!-- <v-divider v-if="index + 1 < items.length" :key="index"></v-divider> -->
            </template>
          </v-list-item-group>
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