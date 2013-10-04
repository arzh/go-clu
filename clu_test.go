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

func testAppInit(a ArgSet) {
	a.SetFlag("verbose", "v", "turns on verbose logging")
	a.SetFlag("debug", "d", "adds debug hooks")

	a.SetFlag("mark", "x", "mark todo item as done")
	a.SetFlag("update", "u", "updates todo item")
	a.SetFlag("last_updated", "lu", "changes ordering")

	a.SetVar("list_name", "name", "default", "list to apply to")
}

// TODO: Add more robust testing
func TestParsing(t *testing.T) {
	a := newArgs()
	testAppInit(a)
	parser(a, lex([]string{"'first note'"}))
	if a.LenLoose() != 1 {
		t.Error("Loosie length incorrect Exp:", 1, "Got:", a.LenLoose())
	} else if a.Loosie(0) != "first note" {
		t.Error("Loosie detection failed Exp: first note Got:", a.Loosie(0))
	}

}