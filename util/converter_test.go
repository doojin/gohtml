package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_IntSliceToString_shouldConvertIntSliceToString(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.Equal(t, "1 2 3 4 5", IntSliceToString(slice))
}
