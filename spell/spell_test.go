package spell_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/kohpai/go-test-101/spell"
)

func TestAutoCorrect(t *testing.T) {
	words := []string{
		"accaptable",
		"accidentilly",
		"accommadate",
		"acqire",
		"alot",
		"amatuer",
		"apparint",
	}

	expects := []string{
		"acceptable",
		"accidentally",
		"accommodate",
		"acquire",
		"allot",
		"amateur",
		"apparent",
	}

	for i, word := range words {
		got := spell.AutoCorrect(word)
		exp := expects[i]
		if got != exp {
			t.Errorf("unexpected value for %q. got: %q, exp: %q", word, got, exp)
		}
	}
}

func TestAutoCorrectFromFile(t *testing.T) {
	data := loadFile(t, "./testdata/terms.txt")
	terms := strings.Split(data, "\n")
	for i, term := range terms {
		if len(term) > 0 {
			f := strings.Fields(term)
			if len(f) != 2 {
				t.Errorf("unexpected input: %s, line: %d", term, i+1)
				continue
			}
			if got, exp := spell.AutoCorrect(f[0]), f[1]; got != exp {
				t.Errorf("unexpected value for %q. got: %q, exp: %q", term, got, exp)
			}
		}
	}
}

func loadFile(t *testing.T, file string) string {
	t.Helper()
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(b)
}
