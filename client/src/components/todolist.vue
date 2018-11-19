<template>
  <div class="todolist">
    <h3>Todo List</h3>
    <hr>
    <table v-if="allTodos.length > 0">
      <thead>
           <tr>
               <th><input type="checkbox" v-model="checkAll"></th>
               <!-- <th>Index</th> -->
               <th>Summary</th>
               <th>Priority</th>
               <th>Status</th>
               <th>Actions</th>
           </tr>
      </thead>
      <tbody>
             <tr v-for="(item) in allTodos" :key="item.id">
                  <th><input type="checkbox" :id="item.id" :value="item.id" v-model="checked"></th>
                  <!-- <td>{{ index + 1 }}</td> -->
                  <td>{{ item.summary }}</td>
                  <td>{{item.priority}}</td>
                  <td>{{item.status}}</td>
                  <td><button>Update</button> | <button v-on:click="deleteTodo(item.id)">Delete</button></td>
            </tr>
      </tbody>
    </table>
    <h4 v-else> - No data available- </h4>
    <div>{{checked}}</div>
  </div>
  
</template>

<script>
import Vue from "vue";
import axios from "axios";
import VueAxios from "vue-axios";

Vue.use(VueAxios, axios);

export default {
  name: "TodoList",

  data: function() {
    return {
      allTodos: [],
      checked: []
    };
  },

  created() {
    this.listAll();
  },

  computed: {
    checkAll: {
            get: function () {
                return this.allTodos ? this.checked.length == this.allTodos.length : false;
            },
            set: function (value) {
                var selected = [];

                if (value) {
                    this.allTodos.forEach(todo => {
                        selected.push(todo.id);
                    });
                }

                this.checked = selected;
            }
        }
  },

  methods: {
    listAll() {
      Vue.axios.get("http://localhost:9999/v1/todos").then(resp => {
        this.allTodos = resp.data.todoItems;
      });
    },

    deleteTodo(todoId) {
      Vue.axios
        .delete("http://localhost:9999/v1/todos/" + todoId)
        .then(resp => {
          if (resp.status == 200) {
            this.listAll();
          }
        })
        .catch(err => {
          alert(err);
        });
    },

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

table {
  /* Just to center the form on the page */
  margin: 0 auto;
  width: 800px;
  /* To see the outline of the form */
  padding: 1em;
  text-align: left;
}

th {
  background-color: #f2f2f2;
}

td {
  padding: 8px;
}

tr:nth-child(even) {
  background-color: #f2f2f2;
}

hr {
  margin: 0 auto;
  width: 800px;
}
</style>
