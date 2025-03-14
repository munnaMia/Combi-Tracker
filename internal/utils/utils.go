package utils

import "errors"

func validateArgs(arg string, args []string) (string, error) {
	// Running a loop over all commands
	for _, argValue := range args {
		if arg == argValue {
			// If it's a valid command from the command list then retrun the command
			return arg, nil
		}
	}
	// Else retrun a error that the user command is not exist
	return "", errors.New("Enter a valid argument!")
}
