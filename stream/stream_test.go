package stream

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

//  Ending a file's name with _test.go tells the go test command that this file contains test functions.

// TODO try Table tests, see reflect.DeepEqual, try -cover flag and -coverprofile=coverage.out and go tool cover -html=coverage.out
// mocking with gomock (and testify?)

func TestShape(t *testing.T) {

	list2 := List([]string{"a", "b", "c"}).
		Filter(func(item interface{}) bool {
			return strings.HasPrefix(item.(string), "a")
		}).
		Map(func(item interface{}) string {
			return item.(string) + "1"
		}).
		Collect()

	assert.Equal(t, []string{"a1"}, list2)

	assert.Equal(t, 1, 1)
}
