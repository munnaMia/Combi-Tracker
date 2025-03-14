package main

import (
	"fmt"
	"os"

	"github.com/munnaMia/Combi-Tracker/cmd"
	"github.com/munnaMia/Combi-Tracker/internal/utils"
)

func main() {
	app := &cmd.Application{
		Commands: []string{"add", "update", "delete" ,"mark-in-progress", "mark-done", "list"},
		SubCommands: []string{"done", "todo", "in-progress"},
	}


	// Taking the command line args without progs
	argsWithoutProg, err := utils.ValidateArgs(os.Args, app.Commands)

	utils.HandleError(err) // If any error exist program will imediate shutdown
	
	fmt.Println(argsWithoutProg)

}
