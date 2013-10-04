package clu

import (
	"strings"
)

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

type lexState func(*lexer, string) lexState

type lexer struct {
	state lexState
	out chan string
	item string
	raw []string
}

// ranges over raw data and runs it though the lexer
func (l *lexer) run() {
	for _, e := range l.raw {
		l.state = l.state(l, e)
	}

	close(l.out)
}

// pushes the currect saved item and resets
func (l *lexer) putItem() {
	l.out <- l.item
	l.item = ""
}

// pushes over the out channel
func (l *lexer) put(s string) {
	l.out <- s
}

// standard arg lexing
func lstArg(l *lexer, s string) lexState {
	offset := 0

	// Test for prefixs
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "/") {
		offset = 1
		if strings.HasPrefix(s, "--") {
			offset = 2
		}
	}

	// Split on '=' to test for value setting
	sp := strings.Split(s[offset:], "=")
	offset = 0
	if len(sp) > 1 {
		l.put(sp[0])
		offset = 1
	}

	i := sp[offset]
	
	// Are starting a quoted section?
	if strings.HasPrefix(i, "'") {
		if strings.HasSuffix(i, "'") {
			i = strings.Trim(i, "'")
		} else {
			l.item = strings.Trim(i, "'")
			return lstInQuote
		}
	}

	l.put(i)
	return lstArg
}

// We keep adding to item until we find an end quote
func lstInQuote(l *lexer, s string) lexState {

	l.item = strings.Join([]string{l.item, strings.Trim(s, "'")}, " ")

	if strings.HasSuffix(s, "'") {
		l.putItem()
		return lstArg
	}

	return lstInQuote
}

// Make and run the lexer, passing back the arg stream
func lex(in []string) (chan string) {
	l := &lexer {
		state: lstArg,
		out: make(chan string, 3),
		item: "",
		raw: in,
	}

	go l.run()

	return l.out
}

// TODO: Add error checking logic!
func parser(a *iArgs, in chan string) {
	for arg := range in {

		if a.putFlag(arg) {
			continue
		} else if a.isVar(arg) {
			if v, ok := <-in; ok {
				a.putVar(arg, v)			 
			} else {
				return
			}
		} else {
			a.putLoosie(arg)
		}
	}
}