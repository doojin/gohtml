package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_SortedMapKeys_shouldReturnListOfSortedMapKeys(t *testing.T) {
	m := map[string]interface{}{
		"c": 1,
		"a": 2,
		"b": 3,
	}
	sortedKeys := SortedMapKeys(m)
	assert.Equal(t, 3, len(sortedKeys))
	assert.Equal(t, "a", sortedKeys[0])
	assert.Equal(t, "b", sortedKeys[1])
	assert.Equal(t, "c", sortedKeys[2])
}
