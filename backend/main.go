package main

/*
 * Imports required to run the program
 */
import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const ID = 0

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
	ID        int
	Name      string
	Status    TaskStatus
	StartDate string
	EndDate   string
}

var tasks []Task

/*
 * Struct that defines a response
 */
type CustomResponse struct {
	HttpCode int
	Message  string
	Response interface{}
}

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
	json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: tasks})
}
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	log.Print("Get Task with ID: " + id + " requested")
	if value, err := strconv.Atoi(id); err == nil {
		for _, item := range tasks {
			if item.ID == value {
				json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: item})
				return
			}
		}
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 404, Message: "Not Found", Response: "No task found with id: " + id})
	} else {
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 400, Message: "Bad request", Response: "Id is not a number"})
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = getLastID()
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 201, Message: "Created", Response: task})
	log.Print("Created Task with ID: " + strconv.Itoa(task.ID))
}

func getLastID() int {
	var task Task
	task = tasks[len(tasks)-1]
	return task.ID + 1
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {}

func CreateMockData() {
	tasks = append(tasks, Task{ID: 1, Name: "Create a REST API", Status: Ongoing, StartDate: "20/07/2018"})
	tasks = append(tasks, Task{ID: 2, Name: "Create a Client in IONIC", Status: Pending, StartDate: "27/07/2018"})
	tasks = append(tasks, Task{ID: 3, Name: "Publish the result", Status: Blocked})
}
