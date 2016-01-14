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

	assert.Equal(t, "<htmlElement><child1></child1></htmlElement>", htmlElement.Compile())
}

func Test_Compile_shouldReplaceAttributesPlaceholder(t *testing.T) {
	htmlElement := new(HTMLElement)
	htmlElement.pattern = "<htmlElement{{attributes}}/>"
	htmlElement.Attribute("width", "250")
	htmlElement.Attribute("height", "100")
	assert.Equal(t, "<htmlElement height=\"100\" width=\"250\"/>", htmlElement.Compile())
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

func Test_sortedAttrNames_shouldReturnSortedAttributeNamesOfHTMLElement(t *testing.T) {
	htmlElement := new(HTMLElement)
	htmlElement.attributes = map[string]string{
		"uDummyAttribute": "uDummyValue",
		"iDummyAttribute": "iDummyValue",
		"aDummyAttribute": "aDummyValue",
	}
	expected := []string{"aDummyAttribute", "iDummyAttribute", "uDummyAttribute"}
	assert.Equal(t, expected, htmlElement.sortedAttrNames())
}

func Test_compileAttributes_shouldCompileAttributesOfHTMLElementCorrectly(t *testing.T) {
	htmlElement := new(HTMLElement)
	htmlElement.attributes = map[string]string{
		"uDummyAttribute": "uDummyValue",
		"iDummyAttribute": "iDummyValue",
		"aDummyAttribute": "aDummyValue",
	}
	expected := " aDummyAttribute=\"aDummyValue\" iDummyAttribute=\"iDummyValue\" uDummyAttribute=\"uDummyValue\""
	assert.Equal(t, expected, htmlElement.compileAttributes())
}