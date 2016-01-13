package tag

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Compile_shouldReplaceChildrenPlaceholder(t *testing.T) {
	child1 := new(HTMLElement)
	child1.pattern = "<child1></child1>"

	htmlElement := new(HTMLElement)
	htmlElement.Append(child1)
	htmlElement.pattern = "<htmlElement>{{children}}</htmlElement>"

	assert.Equal(t, htmlElement.Compile(), "<htmlElement><child1></child1></htmlElement>")
}

func Test_compileChildren_shouldCompileChildrenElementsCorrectly(t *testing.T) {
	child1 := new(HTMLElement)
	child1.pattern = "<child1></child1>"

	child2 := new(HTMLElement)
	child2.pattern = "<child2></child2>"

	htmlElement := new(HTMLElement)
	htmlElement.Append(child1)
	htmlElement.Append(child2)

	assert.Equal(t, htmlElement.compileChildren(), "<child1></child1><child2></child2>")
}
