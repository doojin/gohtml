package tag

import (
	"strings"
	"sort"
)

// HTMLElement represents a DOM element
type HTMLElement struct {
	pattern    string
	children   []*HTMLElement
	attributes map[string]string
}

// Append adds child element to children slice
func (el *HTMLElement) Append(node *HTMLElement) {
	el.children = append(el.children, node)
}

// Attribute adds attribute to attributes map
func (el *HTMLElement) Attribute(name, value string) {
	if el.attributes == nil {
		el.attributes = map[string]string{}
	}
	el.attributes[name] = value
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
	template := " {{attributeName}}=\"{{attributeValue}}\""
	attrNames := el.sortedAttrNames()
	for _, attrName := range attrNames {
		attribute := strings.Replace(template, "{{attributeName}}", attrName, -1)
		attribute = strings.Replace(attribute, "{{attributeValue}}", el.attributes[attrName], -1)
		output += attribute
	}
	return
}

func (el *HTMLElement) sortedAttrNames() (attrNames []string) {
	for attrName := range el.attributes {
		attrNames = append(attrNames, attrName)
	}
	sort.Strings(attrNames)
	return
}