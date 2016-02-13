package tag

import (
	"strings"
	"fmt"
	"github.com/doojin/gohtml/util"
)

// Compilable interface represents the compilable to string structure
type Compilable interface {
	Compile() string
}

// HTMLElement represents a DOM element
type HTMLElement struct {
	pattern           string
	children          []Compilable
	attributes        map[string]interface{}
	booleanAttributes map[string]interface{}
}

// Append adds child element to children slice
func (el *HTMLElement) Append(node Compilable) {
	el.children = append(el.children, node)
}

// Attribute adds attribute to HTML element
func (el *HTMLElement) Attribute(name, value string) {
	if el.attributes == nil {
		el.attributes = map[string]interface{}{}
	}
	el.attributes[name] = value
}

// Boolean attribute adds or removes a boolean attribute to/from HTML element
func (el *HTMLElement) BooleanAttribute(name string, value bool) {
	if el.booleanAttributes == nil {
		el.booleanAttributes = map[string]interface{}{}
	}
	el.booleanAttributes[name] = value
}

// Compile returns a textual representation of HTMLElement
func (el *HTMLElement) Compile() (result string) {
	result = strings.Replace(el.pattern, "{{children}}", el.compileChildren(), -1)
	result = strings.Replace(result, "{{attributes}}", el.compileAttributes(), -1)
	return result
}

func (el *HTMLElement) compileChildren() (output string) {
	for _, child := range el.children {
		output += child.Compile()
	}
	return
}

func (el *HTMLElement) compileAttributes() (output string) {
	output = el.compileNonBooleanAttributes()
	output += el.compileBooleanAttributes()
	return
}

func (el *HTMLElement) compileNonBooleanAttributes() (output string) {
	attrNames := util.SortedMapKeys(el.attributes)
	for _, attrName := range attrNames {
		output += el.compileNonBooleanAttribute(attrName, el.attributes[attrName])
	}
	return
}

func (el *HTMLElement) compileBooleanAttributes() (output string) {
	attrNames := util.SortedMapKeys(el.booleanAttributes)
	for _, attrName := range attrNames {
		if el.booleanAttributes[attrName].(bool) {
			output += el.compileBooleanAttribute(attrName)
		}
	}
	return
}

func (el *HTMLElement) compileNonBooleanAttribute(attrName string, attrValue interface{}) string {
	template := " %s=\"%s\""
	return fmt.Sprintf(template, attrName, attrValue)
}

func (el *HTMLElement) compileBooleanAttribute(attrName interface{}) string {
	template := " %s"
	return fmt.Sprintf(template, attrName)
}