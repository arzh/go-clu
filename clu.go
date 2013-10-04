package clu

import "os"


type SetupFunc func(ArgSet)

func Parse(init SetupFunc) Args {
	// Set it all up
	a := newArgs()
	init(a)

	// Parse all args after the name of the app
	parser(a, lex(os.Args[1:])) // Grab all the args after the app name

	return a
}


