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
		DbPath: "internal/database/db.go",
	}
	
	argsArray, err := utils.ValidateArgs(os.Args, app.Commands) // validating the arguments. and storing all the arguments in a slice.

	utils.HandleError(err) // If any error occur program will be shutdown

	
	switch argsArray[0] {
	case "add":
		app.Add(argsArray, app.DbPath) // Adding a task to JSON
	}

}
