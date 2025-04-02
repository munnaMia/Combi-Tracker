package utils

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func ValidateArgs(arg []string, args []string) ([]string, error) {

	// Close the program if user doesn't provide any commnads
	if len(arg) < 1 {
		log.Print("Usage: combi-tracker <command> [arguments]")
		os.Exit(1)
	}

	// If it's a valid command from the command list then retrun the command
	if slices.Contains(args, arg[1]) {
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

func CreateFileIfNotExist(filePath string) {
	absFilepath, err := filepath.Abs(filePath)
	HandleError(err) // handling the error

	// creating the directory if doen't exist

	// start working here next day
	// .......

}
