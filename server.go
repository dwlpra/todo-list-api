package main

import (
	"fmt"
	"net/http"

	"github.com/dwlpra/todo-list-api/libs/database"
	"github.com/dwlpra/todo-list-api/pkg/activity"
	"github.com/dwlpra/todo-list-api/pkg/todo"
)

func main() {
	// create DB connection
	connDB, err := database.CreateConnection()
	if err != nil {
		fmt.Println("failed connect to database")
		return
	}
	activityRepo := activity.NewRepository(connDB)
	activityUc := activity.NewUseCase(&activityRepo)
	activityHl := activity.NewHandler(&activityUc)

	todoRepo := todo.NewRepository(connDB)
	todoUc := todo.NewUseCase(&todoRepo)
	todoHl := todo.NewHandler(&todoUc)

	activityHl.Route()
	todoHl.Route()

	fmt.Println("starting web server at http://localhost:3030/")
	http.ListenAndServe(":3030", nil)
}
