<template>
  <v-app>
    <v-content>
      {{info}}
      <TodoistView />
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
    info: null
  }),

  mounted() {
    axios
      .get("http://localhost:8181/course/inferential-statistics-intro")
      .then(response => {
        this.info = response;
      });
  }
};
</script>
