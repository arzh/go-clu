package clu

import "os"

// Function typedef of the initilizer function
// 	User will need to provide this function to set
// 	up any flags or values to be parsed
type SetupFunc func(ArgSet)

// Parses the command line arguments and fills out an Args 
// based on the initilization from User
func Parse(init SetupFunc) Args {
	// Set it all up
	a := newArgs()
	init(a)

	// Parse all args after the name of the app
	parser(a, lex(os.Args[1:])) // Grab all the args after the app name

	return a
}


