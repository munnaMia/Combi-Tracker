package cmd

import (
	"fmt"
	"time"

	datamodel "github.com/munnaMia/Combi-Tracker/Model"
	"github.com/munnaMia/Combi-Tracker/internal/utils"
)

type Application struct {
	Commands    []string
	SubCommands []string
	TodoDb      string
	ProgDb      string
	DoneDb      string
}

// Adding a task
func (app *Application) Add(argsArray []string, todoDbPath string) {
	taskDiscription := utils.ConvertArrayToString(argsArray[1:]) // taking the task discription and remove the args

	// i have to find the current ID for a new task
	todoArray := utils.ReadJson(todoDbPath) // pending work list

	// Create a new task
	newTask := datamodel.Model{
		Id:          len(todoArray),
		Description: taskDiscription,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}
	
	// appending the new task to the pending tasks array
	todoArray = append(todoArray, newTask)
	
	// Write into the todo JSON file or database
	// utils.WriteJson(todoDbPath, todoArray)

}
