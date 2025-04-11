package cmd

import (
	"errors"
	"strconv"
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
	// Check the title is empty or not
	if taskDiscription == "" {
		utils.HandleError(errors.New("title is empty"))
	}
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

	utils.PrintData(utils.SuccessMsg("Task added successfully ID: ", newTask.Id))
	utils.PrintTask(newTask) // Print the new task
}

// Delete a task
func (app *Application) Delete(argsArray []string, todoDbPath string) {
	if len(argsArray) != 2 {
		utils.HandleError(errors.New("usage: combi-tracker <command> [arguments]")) // [delete <id>] if more than 2 element as input return err
	}

	taskIDItToDelete, err := strconv.Atoi(argsArray[1]) // get the id of the task to delete
	// Overwriting the error for this case only.
	if err != nil {
		utils.HandleError(errors.New("id must have to be a integer value"))
	}

	todoDatas := utils.ReadJson(todoDbPath) // Read the JSON DB.

	// Checking the given id exist or not
	if exist := utils.SearchId(todoDatas, taskIDItToDelete); !exist {
		utils.HandleError(errors.New("the given id isn't exist"))
	}

	deletedTask := todoDatas[taskIDItToDelete-1] // pick the task that select to delete for letter show to the user as output which one is deleted

	todoDatas = utils.DeleteTask(todoDatas, taskIDItToDelete) // Task deleted successfully.

	todoDatas = utils.SortTask(todoDatas) // Store the sorted list after delete the task.

	utils.WriteJson(todoDbPath, todoDatas) // Write the update data after delete.

	utils.PrintData(utils.SuccessMsg("Task deleted successfully ID: ", deletedTask.Id))
	utils.PrintTask(deletedTask) // Print the deleted task to user

}

// Update a Task
func (app *Application) Update(argsArray []string, todoDbPath string) {

	taskIdToUpdate, err := strconv.Atoi(argsArray[1]) // get the id of the task to delete
	// Overwriting the error for this case only.
	if err != nil {
		utils.HandleError(errors.New("id must have to be a integer value"))
	}

	updatedDiscription := utils.ConvertArrayToString(argsArray[2:])

	todoDatas := utils.ReadJson(todoDbPath) // Read the JSON DB.

	if exist := utils.SearchId(todoDatas, taskIdToUpdate); !exist {
		utils.HandleError(errors.New("the given id isn't exist"))
	}

	todoDatas[taskIdToUpdate-1].Description = updatedDiscription // Update the discription
	// adding the updating time
	now := time.Now()                            // Current time
	todoDatas[taskIdToUpdate-1].UpdatedAt = &now // adding the current updating time

	utils.WriteJson(todoDbPath, todoDatas)

	utils.PrintData(utils.SuccessMsg("Task updated successfully ID: ", taskIdToUpdate))
	utils.PrintTask(todoDatas[taskIdToUpdate-1])
}

// Mark in progress a task
func (app *Application) MarkInProgress(argsArray []string, todoDbPath string) {
	if len(argsArray) != 2 {
		utils.HandleError(errors.New("usage: combi-tracker <command> [arguments]")) // [delete <id>] if more than 2 element as input return err
	}

	taskIdItToMarkInProgress, err := strconv.Atoi(argsArray[1]) // get the id of the task to delete
	// Overwriting the error for this case only.
	if err != nil {
		utils.HandleError(errors.New("id must have to be a integer value"))
	}

	todoDatas := utils.ReadJson(todoDbPath) // Read the JSON DB.

	// Checking the given id exist or not
	if exist := utils.SearchId(todoDatas, taskIdItToMarkInProgress); !exist {
		utils.HandleError(errors.New("the given id isn't exist"))
	}

	todoDatas[taskIdItToMarkInProgress-1].Status = "in-progress" // Update the status to progress
	// adding the updating time
	now := time.Now()                                      // Current time
	todoDatas[taskIdItToMarkInProgress-1].UpdatedAt = &now // adding the current updating time

	utils.WriteJson(todoDbPath, todoDatas)

	utils.PrintData(utils.SuccessMsg("Task mark in progress successfully ID: ", taskIdItToMarkInProgress))
	utils.PrintTask(todoDatas[taskIdItToMarkInProgress-1])
}
