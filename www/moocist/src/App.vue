<template>
  <v-app>
    <v-content>
      <v-container fluid>
        <v-row align="center">
          <v-col>
            <input v-model="message" placeholder="edit me" />
            <p>Message is: {{ message }}</p>
            <button v-on:click="getData">Get course info</button>
          </v-col>
        </v-row>
        <TodoistView :course="info" />
      </v-container>
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
    info: {},
    message: ""
  }),

  mounted() {
    this.getData();
  },

  methods: {
    getUrl: function(slug) {
      return "http://localhost:8181/course/" + slug;
    },
    getData: function() {
      axios.get(this.getUrl(this.message)).then(response => {
        this.info = response.data;
      });
    }
  }
};
</script>
