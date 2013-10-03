package clu




type SetupFunc func(*ArgSet)

func Parse(init SetupFunc) *Args {
	a := newArgs()
	init(a)

	//Parsing though all the args here
	parseCMD(a)

	return a
}


