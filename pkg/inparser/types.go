package inparser

import (
	"errors"
	"strings"
)

var (
	errUnexpectedBeyondEnd error = errors.New("unexpected characters beyond the semantic end of the command")
	errTooManyCommas             = errors.New("too many commas")
	errUnexpectedChar            = errors.New("unexpected character")
	errUnexpectedState           = errors.New("unexpected state")
	errInvalidCommand            = errors.New("invalid command")
	errInvalidFrom               = errors.New("invalid from")
	errNoCommaAllowed            = errors.New("no comma allowed")
	errReadingInput              = errors.New("error reading input")
)

type Command struct {
	From       string   // could be a <number>, blank, ".", or "$"
	Comma      bool     // whether the command had a comma
	To         string   // could be a <number>, blank, ".", or "$"
	Command    string   // should be a single character
	Parameters []string // could be an array of strings
}

func (lh *Command) Equal(rh *Command) bool {
	return (rh != nil &&
		lh.From == rh.From &&
		lh.To == rh.To &&
		lh.Command == rh.Command &&
		strings.Join(lh.Parameters, "-") == strings.Join(rh.Parameters, "-"))
}

type CommandConfig struct {
	canHaveParams bool
	canHaveComma  bool
}

var validCommands = map[rune]CommandConfig{
	'a': CommandConfig{false, true},                               // add lines
	'c': CommandConfig{false, true},                               // change lines
	'd': CommandConfig{false, true},                               // delete line
	'i': CommandConfig{true, false},                               // insert lines
	'p': CommandConfig{false, true},                               // print lines
	't': CommandConfig{true, true},                                // transfer / copy lines
	'n': CommandConfig{false, true},                               // print with numbers
	'm': CommandConfig{true, true},                                // move lines
	'x': CommandConfig{canHaveParams: false, canHaveComma: false}, // tag?
	'q': CommandConfig{canHaveParams: false, canHaveComma: false}, // quit
	'w': CommandConfig{canHaveParams: false, canHaveComma: false}, // write
	'=': CommandConfig{canHaveParams: false, canHaveComma: false}, // show current line number
}
