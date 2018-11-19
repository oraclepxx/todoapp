package db

import (
	"encoding/json"
	"errors"
	"time"

	"../types"
	log "github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
)

const TODO_DB = "todos.db"

var TODO_BUCKET = []byte("todo")

func InitDB() error {

	// rwxrwxrwx
	db, err := bolt.Open(TODO_DB, 0600, &bolt.Options{Timeout: 15 * time.Second})
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Info("BoltDB connected")

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(TODO_BUCKET)
		if err != nil {
			log.Error("could not create root bucket: %v", err)
			return err
		}
		log.Info("<todo> bucket created")
		return nil
	})

	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}

func GetBolt() (*bolt.DB, error) {
	db, err := bolt.Open(TODO_DB, 0600, &bolt.Options{Timeout: 15 * time.Second})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func GreateTodo(db *bolt.DB, todoItem types.TodoItem) error {
	uid := todoItem.Id
	itemBytes, err := json.Marshal(todoItem)
	if err != nil {
		log.Error("could not marshal json: %v", err)
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket(TODO_BUCKET).Put([]byte(uid), itemBytes)
		if err != nil {
			log.Error("could not set todo: %v", err)
			return err
		}
		return nil
	})

	if err != nil {
		log.Error("Failed to create todo item %+v", err)
		return err
	}

	log.Info("todo item <" + uid + "> saved")
	return nil
}

func GetTodo(db *bolt.DB, id string) (*types.TodoItem, error) {
	var todoItem types.TodoItem
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(TODO_BUCKET)
		if bucket == nil {
			return errors.New("could not get bolt bucket: %v")
		}

		value := bucket.Get([]byte(id))

		err := json.Unmarshal(value, &todoItem)
		if err != nil {
			log.Error("Failed to umarshal data")
			return err
		}

		return nil
	})

	if err != nil {
		log.Error("Failed to get todo item")
		return nil, err
	}

	log.Info("Get the todo item")

	return &todoItem, nil
}

func GetTodos(db *bolt.DB) (*types.TodoItems, error) {
	var todoItems []types.TodoItem
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(TODO_BUCKET)

		if bucket == nil {
			return errors.New("could not get bolt bucket: %v")
		}

		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var todoItem types.TodoItem
			err := json.Unmarshal(v, &todoItem)
			if err != nil {
				log.Error("Failed to umarshal data")
				return err
			}
			todoItems = append(todoItems, todoItem)
		}

		return nil
	})

	if err != nil {
		log.Error("Failed to get all todo items")
		return nil, err
	}

	log.Info("Get the todo list")

	return &types.TodoItems{TodoItems: todoItems}, nil
}

func UpdateTodo(db *bolt.DB, todoItem types.TodoItem) error {
	itemBytes, err := json.Marshal(todoItem)
	if err != nil {
		log.Error("could not marshal json: %v", err)
		return err
	}

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(TODO_BUCKET)
		if bucket == nil {
			return errors.New("could not get bolt bucket: %v")
		}
		err := bucket.Put([]byte(todoItem.Id), itemBytes)
		return err
	})

	if err != nil {
		log.Error("Failed to update todo item %v", err)
		return err
	}

	log.Info("todo item <" + todoItem.Id + "> updated")

	return nil
}

func DeleteTodo(db *bolt.DB, todoId string) error {
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(TODO_BUCKET)
		if bucket == nil {
			return errors.New("could not get bolt bucket: %v")
		}
		err := bucket.Delete([]byte(todoId))
		if err != nil {
			log.Error("Failed to delete todo")
			return err
		}
		return nil
	})

	log.Info("Todo item <" + todoId + "> deleted")

	return nil
}
