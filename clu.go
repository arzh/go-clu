package clu

import (
	"os"
	"fmt"
)

// Function typedef of the initializer function
// 	User will need to provide this function to set
// 	up any flags or values to be parsed
type ArgInitializer func(ArgConf)
type AppInitializer func(AppConf)

// Parses the command line arguments and fills out an Args 
// based on the initialization from User
func Parse(init ArgInitializer) Args {
	// Set it all up
	args := newArgs()
	init(args)

	parser(args, lex(os.Args[1:])) // Grab all the args after the exe name

	return *args
}

func Run(init AppInitializer) {
	app := newApp()
	init(app)

	parser(&app.Args, lex(os.Args[1:]))

	if err := app.Run(); err != nil {
		fmt.Println(err.Error())
	}
}


