package clu

// Command Line Utility libary
// Small helper libary for creating apps with command line arguments
//
// Flags: 
//		Should except windows (/f) or UNIX (-|--) style
//	SetFlag(name, shortcut, usage)
//		Flags only test for presence so they are in effect just bool flags
//	Flag(string) bool
//
// Vars: Much like go Flags
//		Will accept win and unix style as well as = or space notation
//	SetVar(name, shortcut, usage)
//	Var(name) string
//		There should also be a conversion get for most types
//			int, uint64, float64, duration (bool is a flag, and string is the standard return)
//		syntax will be Var<type>(name) <type>
//
// Loosies: Any args that are not a defined flag or var
//		Look up only
//	LenLoose() int
//	Loosie(int) string
//		The biggest part of this is that order matters
//
// App: Struct that will hold the Flag and Loosie data
//	Flags should be in a map[string]bool
//	Vars should be a map[string]string, converstion can be gone at 'get' time
//	Loosies can just be a []string
