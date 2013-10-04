package clu

import (
	"testing"
)

func TestLexing(t *testing.T) {
	tests := []string{
		"/help", "/h", "/?", "--h", "-h", "-var=10", "--really_long_flag=this_is_just_getting_out_of_hand", "/Ddeveloper=true"
	}

	units := []string{
		"help", nilValue, "h", nilValue, "?", nilValue, "h", nilValue, "h", nilValue, "var", "10", "really_long_flag", "this_is_just_getting_out_of_hand",
	}

	for it, iu := 0, 0; it < len(tests); it++ {
		n, v := lexArg(tests[it])
		if n != units[iu] || v != units[iu+1] {
			t.Error("Failed test", it, "\nExp:", units[iu], units[iu+1], "\nGot:", n, v)
		}
		iu += 2
	}
}

func testAppInit(a *ArgSet) {
	a.SetFlag("verbose", "v", "turns on verbose logging")
	a.SetFlag("debug", "d", "adds debug hooks")

	a.SetFlag("mark", "x", "mark todo item as done")
	a.SetFlag("update", "u", "updates todo item")
	a.SetFlag("last_updated", "lu", "changes ordering")

	a.SetVar("list_name", "name", "default", "list to apply to")
}

func TestParsing(t *testing.T) {
	a := newArgs()
	testAppInit(a)
	// TODO: Testing stuff here, idk
	parseCMD(a, )
}