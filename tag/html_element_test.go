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

func Test_compileNonBooleanAttribute_shouldCompileAttributeWithValue(t *testing.T) {
	el := new(HTMLElement)
	assert.Equal(t, el.compileNonBooleanAttribute("attrName", "attrValue"), " attrName=\"attrValue\"")
}

func Test_compileBooleanAttribute_shouldCompileAttributeWithoutValue(t *testing.T) {
	el := new(HTMLElement)
	assert.Equal(t, el.compileBooleanAttribute("attrName"), " attrName")
}

func Test_BooleanAttribute_shouldAddBooleanAttributeToBooleanAttributesMap(t *testing.T) {
	el := new(HTMLElement)
	el.BooleanAttribute("attr1", true)
	el.BooleanAttribute("attr2", false)
	assert.Equal(t, 2, len(el.booleanAttributes))
	assert.Equal(t, true, el.booleanAttributes["attr1"])
	assert.Equal(t, false, el.booleanAttributes["attr2"])
}

func Test_Attribute_shouldAddAttributeToAttributesMap(t *testing.T) {
	el := new(HTMLElement)
	el.Attribute("attr1", "val1")
	el.Attribute("attr2", "val2")
	assert.Equal(t, "val1", el.attributes["attr1"])
	assert.Equal(t, "val2", el.attributes["attr2"])
}

func Test_compileNonBooleanAttributes_shouldCompileNonBooleanAttributes(t *testing.T) {
	el := new(HTMLElement)
	el.Attribute("attr1", "val1")
	el.Attribute("attr2", "val2")
	assert.Equal(t, " attr1=\"val1\" attr2=\"val2\"", el.compileNonBooleanAttributes())
}

func Test_compileBooleanAttributes_shouldCompileBooleanAttributes(t *testing.T) {
	el := new(HTMLElement)
	el.BooleanAttribute("attr1", true)
	el.BooleanAttribute("attr2", false)
	el.BooleanAttribute("attr3", true)
	assert.Equal(t, " attr1 attr3", el.compileBooleanAttributes())
}

func Test_compileAttributes_shouldCompileAllAttributes(t *testing.T) {
	el := new(HTMLElement)
	el.Attribute("attr1", "val1")
	el.Attribute("attr2", "val2")
	el.BooleanAttribute("attr3", true)
	el.BooleanAttribute("attr4", true)
	assert.Equal(t, " attr1=\"val1\" attr2=\"val2\" attr3 attr4", el.compileAttributes())
}