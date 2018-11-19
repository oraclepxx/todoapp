<template>
  <div class="todomain">
    <h1>Hello, xYZ ToDo List</h1>
    <todo-creator />
    <todo-list />
  </div>
</template>

<script>
import TodoCreator from './todocreator.vue';
import TodoList from './todolist.vue';

export default {
  name: 'TodoMain',
  components: {
    TodoCreator,
    TodoList
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
