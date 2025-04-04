package utils

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

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
	return nil, errors.New("enter a valid argument")
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// This function used to print data as a helper function it can take array or a single string
func PrintData[T string | []string](text T) {
	log.Println(text)
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
		err := os.Mkdir(dir, 0755) // Creating a directory if not exist.
		HandleError(err)
	}

	// Create the file if doex'nt exist.
	_, err = os.Stat(absFilepath)

	if os.IsNotExist(err) {
		file, err := os.Create(absFilepath)
		HandleError(err)
		defer file.Close()
	}
}
