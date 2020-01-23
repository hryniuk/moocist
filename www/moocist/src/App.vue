<template>
  <v-app>
    <v-content>
      <TodoistView/>
    </v-content>
  </v-app>
</template>

<script>
import TodoistView from './components/TodoistView';

export default {
  name: 'App',

  components: {
    TodoistView,
  },

  data: () => ({
    //
  }),
};
</script>
