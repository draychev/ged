package inparser

import (
	"fmt"
	"testing"
)

type Test struct {
	input    string
	expected *Command
}

var tests = []Test{
	{"p" /*------*/, &Command{From: ".", Comma: false, To: ".", Command: "p", Parameters: nil}}, // print current line only
	{",p" /*-----*/, &Command{From: "1", Comma: true, To: "$", Command: "p", Parameters: nil}},  // prints everything
	{"." /*------*/, &Command{From: ".", Comma: false, To: ".", Command: "p", Parameters: nil}}, // shows current line only
	{"$" /*------*/, &Command{From: "$", Comma: false, To: "$", Command: "p", Parameters: nil}}, // shows last line
	{"1" /*------*/, &Command{From: "1", Comma: false, To: "1", Command: "p", Parameters: nil}}, // shows first line
	{"1,5" /*----*/, &Command{From: "5", Comma: true, To: "5", Command: "p", Parameters: nil}},  // shows only 5!!
	{"1,5p" /*---*/, &Command{From: "1", Comma: true, To: "5", Command: "p", Parameters: nil}},  // shows 1 through 5
	{",n" /*-----*/, &Command{From: "1", Comma: true, To: "$", Command: "n", Parameters: nil}},
	{"3,n" /*----*/, &Command{From: "3", Comma: true, To: "3", Command: "n", Parameters: nil}},
	{",5n" /*----*/, &Command{From: "1", Comma: true, To: "5", Command: "n", Parameters: nil}},
	{".p" /*-----*/, &Command{From: ".", Comma: false, To: ".", Command: "p", Parameters: nil}},
	{".n" /*-----*/, &Command{From: ".", Comma: false, To: ".", Command: "n", Parameters: nil}},
	{"$" /*------*/, &Command{From: "$", Comma: false, To: "$", Command: "p", Parameters: nil}},
	{"." /*------*/, &Command{From: ".", Comma: false, To: ".", Command: "p", Parameters: nil}},
	{".,$n" /*---*/, &Command{From: ".", Comma: true, To: "$", Command: "n", Parameters: nil}},
	{"2,4t279" /**/, &Command{From: "2", Comma: true, To: "4", Command: "t", Parameters: []string{"279"}}},
	{"2,4" /*----*/, &Command{From: "4", Comma: true, To: "4", Command: "p", Parameters: nil}},
	{"2,2" /*----*/, &Command{From: "2", Comma: true, To: "2", Command: "p", Parameters: nil}},
	{"2,$" /*----*/, &Command{From: "$", Comma: true, To: "$", Command: "p", Parameters: nil}},
	{"2,." /*----*/, &Command{From: ".", Comma: true, To: ".", Command: "p", Parameters: nil}},
	{"1,5u" /*---*/, nil}, // no-such-command
	{"1u" /*-----*/, nil}, // no-such-command
	{"u" /*------*/, nil}, // no-such-command
	{"1,5p6" /*--*/, nil}, // no-such-command
	{"5,1p" /*---*/, nil}, // no-such-command
	{"2,4m6" /*--*/, &Command{From: "2", Comma: true, To: "4", Command: "m", Parameters: []string{"6"}}},
	{"2m6" /*----*/, &Command{From: "2", Comma: true, To: "2", Command: "m", Parameters: []string{"6"}}},
	{"2x" /*-----*/, &Command{From: "2", Comma: true, To: "2", Command: "x", Parameters: nil}},
	{"x" /*------*/, &Command{From: ".", Comma: true, To: ".", Command: "x", Parameters: nil}},
	{".x" /*-----*/, &Command{From: ".", Comma: true, To: ".", Command: "x", Parameters: nil}},
	{"1,2x" /*---*/, nil},
	{"x3" /*-----*/, nil},
	{"=" /*---_--*/, &Command{From: ".", Comma: false, To: ".", Command: "=", Parameters: nil}},
	{"10in" /*_--*/, &Command{From: "10", Comma: false, To: "10", Command: "i", Parameters: []string{"n"}}},
	// {"1,5p6" /*---*/, &Command{From: "1", To: "6", Command: "p", Parameters: []string{"6"}}}, /** THIS SHOULD FAIL **/
}

func TestCommandInput(t *testing.T) {
	for idx, test := range tests {
		if test.expected == nil {
			fmt.Printf("\n\n----------\n%00d: Testing w/ '%s' (expecting a no-such-command)\n", idx, test.input)
		} else {
			fmt.Printf("\n\n----------\n%00d: Testing w/ '%s' (expecting: %+v)\n", idx, test.input, *test.expected)
		}
		actual, err := ReadCommand(test.input)
		if err != nil {
			if test.expected == nil { // this is expected
				// special case for when we actually expect an error
				fmt.Printf("%00d: OK -- (got the error expected)\n", idx)
				continue // move on to the next test
			}
			fmt.Printf("Test failed w/ %s: \n\texpected: %+v \n\tactually: error %+v\n", test.input, *test.expected, err)
			t.Fatal(err)
		}

		if test.expected == nil && err != nil {
			fmt.Printf("For input %s --> \nexpected: no-such-command \nactually: %+v\n", test.input, *actual)
			t.Errorf("For input %s --> \nexpected: no-such-command \nactually: %+v\n", test.input, *actual)
		}

		if !actual.Equal(test.expected) && test.expected == nil {
			fmt.Printf("For input %s --> \nexpected: an error \nactually: %+v\n", test.input, *actual)
			t.Errorf("For input %s --> \nexpected: an error \nactually: %+v\n", test.input, *actual)
		} else if !actual.Equal(test.expected) {
			fmt.Printf("For input %s --> \nexpected: %+v \nactually: %+v\n", test.input, *test.expected, *actual)
			t.Errorf("For input %s --> \nexpected: %+v \nactually: %+v\n", test.input, *test.expected, *actual)
		}

		fmt.Printf("%00d: OK -- %+v\n", idx, *actual)
	}
}
