package stream

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

// TODO
// generify the stream
// mocking with gomock ("github.com/golang/mock/gomock") might be better to test on regular app service

func TestFilterMap(t *testing.T) {

	list2 := Stream{list: []string{"a", "b", "c"}}.
		Filter(func(item interface{}) bool {
			return strings.HasPrefix(item.(string), "a")
		}).
		Map(func(item interface{}) string {
			return item.(string) + "1"
		}).
		ToList()

	assert.Equal(t, []string{"a1"}, list2)

	assert.Equal(t, 1, 1)
}

func TestFilterSuccesses(t *testing.T) {

	startsWithA := func(item interface{}) bool {
		return strings.HasPrefix(item.(string), "a")
	}

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

	// There is a DeepEqual method to compare complex data structures,
	// but you don't need it for assertions... assert.Equal will ultimately call DeepEqual under the hood if necessary
	reflect.DeepEqual([]string{"a"}, []string{"a"})

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Stream{list: test.input}.Filter(startsWithA).ToList()
			assert.Equal(t, test.expected, actual)
		})
	}

}
