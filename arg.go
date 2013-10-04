package clu

import (
	"time"
	"strconv"
)

// Interface for setting Args, exposed to the initlization function
type ArgSet interface {
	SetFlag(string, string, string)
	SetVar(string, string, string, string)
}

// Interface for getting command line arguments
type Args interface {
	Var(string) string
	VarInt(string) (int64, error)
	VarFloat(string) (float64, error)
	VarDuration(string) (time.Duration, error)

	Flag(string) bool

	LenLoose() int
	Loosie(uint) string
}

// Internal implementation of Args interface
type iArgs struct {
	values map[string]*string
	flags map[string]*bool
	help map[string]*string
	loosies []string
}

// Creates an initilized iArgs pointer
func newArgs() *iArgs {
	a := new(iArgs)

	a.values = make(map[string]*string)
	a.flags = make(map[string]*bool)
	a.help = make(map[string]*string)
	a.loosies = make([]string, 0)

	return a
}

// helper to add to the 'help' section
func (a *iArgs) addHelp(name, shortcut, usage string) {
	h := new(string)
	(*h) = usage
	a.help[name] = h
	a.help[shortcut] = h
}

// Create a new Var to track
func (a *iArgs) SetVar(name, shortcut, dvalue, usage string) {
	s := new(string)
	(*s) = dvalue
	a.values[name] = s
	a.values[shortcut] = s

	a.addHelp(name, shortcut, usage)
}

// Get Var
// return empty string if not found
func (a *iArgs) Var(s string) string {
	if sp, ok := a.values[s]; ok {
		return (*sp)
	}

	return ""
}

// Get Var as int64
func (a *iArgs) VarInt(s string) (int64, error) {
	return strconv.ParseInt(a.Var(s), 10, 64)
}

// Get Var as uint64 (currently not exposed)
func (a *iArgs) VarUInt(s string) (uint64, error) {
	return strconv.ParseUint(a.Var(s), 10, 64)
}

// Get Var as float64
func (a *iArgs) VarFloat(s string) (float64, error) {
	return strconv.ParseFloat(a.Var(s), 64)
}

// Get Var as time.Duration
func (a *iArgs) VarDuration(s string) (time.Duration, error) {
	return time.ParseDuration(a.Var(s))
}

// Create a new Flag to track
func (a *iArgs) SetFlag(name, shortcut, usage string) {
	f := new(bool)
	(*f) = false
	a.flags[name] = f
	a.flags[shortcut] = f

	a.addHelp(name, shortcut, usage)
}

// Get Flag
// return false if not found
func (a *iArgs) Flag(s string) bool {
	if bp, ok := a.flags[s]; ok {
		return (*bp)
	}

	return false
}

// Get the number of loosies
func (a *iArgs) LenLoose() int {
	return len(a.loosies)
}

// Get loosie at an index
// return empty string if invalid index
func (a *iArgs) Loosie(i uint) string {
	if i >= uint(len(a.loosies)) {
		return ""
	}

	return a.loosies[i]
}

// Tries to set a flag,
// if the given flag is not present in the map return false
func (a *iArgs) putFlag(s string) bool {
	fp, ok := a.flags[s]
	
	if ok {
		(*fp) = true;
	}

	return ok
}

func (a *iArgs) isVar(s string) bool {
	_, ok := a.values[s]
	return ok
}

// Tries to set a value
// If value isn't present then return false
func (a *iArgs) putVar(s, v string) bool {
	vp, ok := a.values[s]

	if ok {
		(*vp) = v
	}

	return ok
}

// Adds a loosie to the list,
// returns true to maintain a bit of consistancy between all the 'puts'
func (a *iArgs) putLoosie(l string) bool {
	a.loosies = append(a.loosies, l)
	return true
}