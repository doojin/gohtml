package tag

import "strings"

// HTMLElement represents a DOM element
type HTMLElement struct {
	children []*HTMLElement
	pattern  string
}

// Append adds child element to children slice
func (el *HTMLElement) Append(node *HTMLElement) {
	el.children = append(el.children, node)
}

// Compile returns a textual representation of HTMLElement
func (el *HTMLElement) Compile() string {
	return strings.Replace(el.pattern, "{{children}}", el.compileChildren(), -1)
}

func (el *HTMLElement) compileChildren() (output string) {
	for _, child := range el.children {
		output += child.Compile()
	}
	return output
}