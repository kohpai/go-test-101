package strings_test

import (
	"strings"
	"testing"
)

type Input struct {
	sentence string
	term     string
}

type TestCase struct {
	name  string
	input Input
	exp   int
}

func TestIndex(t *testing.T) {
	tcs := []TestCase{
		{"gopher", Input{"Gophers are amazing!", "are"}, 8},
		{"testing", Input{"Testing in Go is fun.", "fun"}, 17},
		{"answer", Input{"The answer is 42.", "is"}, 11},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			input := tc.input
			got := strings.Index(input.sentence, input.term)
			exp := tc.exp
			if got != exp {
				t.Errorf("unexpected value for %v. got: %v, exp: %v", input, got, exp)
			}
		})
	}
}
