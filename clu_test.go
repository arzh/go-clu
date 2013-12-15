package clu

import (
	"testing"
)

func TestLexing(t *testing.T) {
	tests := []string{
		"/help", "/h", "/?", "--h", "-h", "-var=10", "--really_long_flag=this_is_just_getting_out_of_hand", 
		"/Ddeveloper=true", `/poopie='this is in quotes'`,
	}

	units := []string{
		"help", "h", "?", "h", "h", "var", "10", "really_long_flag", "this_is_just_getting_out_of_hand",
		"Ddeveloper", "true", "poopie", "this is in quotes",
	}

	out := lex(tests)

	i := 0
	for arg := range out {
		if arg != units[i] {
			t.Error("Failed test", i, "\nExp:", units[i], "\nGot:", arg)
		}
		i++
	}

	if i < len(units) {
		t.Error("Lexer chan closed before we go though all tests. Got to", i, "out of", len(units))
	}
}

func testArgInit(a ArgConf) {
	a.AddFlag("verbose", "v", "turns on verbose logging")
	a.AddFlag("debug", "d", "adds debug hooks")

	a.AddFlag("mark", "x", "mark todo item as done")
	a.AddFlag("update", "u", "updates todo item")
	a.AddFlag("last_updated", "lu", "changes ordering")

	a.AddVar("list_name", "name", "default", "list to apply to")
}

// TODO: Add more robust testing
func TestParsing(t *testing.T) {
	a := newArgs()
	testArgInit(a)
	parser(a, lex([]string{"'first note'"}))
	if a.LenLoose() != 1 {
		t.Error("Loosie length incorrect Exp:", 1, "Got:", a.LenLoose())
	} else if a.Loosie(0) != "first note" {
		t.Error("Loosie detection failed Exp: first note Got:", a.Loosie(0))
	}

}

var appTestRunner int
const (
	firstRun = 100
	secondRun = 1000
	thirdRun = 10000
)

func firstTest(args Args) {
	appTestRunner = firstRun;
}

func testApp1Init(a AppConf) {
	a.AddCmd("test", firstTest)
}

func secondTest(args Args) {
	if args.Flag("verbose") {
		appTestRunner = secondRun;
		return
	}
	appTestRunner = 0
}

func testApp2Init(a AppConf) {
	a.AddFlag("verbose", "v", "verbose flag")
	a.AddCmd("test2", secondTest)
}

func thirdTest(args Args) {
	if args.Loosie(0) == "run" {
		appTestRunner = thirdRun
		return
	}
	appTestRunner = 0
}

func testApp3Init(a AppConf) {
	a.AddCmd("test3", thirdTest)
}

func TestApp(t *testing.T) {
	t1 := []string{"test"}
	t2 := []string{"test2", "-v"}
	t3 := []string{"test3", "run"}

	appTestRunner = 0;
	a1 := newApp()
	testApp1Init(a1)
	parser(&a1.Args, lex(t1))
	a1.Run()

	if appTestRunner != firstRun {
		t.Error("First Test failed");
	}


	appTestRunner = 0;
	a2 := newApp()
	testApp2Init(a2)
	parser(&a2.Args, lex(t2))
	a2.Run()

	if appTestRunner != secondRun {
		t.Error("Second Test failed");
	}


	appTestRunner = 0;
	a3 := newApp()
	testApp3Init(a3)
	parser(&a3.Args, lex(t3))
	a3.Run()

	if appTestRunner != thirdRun {
		t.Error("Third Test failed");
	}
}