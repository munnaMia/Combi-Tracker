package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	datamodel "github.com/munnaMia/Combi-Tracker/Model"
	maltacolor "github.com/munnaMia/Combi-Tracker/internal/maltaColor"
)

// Format the table output for terminal.
const (
	tableFormaterTitle = "%s%-4s %-32s %-8s %-25s %-25s%s\n"
	tableFormaterData  = "%-4d %-32s %s%-8s%s %-25s %-25s\n"
)

// validate argument based on the application arguments
func ValidateArgs(arg []string, applicationCmds []string) ([]string, error) {

	// Close the program if user doesn't provide any commnads
	if len(arg) < 2 {
		HandleError(errors.New("usage: combi-tracker <command> [arguments]"))
	}

	// If it's a valid command from the command list then retrun the command
	if slices.Contains(applicationCmds, arg[1]) {
		return arg[1:], nil
	}

	// Else retrun a error that the user command is not exist
	return nil, errors.New("usage: combi-tracker <command> [arguments]")
}

// handling any err with log.fatal
func HandleError(err error) {
	if err != nil {
		log.Fatal(maltacolor.Red, err, maltacolor.Reset) // Error in red color
	}
}

// This function used to print data as a helper function it can take array or a single string
func PrintData[T string | []string | []datamodel.Model | datamodel.Model](text T) {
	log.Println(text)
}

// Print a success message and return a string with a green color text
func SuccessMsg(msg string, id int) string {
	// Formate the Success text.
	return fmt.Sprintf("%s %s %d%s \n", maltacolor.Green, msg, id, maltacolor.Reset)
}

// converted a array of string to a single string
func ConvertArrayToString(array []string) string {
	return strings.Join(array, " ")
}

// Create a file included it's directory if not exist.
func CreateFileIfNotExist(filePath string) {
	absFilepath, err := filepath.Abs(filePath)
	HandleError(err) // handling the error

	// Create the directory if doesn't exixt
	dir := filepath.Dir(absFilepath)

	_, err = os.Stat(dir) // Checking the file information

	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755) // Creating a directory if not exist.
		HandleError(err)
	}

	// Create the file if doex'nt exist.
	_, err = os.Stat(absFilepath)

	if os.IsNotExist(err) {
		file, err := os.Create(absFilepath)
		HandleError(err)
		defer file.Close()

		_, err = file.Write([]byte("[]")) // Write an empty array to the json
		HandleError(err)
	}
}

// Read a json file and return the json data as a Array
func ReadJson(filePath string) []datamodel.Model {
	var data []datamodel.Model

	file, err := os.Open(filePath)
	HandleError(err)

	defer file.Close()

	byteValue, errRead := io.ReadAll(file)
	HandleError(errRead)

	err = json.Unmarshal(byteValue, &data)
	HandleError(err)

	return data
}

// Writing on a json file
func WriteJson(filePath string, tasks []datamodel.Model) {
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	HandleError(err)

	file, errC := os.Create(filePath)
	HandleError(errC)
	defer file.Close()

	_, err = file.Write(jsonData)
	HandleError(err)
}

// Print Data in a table form to the CLI
func PrintTask(task datamodel.Model) {
	/*
		Using Bright Yellow color to highlight the Title bar
		using setStatusColor to set the status color based on "todo", "done", "in-progress"
	*/
	fmt.Printf(tableFormaterTitle, maltacolor.BrightYellow, "ID", "Description", "Status", "CreatedAt", "UpdatedAt", maltacolor.Reset)
	fmt.Printf(tableFormaterData, task.Id, task.Description, setStatusColor(task.Status), task.Status, maltacolor.Reset, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
}

// Set a color for the task.Status attribute
func setStatusColor(status string) string {
	if status == "todo" {
		return maltacolor.Red
	} else if status == "done" {
		return maltacolor.Green
	} else if status == "in-progress" {
		return maltacolor.Blue
	}
	return ""
}

// Search if the id exist.
func SearchId(tasks []datamodel.Model, id int) bool {
	for _, task := range tasks {
		if task.Id == id {
			return true
		}
	}
	return false
}

// Delete a task form []task
func DeleteTask(tasks []datamodel.Model, id int) []datamodel.Model {
	tasks = slices.Delete(tasks, id-1, id) // Remove a task form the slice using id
	return tasks
}

// Sort a task list and update there IDs
func SortTask(tasks []datamodel.Model) []datamodel.Model {
	for idx := range tasks {
		tasks[idx].Id = idx + 1
	}

	return tasks // return the sort list

}
