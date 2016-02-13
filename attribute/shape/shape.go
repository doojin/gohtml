package shape

import "github.com/doojin/gohtml/attribute"

// Shape represents shape attribute value
type Shape struct {
	attribute.Attribute
}

// New returns a new instance of Shape
func New(value string) Shape {
	shape := Shape{}
	shape.Value = value
	return shape
}

var (
	Circle = New("circle")
	Default = New("default")
	Poly = New("poly")
	Rect = New("rect")
)