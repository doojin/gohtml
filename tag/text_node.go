package tag

type textNode struct {
	HTMLElement
}

// NewTextNode returns an instance of textNode element
func NewTextNode(text string) *textNode {
	textNodeElement := new(textNode)
	textNodeElement.pattern = text
	return textNodeElement
}