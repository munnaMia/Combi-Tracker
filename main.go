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
		TodoDb: "internal/database/todoDb.json",
		PendDb: "internal/database/pendingDb.json",
		DoneDb: "internal/database/doneDb.json",
		
	}
	
	argsArray, err := utils.ValidateArgs(os.Args, app.Commands) // validating the arguments. and storing all the arguments in a slice.
	utils.HandleError(err) // If any error occur program will be shutdown

	// Creating all the nessary files and directory.
	utils.CreateFileIfNotExist(app.TodoDb)
	utils.CreateFileIfNotExist(app.PendDb)
	utils.CreateFileIfNotExist(app.DoneDb)

		switch argsArray[0] {
	case "add":
		app.Add(argsArray) // Adding a task to JSON
	}

}
