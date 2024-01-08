package inparser

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/draychev/go-toolbox/pkg/logger"
)

var log = logger.NewPretty("inparser")

// Reading the ed command is going to be a state machine.
// ed commands adhere to a strict grammar.
// See the grammar with `info ed`.
// Example:
//
//	'(.,.)t(.)'
//	     Copies (i.e., transfers) the addressed lines to after the right-hand
//	     destination address. If the destination address is '0' (zero), the
//	     lines are copied at the beginning of the buffer. The current address is
//	     set to the address of the last line copied.
func ReadCommand(input string) (*Command, error) {

	/*
	   source: https://www.gnu.org/software/ed/manual/ed_manual.html
	   >>> [address[,address]]command[parameters]
	*/

	var ( /** --- states --- **/
		from       []rune
		hasComma   bool
		to         []rune
		command    []rune
		parameters []rune
	)

	var curState *[]rune   // pointer to the current sate (block above)
	var canHaveParams bool // is the command allowed to have params
	var done bool          // indicates when we are expected to be done
	curState = &from       // we start with FROM

	reader := bufio.NewReader(strings.NewReader(input))

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			if err.Error() == "EOF" {
				break // we are done
			}
			log.Error().Err(err).Msg("Error reading input") // TODO: return error here?
			return nil, errReadingInput
		}

		if curState == &parameters && !canHaveParams {
			return nil, errUnexpectedChar
		}

		// Check to see if the state machine found everything it needed.
		// This check is only to return an error if there are characters beyond.
		if done {
			// should have been done by now
			return nil, errUnexpectedBeyondEnd
		}

		// decisions

		if char == '=' {
			return &Command{
				From:    ".",
				Comma:   false,
				To:      ".",
				Command: "=",
			}, nil
		}

		// Once you reach a COMMA - switch accumulation (TO) and move on
		if char == ',' {
			hasComma = true
			if curState != &from {
				return nil, errTooManyCommas
			}
			curState = &to
			continue
		}

		if char >= '0' && char <= '9' {
			*curState = append(*curState, char)
			continue
		}

		// this is for command only (not params)
		if char >= 'a' && char <= 'z' && curState != &command && curState != &parameters {
			commandConfig, isValidCommand := validCommands[char]
			if !commandConfig.canHaveComma && hasComma {
				return nil, errNoCommaAllowed
			}
			canHaveParams = commandConfig.canHaveParams
			if !isValidCommand {
				return nil, errInvalidCommand
			}

			if curState == &command && len(command) > 0 { // only 1 letter allowed
				return nil, errInvalidCommand
			}

			curState = &command
			*curState = append(*curState, char)
			if !canHaveParams { // commands such as 'p' are terminal - nothing should follow them
				done = true // this won't work for 1,2t3
			}

			curState = &parameters
			continue
		}

		if char == '.' {
			*curState = append(*curState, char)
			continue
		}

		if char == '$' {
			*curState = append(*curState, char)
			continue
		}

		if curState == &parameters {
			// TODO: check for allowed parameters
			// commandConfig, isValidCommand := validCommands[char]
			*curState = append(*curState, char)
		}

	}

	cmd := &Command{
		From:    combineRunes(from),
		Comma:   hasComma,
		To:      combineRunes(to),
		Command: combineRunes(command),
	}

	// TODO: this is not complete or even well thought out
	if parameters != nil {
		cmd.Parameters = []string{combineRunes(parameters)}
	}

	if err := validate(cmd); err != nil {
		return nil, err
	}

	return makeCommandExplicit(cmd), nil
}

func validate(c *Command) error {
	// TODO: check if from and to are only digits - can't have spaces, can't have ".1" etc can't have "$4" etc
	// TODO: command should be a single character - not more than one
	if c.Command == "p" && c.Parameters != nil {
		return errUnexpectedBeyondEnd
	}

	fromNum, errFrom := strconv.Atoi(c.From)
	toNum, errTo := strconv.Atoi(c.To)
	if errFrom == nil && errTo == nil && toNum < fromNum {
		return errInvalidFrom
	}

	return nil
}

func makeCommandExplicit(c *Command) *Command {
	// This function aims to make it as easy as posible for the next stage
	// to understand the intent of the user. So it will turn the implicit
	// into explicit.

	if c.Comma && c.From == "" && c.To == "" {
		c.From = "1"
		c.To = "$"
	}

	if c.Comma && c.From == "" && c.To != "" {
		c.From = "1"
	}

	// covers commands such as "." and "2,"
	if c.From != "" && c.To == "" {
		c.To = c.From
	}

	if c.Command == "" {
		if c.From == "" {
			c.From = "."
		}
		if c.To == "" {
			c.To = "."
		}
		if c.Comma {
			c.From = c.To
		}
		c.From = c.To // when there is no command - it p and the from becomes t
	}

	// this must be last
	if c.Command == "" {
		c.Command = "p"
	}

	if c.From == "" && c.To == "" {
		c.From = "."
		c.To = "."
	}

	return c
}

func combineRunes(runes []rune) string {
	s := make([]string, len(runes))
	for _, r := range runes {
		s = append(s, string(r))
	}
	return strings.Join(s, "")
}
