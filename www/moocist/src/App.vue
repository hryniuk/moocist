<template>
  <v-app>
    <v-content>
      <TodoistView />
      {{info}}
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
      .get("localhost:)
      .then(response => {
        this.info = "XD" + response;
      });
  }
};
</script>
