package cmd

import (
	datamodel "github.com/munnaMia/Combi-Tracker/Model"
	"github.com/munnaMia/Combi-Tracker/internal/utils"
)

type Application struct {
	dataModel   datamodel.Model
	Commands    []string
	SubCommands []string
	DbPath      string
}

// Add a task
func (app *Application) Add(argsArray []string, dbPath string) {

	// taskDiscription := utils.ConvertArrayToString(argsArray[1:]) // taking the task discription and remove the args

	utils.CreateFileIfNotExist(dbPath)
	// create or store to the JSON db
	// return if success of failed

}
