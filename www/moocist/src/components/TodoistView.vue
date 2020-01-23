<template>
    <v-container fluid fill-height class="todist-view">
        <v-layout justify-right align-left column pa-5>
            <v-list-item-group>
            <v-list-item
              v-for="item in items"
              :key="item"
            >
              <v-list-item-content>
                <v-list-item-title v-text="item"></v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            </v-list-item-group>
        </v-layout>
    </v-container>
</template>

<script>
export default {
    name: 'TodoistView',
    data: () => ({
        items: ['a', 'b', 'c'],
        count: 3,
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