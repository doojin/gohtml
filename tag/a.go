package tag

import (
	"github.com/doojin/gohtml/util"
	"github.com/doojin/gohtml/attribute/lang"
	"github.com/doojin/gohtml/attribute/rel"
	"github.com/doojin/gohtml/attribute/shape"
	"strconv"
)

type a struct {
	HTMLElement
}

// NewA returns an instance of a element
func NewA() *a {
	aElement := new(a)
	aElement.pattern = "<a{{attributes}}>{{children}}</a>"
	return aElement
}

// SetAccessKey sets accesskey attribute
func (aElement *a) SetAccesskey(c byte) {
	aElement.Attribute("accesskey", string(c))
}

// SetCoords sets coords attribute
func (aElement *a) SetCoords(coords ...int) {
	aElement.Attribute("coords", util.IntSliceToString(coords))
}

// SetDownload sets download attribute
func (aElement *a) SetDownload(download bool) {
	aElement.BooleanAttribute("download", download)
}

// SetHref sets href attribute
func (aElement *a) SetHref(href string) {
	aElement.Attribute("href", href)
}

// SetHreflang sets hreflang attribute
func (aElement *a) SetHreflang(hrefLang lang.Lang) {
	aElement.Attribute("hreflang", hrefLang.Value)
}

// SetName sets name attribute
func (aElement *a) SetName(name string) {
	aElement.Attribute("name", name)
}

// SetRel sets rel attribute
func (aElement *a) SetRel(rel rel.Rel) {
	aElement.Attribute("rel", rel.Value)
}

// SetRev sets rev attribute
func (aElement *a) SetRev(rev string) {
	aElement.Attribute("rev", rev)
}

// SetShape sets shape attribute
func (aElement *a) SetShape(shape shape.Shape) {
	aElement.Attribute("shape", shape.Value)
}

// SetTabindex sets tabindex attribute
func (aElement *a) SetTabindex(tabIndex int) {
	aElement.Attribute("tabindex", strconv.Itoa(tabIndex))
}

// SetTarget sets target attribute
func (aElement *a) SetTarget(target string) {
	aElement.Attribute("target", target)
}

// SetTitle sets title attribute
func (aElement *a) SetTitle(title string) {
	aElement.Attribute("title", title)
}

// TODO: create a separate struct one day
// SetType sets type attribute
func (aElement *a) SetType(t string) {
	aElement.Attribute("type", t)
}