package clu

import (
	"time"
	"strconv"
)

// Interface for setting Args, exposed to the initlization function
type ArgConf interface {
	AddFlag(string, string, string)
	AddVar(string, string, string, string)
}

// Internal implementation of Args and ArgSet interface
type Args struct {
	values map[string]*string
	flags map[string]*bool
	help map[string]*string
	loosies []string
}

// Creates an initilized Args pointer
func newArgs() *Args {
	args := new(Args)

	args.values = make(map[string]*string)
	args.flags = make(map[string]*bool)
	args.help = make(map[string]*string)
	args.loosies = make([]string, 0)

	return args
}

// helper to add to the 'help' section
func (args *Args) addHelp(name, shortcut, usage string) {
	h := new(string)
	(*h) = usage
	args.help[name] = h
	args.help[shortcut] = h
}

// Create args new Var to track
func (args *Args) AddVar(name, shortcut, dvalue, usage string) {
	s := new(string)
	(*s) = dvalue
	args.values[name] = s
	args.values[shortcut] = s

	args.addHelp(name, shortcut, usage)
}

// Get Var
// return empty string if not found
func (args *Args) Var(s string) string {
	if sp, ok := args.values[s]; ok {
		return (*sp)
	}

	return ""
}

// Get Var as int64
func (args *Args) VarInt(s string) (int64, error) {
	return strconv.ParseInt(args.Var(s), 10, 64)
}

// Get Var as uint64 (currently not exposed)
func (args *Args) VarUInt(s string) (uint64, error) {
	return strconv.ParseUint(args.Var(s), 10, 64)
}

// Get Var as float64
func (args *Args) VarFloat(s string) (float64, error) {
	return strconv.ParseFloat(args.Var(s), 64)
}

// Get Var as time.Duration
func (args *Args) VarDuration(s string) (time.Duration, error) {
	return time.ParseDuration(args.Var(s))
}

// Create args new Flag to track
func (args *Args) AddFlag(name, shortcut, usage string) {
	f := new(bool)
	(*f) = false
	args.flags[name] = f
	args.flags[shortcut] = f

	args.addHelp(name, shortcut, usage)
}

// Get Flag
// return false if not found
func (args *Args) Flag(s string) bool {
	if bp, ok := args.flags[s]; ok {
		return (*bp)
	}

	return false
}

// Get the number of loosies
func (args *Args) LenLoose() int {
	return len(args.loosies)
}

// Get loosie at an index
// return empty string if invalid index
func (args *Args) Loosie(i uint) string {
	if i >= uint(len(args.loosies)) {
		return ""
	}

	return args.loosies[i]
}

// Returns all loosies
func (args *Args) Loosies() []string {
	return args.loosies
}

// Tries to set args flag,
// if the given flag is not present in the map return false
func (args *Args) putFlag(s string) bool {
	fp, ok := args.flags[s]

	if ok {
		(*fp) = true;
	}

	return ok
}

func (args *Args) isVar(s string) bool {
	_, ok := args.values[s]
	return ok
}

// Tries to set args value
// If value isn't present then return false
func (args *Args) putVar(s, v string) bool {
	vp, ok := args.values[s]

	if ok {
		(*vp) = v
	}

	return ok
}

// Adds args loosie to the list,
// returns true to maintain args bit of consistancy between all the 'puts'
func (args *Args) putLoosie(l string) bool {
	args.loosies = append(args.loosies, l)
	return true
}
