package main

import (
	"os"

	"github.com/munnaMia/Combi-Tracker/cmd"
	"github.com/munnaMia/Combi-Tracker/internal/utils"
)

func main() {
	app := &cmd.Application{
		Commands:    []string{"add", "update", "delete", "mark-in-progress", "mark-done", "list", "help"},
		SubCommands: []string{"done", "todo", "in-progress"}, // use this as sub commands with list>command
		TodoDb:      "internal/database/todoDb.json",
	}

	argsArray, err := utils.ValidateArgs(os.Args, app.Commands) // validating the arguments. and storing all the arguments in a slice.
	utils.HandleError(err)                                      // If any error occur program will be shutdown

	// Creating all the nessary files and directory.
	utils.CreateFileIfNotExist(app.TodoDb)

	switch argsArray[0] {
	case "add":
		app.Add(argsArray, app.TodoDb) // Adding a task to JSON
	case "delete":
		app.Delete(argsArray, app.TodoDb) // Delete a Single Task
	case "update":
		app.Update(argsArray, app.TodoDb) // Update a Single Task
	case "mark-in-progress":
		app.MarkInProgress(argsArray, app.TodoDb) // Mark in progress a task
	case "mark-done":
		app.MarkDone(argsArray, app.TodoDb) // Mark done a task
	case "list":
		app.List(argsArray, app.SubCommands, app.TodoDb) // Show all task or based on there status
	case "help":
		app.Help() // Show all the command
	}

}
