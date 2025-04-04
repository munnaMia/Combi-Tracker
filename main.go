package main

import (
	"os"

	"github.com/munnaMia/Combi-Tracker/cmd"
	"github.com/munnaMia/Combi-Tracker/internal/utils"
)

func main() {
	app := &cmd.Application{
		Commands:    []string{"add", "update", "delete", "mark-in-progress", "mark-done", "list"},
		SubCommands: []string{"done", "todo", "in-progress"},
		TodoDb:      "internal/database/todoDb.json",
		ProgDb:      "internal/database/progDb.json",
		DoneDb:      "internal/database/doneDb.json",
	}

	argsArray, err := utils.ValidateArgs(os.Args, app.Commands) // validating the arguments. and storing all the arguments in a slice.
	utils.HandleError(err)                                      // If any error occur program will be shutdown

	// Creating all the nessary files and directory.
	utils.CreateFileIfNotExist(app.TodoDb)
	utils.CreateFileIfNotExist(app.ProgDb)
	utils.CreateFileIfNotExist(app.DoneDb)

	switch argsArray[0] {
	case "add":
		app.Add(argsArray, app.TodoDb) // Adding a task to JSON
	}

}
