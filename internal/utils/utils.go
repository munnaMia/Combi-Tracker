package utils

import (
	"errors"
	"log"
	"os"
	"slices"
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
