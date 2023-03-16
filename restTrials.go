package main

import ( //importing all required packages
	"errors"   //create custom error messages
	"net/http" //returns the status code of the response

	"github.com/gin-gonic/gin" //imports the
)

type todo struct { //creating a JSON message structure with name todo
	ID        string `json: "id"` //the content within the `` is used for sending data to the local server
	Item      string `json: "item"`
	Completed bool   `json: "completed"`
}

var todos = []todo{ //array of todos in the local server
	{ID: "1", Item: "Clean room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodos(context *gin.Context) { //function to return all todos in local server in json
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) { //function to add a new todo to the local server
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(id string) (*todo, error) { //function to return the todo from list of todos
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getSingleTodo(context *gin.Context) { //function to return the todo in json format
	id := context.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func updateTodoStatus(context *gin.Context) { //function to update status of todo, return in json
	id := context.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todos)
}

func main() { //main method to call all API functions
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getSingleTodo)
	router.PUT("/todos/:id", updateTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}

//all results are seen using the Postman application
