package tag

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/doojin/gohtml/attribute/lang"
	"github.com/doojin/gohtml/attribute/rel"
	"github.com/doojin/gohtml/attribute/shape"
)

func Test_A_Pattern(t *testing.T) {
	a := NewA()
	a.Attribute("attr1", "val1")
	a.Attribute("attr2", "val2")
	a.Append(NewTextNode("child1"))
	assert.Equal(t, "<a attr1=\"val1\" attr2=\"val2\">child1</a>", a.Compile())
}

func Test_A_SetAccesskey_shouldSetAccesskeyAttribute(t *testing.T) {
	a := NewA()
	a.SetAccesskey('a')
	assert.Equal(t, "a", a.attributes["accesskey"])
}

func Test_A_SetCoords_shouldSetCoordsAttribute(t *testing.T) {
	a := NewA()
	a.SetCoords(1, 2, 3, 4)
	assert.Equal(t, "1 2 3 4", a.attributes["coords"])
}

func Test_A_SetDownload_shouldSetDownloadAttribute(t *testing.T) {
	a := NewA()
	a.SetDownload(true)
	assert.Equal(t, true, a.booleanAttributes["download"])
}

func Test_A_SetHreflang_shouldSetHreflangAttribute(t *testing.T) {
	a := NewA()
	a.SetHreflang(lang.Abkhazian)
	assert.Equal(t, "ab", a.attributes["hreflang"])
}

func Test_A_SetRel_shouldSetRelAttribute(t *testing.T) {
	a := NewA()
	a.SetRel(rel.Answer)
	assert.Equal(t, "answer", a.attributes["rel"])
}

func Test_A_SetShape_shouldSetShapeAttribute(t *testing.T) {
	a := NewA()
	a.SetShape(shape.Circle)
	assert.Equal(t, "circle", a.attributes["shape"])
}

func Test_A_SetTabindex_shouldSetTabIndexAttribute(t *testing.T) {
	a := NewA()
	a.SetTabindex(12)
	assert.Equal(t, "12", a.attributes["tabindex"])
}