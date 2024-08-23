package stream

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func TestFilterMap(t *testing.T) {

	letters := []string{"a", "b", "c"}
	prefixedLetters := Filter(letters, func(item string) bool {
		return strings.HasPrefix(item, "a")
	})
	appendedLetters := Map(prefixedLetters, func(item string) string {
		return item + "1"
	})

	assert.Equal(t, []string{"a1"}, appendedLetters)

	assert.Equal(t, 1, 1)
}

func TestFilterSuccesses(t *testing.T) {

	// Table Test using maps, see https://semaphoreci.com/blog/table-driven-unit-tests-go

	tests := map[string]struct {
		input    []string
		expected []string
	}{
		"filter finds a": {
			input:    []string{"a", "b", "c"},
			expected: []string{"a"},
		},
		"filter works with empty inputs": {
			input:    []string{},
			expected: []string{},
		},
		"filter works with empty strings": {
			input:    []string{""},
			expected: []string{},
		},
	}

	startsWithA := func(item string) bool {
		return strings.HasPrefix(item, "a")
	}

	// There is a DeepEqual method to compare complex data structures,
	// but you don't need it for assertions... assert.Equal will ultimately call DeepEqual under the hood if necessary
	reflect.DeepEqual([]string{"a"}, []string{"a"})

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Filter(test.input, startsWithA)
			assert.Equal(t, test.expected, actual)
		})
	}

}
