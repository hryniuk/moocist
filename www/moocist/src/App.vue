<template>
  <v-app>
    <v-content>
      <TodoistView :course="info" />
    </v-content>
  </v-app>
</template>

<script>
import axios from "axios";

import TodoistView from "./components/TodoistView";

export default {
  name: "App",

  components: {
    TodoistView
  },

  data: () => ({
    info: {}
  }),

  mounted() {
    axios
      .get("http://localhost:8181/course/inferential-statistics-intro")
      .then(response => {
        this.info = response.data;
      });
  }
};
</script>
