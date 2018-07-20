package main

/*
 * Imports required to run the program
 */
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
 * Enum that defines the status of task
 */
type TaskStatus int

const (
	Pending  TaskStatus = 0
	Ongoing  TaskStatus = 1
	Done     TaskStatus = 2
	Blocked  TaskStatus = 3
	Rejected TaskStatus = 4
)

/*
 * Struct that defines a task
 */
type Task struct {
	ID        string
	Name      string
	Status    TaskStatus
	StartDate string
	EndDate   string
}

var tasks []Task

// our main function
func main() {
	CreateMockData()
	router := mux.NewRouter()
	router.HandleFunc("/task", GetTasks).Methods("GET")
	router.HandleFunc("/task/{id}", GetTask).Methods("GET")
	router.HandleFunc("/task", CreateTask).Methods("POST")
	router.HandleFunc("/task/{id}", DeleteTask).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	log.Print("Get Tasks requested")
	json.NewEncoder(w).Encode(tasks)
}
func GetTask(w http.ResponseWriter, r *http.Request)    {}
func CreateTask(w http.ResponseWriter, r *http.Request) {}
func DeleteTask(w http.ResponseWriter, r *http.Request) {}

func CreateMockData() {
	tasks = append(tasks, Task{ID: "1", Name: "Create a REST API", Status: Ongoing, StartDate: "20/07/2018"})
	tasks = append(tasks, Task{ID: "2", Name: "Create a Client in IONIC", Status: Pending, StartDate: "27/07/2018"})
	tasks = append(tasks, Task{ID: "3", Name: "Publish the result", Status: Blocked})
}
