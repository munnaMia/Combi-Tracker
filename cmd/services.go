package cmd

import (
	"time"

	datamodel "github.com/munnaMia/Combi-Tracker/Model"
	"github.com/munnaMia/Combi-Tracker/internal/utils"
)

type Application struct {
	Commands    []string
	SubCommands []string
	TodoDb      string
}

// Adding a task
func (app *Application) Add(argsArray []string, todoDbPath string) {
	taskDiscription := utils.ConvertArrayToString(argsArray[1:]) // taking the task discription and remove the args

	todoArray := utils.ReadJson(todoDbPath) // pending task list

	// Create a new task
	newTask := datamodel.Model{
		Id:          len(todoArray) + 1,
		Description: taskDiscription,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}

	todoArray = append(todoArray, newTask) // appending the new task to the pending tasks array

	utils.WriteJson(todoDbPath, todoArray) // Write into the todo JSON file or database

	utils.PrintTask(newTask) // Print the new task
}

// Delete a task
func (app *Application) Delete(argsArray []string, todoDbPath string){
	// Extract the ID from arguments
	// Validate the ID
	// Check the ID exist or not into the DB
		// If exist then Delete the task by slicing it off
	// Sort the DB IDs
}
