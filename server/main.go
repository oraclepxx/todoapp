package main

import (
	"encoding/json"
	"net/http"

	"./db"
	"./types"
	"./utils"
	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

func init() {
	err := db.InitDB()
	if err != nil {
		log.Error()
	}
}

type TodoServer struct {
	r *httprouter.Router
}

func (s *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, PUT, GET, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Accept")
	w.Header().Set("Content-Type", "application/json")
	s.r.ServeHTTP(w, r)
}

func main() {

	router := httprouter.New()

	// router.OPTIONS("/v1/todos", handleCorsPreflight)
	router.GET("/v1/todos", handleListTodos)
	router.POST("/v1/todos", handleCreateTodo)

	// router.OPTIONS("/v1/todos/:id", handleCorsPreflight)
	router.GET("/v1/todos/:id", handleGetTodo)
	router.PUT("/v1/todos/:id", handlePutTodo)
	router.DELETE("/v1/todos/:id", handleDeleteTodo)

	err := http.ListenAndServe(":9999", &TodoServer{router})
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Magic happens on port: 9999")

}

func handleCreateTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var newItem types.TodoItem
	err := decoder.Decode(&newItem)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Failed to create todo item")
		return
	}

	uuidv4, err := utils.GenerateUUID()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Failed to create uuid for todo item")
		return
	}
	newItem.Id = uuidv4

	boltdb, err := db.GetBolt()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Failed to connect db")
		return
	}

	db.GreateTodo(boltdb, newItem)

	utils.WriteJSON(w, http.StatusOK, newItem)

	defer boltdb.Close()
	return
}

func handleListTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	boltdb, err := db.GetBolt()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Failed to connect db")
		return
	}

	todos, err := db.GetTodos(boltdb)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Failed to get all todos from db")
		return
	}

	utils.WriteJSON(w, http.StatusOK, todos)

	defer boltdb.Close()

	return
}

func handleGetTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoId := params.ByName("id")

	boltdb, err := db.GetBolt()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Failed to connect db")
		return
	}

	todo, err := db.GetTodo(boltdb, todoId)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Failed to get todo from db")
		return
	}

	utils.WriteJSON(w, http.StatusOK, todo)

	defer boltdb.Close()
	return
}

func handlePutTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoId := params.ByName("id")

	var newItem types.TodoItem
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newItem)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Failed to create todo item")
		return
	}

	newItem.Id = todoId

	boltdb, err := db.GetBolt()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Failed to connect db")
		return
	}

	err = db.UpdateTodo(boltdb, newItem)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Failed to get todo from db")
		return
	}

	utils.WriteJSON(w, http.StatusOK, newItem)

	defer boltdb.Close()
	return
}

func handleDeleteTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoId := params.ByName("id")

	boltdb, err := db.GetBolt()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Failed to connect db")
		return
	}

	err = db.DeleteTodo(boltdb, todoId)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Failed to delete todo from db")
		return
	}

	utils.WriteJSON(w, http.StatusOK, "success")

	defer boltdb.Close()
	return 
}

func handleCorsPreflight(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// log.Infof("Received [OPTIONS] request to CorsPreFlight: %+v", r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, PUT, GET, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Accept")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}
