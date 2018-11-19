<template>
  <div id="app">
    <todo-main />
  </div>
</template>

<script>
import TodoMain from './components/todomain.vue'

export default {
  name: 'app',
  components: {
    TodoMain
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
