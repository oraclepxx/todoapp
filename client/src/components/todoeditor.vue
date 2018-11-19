<template>
  <div class="todoeditor">
    <h3>Todo Editor</h3>
    <form>
        <div align="left"><label>Summary:</label><textarea id="summary" name="summary" size="50" v-model="summaryInput"></textarea></div>
        <div align="left"><label>Priority:</label> <select id="priority" name="priority" v-model="priorityInput">
            <option value="high">High</option>
            <option value="medium">Medium</option>
            <option value="low">Low</option>
        </select>
        </div>
        <div>
          <button v-on:click="updateTodo()">Update</button>
          <button v-on:click="cancel()">Cancel</button>
        </div>
    </form>
  </div>
</template>

<script>
import Vue from "vue";
import axios from "axios";
import VueAxios from "vue-axios";

Vue.use(VueAxios, axios);

export default {
  name: "TodoEditor",
  data: function() {
    return {
      currentTodo: Object,
      summaryInput: "",
      priorityInput: "",
      statusInput: ""
    };
  },
  methods: {
    updateTodo: function() {
      const updatedTodo = {
        id: currentTodo.id,
        summary: this.summaryInput,
        priority: this.priorityInput,
        status: this.statusInput
      };
      Vue.axios.put(
        "http://localhost:9999/v1/todos" + currentTodo.id,
        updatedTodo
      );
    }
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

form {
  /* Just to center the form on the page */
  margin: 0 auto;
  width: 400px;
  /* To see the outline of the form */
  padding: 1em;
  border: 1px solid #ccc;
  border-radius: 1em;
}

form div + div {
  margin-top: 1em;
}

label {
  /* To make sure that all labels have the same size and are properly aligned */
  display: inline-block;
  width: 90px;
  text-align: right;
}

input,
textarea {
  /* To make sure that all text fields have the same font settings
       By default, textareas have a monospace font */
  font: 1em sans-serif;

  /* To give the same size to all text field */
  width: 300px;
  -moz-box-sizing: border-box;
  box-sizing: border-box;

  /* To harmonize the look & feel of text field border */
  border: 1px solid #999;
}

input:focus,
textarea:focus {
  /* To give a little highlight on active elements */
  border-color: #000;
}

textarea {
  /* To properly align multiline text fields with their labels */
  vertical-align: top;

  /* To give enough room to type some text */
  height: 5em;

  /* To allow users to resize any textarea vertically
       It does not work on all browsers */
  resize: vertical;
}
</style>
