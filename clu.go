package clu

import "os"


type SetupFunc func(ArgSet)

func Parse(init SetupFunc) Args {
	a := newArgs()
	init(a)

	//Parsing though all the args here
	parseCMD(a, os.Args[1:]) // Grab all the args after the app name
)

	return a
}


