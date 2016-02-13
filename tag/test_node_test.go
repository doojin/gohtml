package tag

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_TextNode_Pattern(t *testing.T) {
	textNode := NewTextNode("text")
	assert.Equal(t, "text", textNode.Compile())
}
