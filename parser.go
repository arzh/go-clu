package clu

import "os"

// Lexing
// Check first rune
//	if '/' then start getting name
//	if '-' check next rune
//		if '-' then start getting name
//		else start getting name
//	else add to Loosies
// Getting name
//	Keep reading till '='
//		if 'name' is a Flag then log error
//		if 'name' is a Var then start getting value
//	Keep reading till end
//		if 'name' is a Flag then set to true
//		if 'name' is a Var then continue to next arg as-var
//		else log error

const ( // lex states
	lstPre = iota // Check for first rune
	lstPre2 // second rune check for '-'
	lstName // read the name
	lstValue // read the value
)

// A string that no one should ever want to use as a value
const nilValue = "NIL_STRING_IFSOMEONEUSESTHISASANAGRICANTHELPTHEM"

// lexes an arge and returns it's parts
// returns name and value string
// will ALWAYS return name only returns value if using '=' notation
func lexArg(arg string) (name, value string){
// This is the long hand way, I dont really care for it much right now but it will do.
	state := lstPre
	name = ""
	value = nilValue // This is to let the arg parser that a value wasn't read

	for _, e := range arg {
		switch(state) {
		case lstPre:
			if e == '/' {
				state = lstName
			} else if e == '-' {
				state = lstPre2
			} else { // This is a loosie
				name = arg
				return // we're done here
			}
		case lstPre2:
			if e != '-' { // if we only have one '-' we need to add the current rune to the name
				name += string(e)
			}
			state = lstName
		case lstName:
			if e == '=' {
				value = "" // set the value to empty so we can fill it properly
				state = lstValue
			} else {
				name += string(e)
			}
		case lstValue:
			value += string(e)

		}
	}

	return
}

// TODO: fillout iArgs with needed helper-func
// TODO: Add error checking logic!
func parseCMD(a *iArgs) {
	args := os.Args[1:] // Grab all the args after the app name

	varName := ""

	for _, e := range args {
		name, value := lexArg(e)

		if a.putFlag(name) {
			continue
		} else if varName != "" {
			a.putVar(varName, name)
			varName = ""
		} else if a.isVar(name) {
			if value != nilValue {
				a.putVar(name, value)
			} else {
				varName = name;
				continue
			}
		} else {
			a.putLoosie(name)
		}
	}
}