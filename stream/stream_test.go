package stream

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// TODO see reflect.DeepEqual, try -cover flag and -coverprofile=coverage.out and go tool cover -html=coverage.out
// mocking with gomock (and testify?)
// generify the stream

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

	type testCase struct {
		description string
		input       []string
		expected    []string
	}

	// TODO try with maps https://semaphoreci.com/blog/table-driven-unit-tests-go

	// This is a Table Test
	for _, scenario := range []testCase{
		{
			description: "filter finds a",
			input:       []string{"a", "b", "c"},
			expected:    []string{"a"},
		},
		{
			description: "filter works with empty inputs",
			input:       []string{},
			expected:    []string{},
		},
		{
			description: "filter works with empty strings",
			input:       []string{""},
			expected:    []string{},
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			output := Stream{list: scenario.input}.Filter(startsWithA).ToList()
			assert.Equal(t, scenario.expected, output)
		})
	}

}
